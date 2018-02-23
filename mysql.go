package main

import (
	"database/sql"
)

const (
	QueryCreateMapping = "create-mapping"
	QueryFindMapping   = "find-mapping"
	QueryExpireMapping = "expire-mapping"
)

type StoreStruct struct {
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

type MySqlConfig struct {
	username string
	password string
	host     string
	port     string
	database string
}

// NewStore will prepare all of our queries on the provided
// database and returns a pointer to a new `mysql.Store` instance.
func NewStore(db *sql.DB) (*StoreStruct, error) {
	unprepared := map[string]string{
		QueryCreateMapping: `
			INSERT INTO mappings (original_url, shortened_url, single_use, expired)
			VALUES(?, ?, ?, ?);
		`,
		QueryFindMapping: `
			SELECT id, original_url, single_use, expired
			FROM mappings
			WHERE shortened_url = ?;
		`,
		QueryExpireMapping: `
			UPDATE mappings
			SET expired = ?
			WHERE id = ?;
		`,
	}

	// prepare all statements to verify syntax
	stmts, err := prepareStmts(db, unprepared)
	if err != nil {
		return nil, err
	}

	s := StoreStruct{
		db:    db,
		stmts: stmts,
	}

	return &s, nil
}

// Close closes the database, releasing any open resources.
func (s *StoreStruct) Close() error {
	return s.db.Close()
}

// prepareStmts will attempt to prepare each unprepared
// query on the database. If one fails, the function returns
// with an error.
func prepareStmts(db *sql.DB, unprepared map[string]string) (map[string]*sql.Stmt, error) {
	prepared := map[string]*sql.Stmt{}
	for k, v := range unprepared {
		stmt, err := db.Prepare(v)
		if err != nil {
			return nil, err
		}
		prepared[k] = stmt
	}

	return prepared, nil
}
