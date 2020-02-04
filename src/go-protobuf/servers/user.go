package servers

import (
	"context"
	"database/sql"
	"go-protobuf/api"
	"go-protobuf/configs"
	"go-protobuf/repositories/caches"
	"go-protobuf/repositories/databases"
	"go-protobuf/repositories/queues"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	BaseServer
	config       *configs.Config
	db           databases.DB
	cache        caches.Cache
	queue        queues.Queue
	userDatabase databases.UserDatabase
}

// NewUserServer creates a new instance
func NewUserServer(
	config *configs.Config,
	db databases.DB,
	cache caches.Cache,
	queue queues.Queue,
	userDatabase databases.UserDatabase,
) api.UserServer {
	return &userServer{
		BaseServer: BaseServer{
			config: config,
			db:     db,
			cache:  cache,
		},
		config:       config,
		db:           db,
		cache:        cache,
		queue:        queue,
		userDatabase: userDatabase,
	}
}

func (s *userServer) GetUser(ctx context.Context, request *api.GetUserRequest) (response *api.GetUserResponse, err error) {
	var (
		user *api.UserModel
	)

	_ = copier.Copy(&user, request)

	tx, err := s.db.Begin()
	if err != nil {
		zap.S().Errorw("Failed to start a database transaction", zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.db.Rollback(tx)
		} else {
			_ = s.db.Commit(tx)
		}
	}()

	user, err = s.userDatabase.Get(ctx, tx, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.New(codes.NotFound, "Not Found").Err()
		}
		return nil, err
	}

	response = &api.GetUserResponse{}
	_ = copier.Copy(response, user)
	return response, nil
}

func (s *userServer) CreateUser(ctx context.Context, request *api.CreateUserRequest) (response *empty.Empty, err error) {
	var (
		user api.UserModel
	)

	_ = copier.Copy(&user, request)

	tx, err := s.db.Begin()
	if err != nil {
		zap.S().Errorw("Failed to start a database transaction", zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.db.Rollback(tx)
		} else {
			_ = s.db.Commit(tx)
		}
	}()

	_, err = s.userDatabase.Insert(ctx, tx, &user)
	if err != nil {
		return nil, err
	}

	response = &empty.Empty{}
	return response, nil
}

func (s *userServer) UpdateUser(ctx context.Context, request *api.UpdateUserRequest) (response *empty.Empty, err error) {
	var (
		user api.UserModel
	)

	_ = copier.Copy(&user, request)

	tx, err := s.db.Begin()
	if err != nil {
		zap.S().Errorw("Failed to start a database transaction", zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.db.Rollback(tx)
		} else {
			_ = s.db.Commit(tx)
		}
	}()

	err = s.userDatabase.Update(ctx, tx, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.New(codes.NotFound, "Not Found").Err()
		}
		return nil, err
	}

	response = &empty.Empty{}
	return response, nil
}

func (s *userServer) DeleteUser(ctx context.Context, request *api.DeleteUserRequest) (response *empty.Empty, err error) {
	tx, err := s.db.Begin()
	if err != nil {
		zap.S().Errorw("Failed to start a database transaction", zap.Error(err))
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = s.db.Rollback(tx)
		} else {
			_ = s.db.Commit(tx)
		}
	}()

	err = s.userDatabase.Delete(ctx, tx, request.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.New(codes.NotFound, "Not Found").Err()
		}
		return nil, err
	}

	response = &empty.Empty{}
	return response, nil
}
