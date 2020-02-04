package databases

import "database/sql"

// DB is repository of database
type DB interface {
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}

type defaultDB struct {
	db *sql.DB
}

// NewDefaultDB creates a new instance
func NewDefaultDB(db *sql.DB) DB {
	return &defaultDB{
		db: db,
	}
}

func (d *defaultDB) Begin() (*sql.Tx, error) {
	return d.db.Begin()
}

func (d *defaultDB) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (d *defaultDB) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
