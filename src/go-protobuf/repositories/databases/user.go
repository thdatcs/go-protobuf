package databases

import (
	"context"
	"database/sql"
	"go-protobuf/api"
	"go-protobuf/utils/jaeger"

	"github.com/opentracing/opentracing-go/ext"
)

// UserDatabase is repository of user
type UserDatabase interface {
	Get(ctx context.Context, tx *sql.Tx, username string) (*api.UserModel, error)
	List(ctx context.Context, tx *sql.Tx) ([]*api.UserModel, error)
	Insert(ctx context.Context, tx *sql.Tx, user *api.UserModel) (int64, error)
	Update(ctx context.Context, tx *sql.Tx, user *api.UserModel) error
	Delete(ctx context.Context, tx *sql.Tx, userCode string) error
}

type userDatabase struct {
	BaseDatabase
}

// NewUserDatabase creates a new instance
func NewUserDatabase() UserDatabase {
	return &userDatabase{}
}

func (d *userDatabase) Get(ctx context.Context, tx *sql.Tx, username string) (user *api.UserModel, err error) {
	span := jaeger.Start(ctx, ">databases.userDatabase/Get", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	var (
		statement = `
			SELECT username
				, password
				, fullname
			FROM user
			WHERE username=?`
	)
	user = &api.UserModel{}
	err = tx.
		QueryRow(statement,
			username).
		Scan(&user.Username,
			&user.Password,
			&user.Fullname)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *userDatabase) List(ctx context.Context, tx *sql.Tx) (users []*api.UserModel, err error) {
	span := jaeger.Start(ctx, ">databases.userDatabase/List", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	var (
		statement = `
			SELECT username
				, password
				, fullname
			FROM user`
	)
	rows, err := tx.Query(statement)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := &api.UserModel{}
		err = rows.
			Scan(&user.Username,
				&user.Password,
				&user.Fullname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (d *userDatabase) Insert(ctx context.Context, tx *sql.Tx, user *api.UserModel) (id int64, err error) {
	span := jaeger.Start(ctx, ">databases.userDatabase/Insert", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	var (
		insertStatement = `
			INSERT INTO user (
				username
				, password
				, fullname
			) VALUES (?, ?, ?)`
	)
	result, err := tx.Exec(insertStatement,
		user.Username,
		user.Password,
		user.Fullname)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (d *userDatabase) Update(ctx context.Context, tx *sql.Tx, user *api.UserModel) (err error) {
	span := jaeger.Start(ctx, ">databases.userDatabase/Update", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	var (
		nums            int64
		updateStatement = `
			UPDATE user
			SET password=?
				, fullname=?
			WHERE username=?`
	)
	result, err := tx.Exec(updateStatement,
		user.Password,
		user.Fullname,
		user.Username)
	if err != nil {
		return err
	}
	nums, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if nums == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (d *userDatabase) Delete(ctx context.Context, tx *sql.Tx, username string) (err error) {
	span := jaeger.Start(ctx, ">databases.userDatabase/Delete", ext.SpanKindRPCClient)
	defer func() {
		jaeger.Finish(span, err)
	}()

	var (
		nums            int64
		deleteStatement = `DELETE FROM user WHERE username=?`
	)
	result, err := tx.Exec(deleteStatement, username)
	if err != nil {
		return err
	}
	nums, err = result.RowsAffected()
	if err != nil {
		return err
	}
	if nums == 0 {
		return sql.ErrNoRows
	}
	return nil
}
