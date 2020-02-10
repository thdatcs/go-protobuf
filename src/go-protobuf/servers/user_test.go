package servers

import (
	"context"
	"database/sql"
	"errors"
	"go-protobuf/api"
	"go-protobuf/repositories/databases"
	databases_mocks "go-protobuf/repositories/databases/mocks"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServerTestSuite struct {
	suite.Suite
	db *databases_mocks.DB
}

func (suite *UserServerTestSuite) SetupTest() {
	suite.db = new(databases_mocks.DB)
	suite.db.On("Begin").Return(&sql.Tx{}, nil).On("Commit", &sql.Tx{}).Return(nil).On("Rollback", &sql.Tx{}).Return(nil)
}

func (suite *UserServerTestSuite) TestGetUser() {
	successUserDatabase := new(databases_mocks.UserDatabase)
	successUserDatabase.On("Get", context.Background(), &sql.Tx{}, "abc").Return(&api.UserModel{
		Username: "abc",
		Password: "abc",
		Fullname: "abc",
	}, nil)
	errorNotFoundUserDatabase := new(databases_mocks.UserDatabase)
	errorNotFoundUserDatabase.On("Get", context.Background(), &sql.Tx{}, "xyz").Return(nil, sql.ErrNoRows)
	errorOtherUserDatabase := new(databases_mocks.UserDatabase)
	errorOtherUserDatabase.On("Get", context.Background(), &sql.Tx{}, "abcxyz").Return(nil, errors.New("Other"))

	tests := []struct {
		name  string
		mocks struct {
			db           databases.DB
			userDatabase databases.UserDatabase
		}
		args struct {
			ctx     context.Context
			request *api.GetUserRequest
		}
		want    *api.GetUserResponse
		wantErr error
	}{
		{
			name: "success",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: successUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.GetUserRequest
			}{
				ctx: context.Background(),
				request: &api.GetUserRequest{
					Username: "abc",
				},
			},
			want: &api.GetUserResponse{
				Username: "abc",
				Password: "abc",
				Fullname: "abc",
			},
			wantErr: nil,
		},
		{
			name: "error not found",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: errorNotFoundUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.GetUserRequest
			}{
				ctx: context.Background(),
				request: &api.GetUserRequest{
					Username: "xyz",
				},
			},
			want:    nil,
			wantErr: status.New(codes.NotFound, "Not Found").Err(),
		},
		{
			name: "error other",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: errorOtherUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.GetUserRequest
			}{
				ctx: context.Background(),
				request: &api.GetUserRequest{
					Username: "abcxyz",
				},
			},
			want:    nil,
			wantErr: errors.New("Other"),
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			userServer := NewUserServer(nil, test.mocks.db, nil, nil, test.mocks.userDatabase)
			response, err := userServer.GetUser(test.args.ctx, test.args.request)
			if !suite.Assert().Equal(response, test.want) {
				t.Errorf("userServer.GetUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
			if !suite.Assert().Equal(err, test.wantErr) {
				t.Errorf("userServer.GetUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
		})
	}
}

func (suite *UserServerTestSuite) TestCreateUser() {
	successUserDatabase := new(databases_mocks.UserDatabase)
	successUserDatabase.On("Insert", context.Background(), &sql.Tx{}, &api.UserModel{
		Username: "abc",
		Password: "abc",
		Fullname: "abc",
	}).Return(int64(1), nil)
	errorOtherUserDatabase := new(databases_mocks.UserDatabase)
	errorOtherUserDatabase.On("Insert", context.Background(), &sql.Tx{}, &api.UserModel{
		Username: "abcxyz",
		Password: "abcxyz",
		Fullname: "abcxyz",
	}).Return(int64(0), errors.New("Other"))

	tests := []struct {
		name  string
		mocks struct {
			db           databases.DB
			userDatabase databases.UserDatabase
		}
		args struct {
			ctx     context.Context
			request *api.CreateUserRequest
		}
		want    *empty.Empty
		wantErr error
	}{
		{
			name: "success",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: successUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.CreateUserRequest
			}{
				ctx: context.Background(),
				request: &api.CreateUserRequest{
					Username: "abc",
					Password: "abc",
					Fullname: "abc",
				},
			},
			want:    &empty.Empty{},
			wantErr: nil,
		},
		{
			name: "error other",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: errorOtherUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.CreateUserRequest
			}{
				ctx: context.Background(),
				request: &api.CreateUserRequest{
					Username: "abcxyz",
					Password: "abcxyz",
					Fullname: "abcxyz",
				},
			},
			want:    nil,
			wantErr: errors.New("Other"),
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			userServer := NewUserServer(nil, test.mocks.db, nil, nil, test.mocks.userDatabase)
			response, err := userServer.CreateUser(test.args.ctx, test.args.request)
			if !suite.Assert().Equal(response, test.want) {
				t.Errorf("userServer.CreateUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
			if !suite.Assert().Equal(err, test.wantErr) {
				t.Errorf("userServer.CreateUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
		})
	}
}

func (suite *UserServerTestSuite) TestUpdateUser() {
	successUserDatabase := new(databases_mocks.UserDatabase)
	successUserDatabase.On("Update", context.Background(), &sql.Tx{}, &api.UserModel{
		Username: "abc",
		Password: "abc",
		Fullname: "abc",
	}).Return(nil)
	errorNotFoundUserDatabase := new(databases_mocks.UserDatabase)
	errorNotFoundUserDatabase.On("Update", context.Background(), &sql.Tx{}, &api.UserModel{
		Username: "xyz",
		Password: "xyz",
		Fullname: "xyz",
	}).Return(sql.ErrNoRows)
	errorOtherUserDatabase := new(databases_mocks.UserDatabase)
	errorOtherUserDatabase.On("Update", context.Background(), &sql.Tx{}, &api.UserModel{
		Username: "abcxyz",
		Password: "abcxyz",
		Fullname: "abcxyz",
	}).Return(errors.New("Other"))

	tests := []struct {
		name  string
		mocks struct {
			db           databases.DB
			userDatabase databases.UserDatabase
		}
		args struct {
			ctx     context.Context
			request *api.UpdateUserRequest
		}
		want    *empty.Empty
		wantErr error
	}{
		{
			name: "success",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: successUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.UpdateUserRequest
			}{
				ctx: context.Background(),
				request: &api.UpdateUserRequest{
					Username: "abc",
					Password: "abc",
					Fullname: "abc",
				},
			},
			want:    &empty.Empty{},
			wantErr: nil,
		},
		{
			name: "error not found",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: errorNotFoundUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.UpdateUserRequest
			}{
				ctx: context.Background(),
				request: &api.UpdateUserRequest{
					Username: "xyz",
					Password: "xyz",
					Fullname: "xyz",
				},
			},
			want:    nil,
			wantErr: status.New(codes.NotFound, "Not Found").Err(),
		},
		{
			name: "error other",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: errorOtherUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.UpdateUserRequest
			}{
				ctx: context.Background(),
				request: &api.UpdateUserRequest{
					Username: "abcxyz",
					Password: "abcxyz",
					Fullname: "abcxyz",
				},
			},
			want:    nil,
			wantErr: errors.New("Other"),
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			userServer := NewUserServer(nil, test.mocks.db, nil, nil, test.mocks.userDatabase)
			response, err := userServer.UpdateUser(test.args.ctx, test.args.request)
			if !suite.Assert().Equal(response, test.want) {
				t.Errorf("userServer.UpdateUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
			if !suite.Assert().Equal(err, test.wantErr) {
				t.Errorf("userServer.UpdateUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
		})
	}
}

func (suite *UserServerTestSuite) TestDeleteUser() {
	successUserDatabase := new(databases_mocks.UserDatabase)
	successUserDatabase.On("Delete", context.Background(), &sql.Tx{}, "abc").Return(nil)
	errorNotFoundUserDatabase := new(databases_mocks.UserDatabase)
	errorNotFoundUserDatabase.On("Delete", context.Background(), &sql.Tx{}, "xyz").Return(sql.ErrNoRows)
	errorOtherUserDatabase := new(databases_mocks.UserDatabase)
	errorOtherUserDatabase.On("Delete", context.Background(), &sql.Tx{}, "abcxyz").Return(errors.New("Other"))

	tests := []struct {
		name  string
		mocks struct {
			db           databases.DB
			userDatabase databases.UserDatabase
		}
		args struct {
			ctx     context.Context
			request *api.DeleteUserRequest
		}
		want    *empty.Empty
		wantErr error
	}{
		{
			name: "success",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: successUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.DeleteUserRequest
			}{
				ctx: context.Background(),
				request: &api.DeleteUserRequest{
					Username: "abc",
				},
			},
			want:    &empty.Empty{},
			wantErr: nil,
		},
		{
			name: "error not found",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: errorNotFoundUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.DeleteUserRequest
			}{
				ctx: context.Background(),
				request: &api.DeleteUserRequest{
					Username: "xyz",
				},
			},
			want:    nil,
			wantErr: status.New(codes.NotFound, "Not Found").Err(),
		},
		{
			name: "error other",
			mocks: struct {
				db           databases.DB
				userDatabase databases.UserDatabase
			}{
				db:           suite.db,
				userDatabase: errorOtherUserDatabase,
			},
			args: struct {
				ctx     context.Context
				request *api.DeleteUserRequest
			}{
				ctx: context.Background(),
				request: &api.DeleteUserRequest{
					Username: "abcxyz",
				},
			},
			want:    nil,
			wantErr: errors.New("Other"),
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			userServer := NewUserServer(nil, test.mocks.db, nil, nil, test.mocks.userDatabase)
			response, err := userServer.DeleteUser(test.args.ctx, test.args.request)
			if !suite.Assert().Equal(response, test.want) {
				t.Errorf("userServer.DeleteUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
			if !suite.Assert().Equal(err, test.wantErr) {
				t.Errorf("userServer.DeleteUser() = (%v, %v), want (%v, %v)", response, err, test.want, test.wantErr)
			}
		})
	}
}

func TestUserServerTestSuite(t *testing.T) {
	suite.Run(t, new(UserServerTestSuite))
}
