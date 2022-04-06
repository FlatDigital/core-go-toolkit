// Package converter methods to database/sql package to DB custom interfaces
// created only for test purposes
//nolint:sqlclosecheck
package converter

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

// DBer database interface
type DBer interface {
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (DBTxer, error)
	Close() error
	Conn(ctx context.Context) (DBConner, error)
	Driver() driver.Driver
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Ping() error
	PingContext(ctx context.Context) error
	Prepare(query string) (DBStmter, error)
	PrepareContext(ctx context.Context, query string) (DBStmter, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	SetConnMaxLifetime(d time.Duration)
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	Stats() sql.DBStats
}

// DBConner connection database interface
type DBConner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Close() error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PingContext(ctx context.Context) error
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Raw(f func(driverConn interface{}) error) (err error)
}

// DBTxer tx database interface
type DBTxer interface {
	Commit() error
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (DBStmter, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Rollback() error
	Stmt(stmt *sql.Stmt) *sql.Stmt
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}

// DBStmter stmt database interface
type DBStmter interface {
	Close() error
	Exec(args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
	Query(args ...interface{}) (DBRowser, error)
	QueryContext(ctx context.Context, args ...interface{}) (DBRowser, error)
	QueryRow(args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, args ...interface{}) *sql.Row
}

// DBRowser rows database interface
type DBRowser interface {
	Close() error
	ColumnTypes() ([]*sql.ColumnType, error)
	Columns() ([]string, error)
	Err() error
	Next() bool
	NextResultSet() bool
	Scan(dest ...interface{}) error
}

// sqlConverter map methods to sql.DB to DBer
type sqlConverter struct {
	db *sql.DB
	DBer
}

// sqlConnConverter map methods to sql.Conn to DBConner
type sqlConnConverter struct {
	conn *sql.Conn
	DBConner
}

// sqlStmtConverter map methods to sql.Stmt to DBStmter
type sqlStmtConverter struct {
	stmt *sql.Stmt
	DBStmter
}

// sqlTxConverter map methods to sql.Tx to DBTxer
type sqlTxConverter struct {
	tx *sql.Tx
	DBTxer
}

// SQLToDBer converts sql.DB to an struct that implement DBer interface
func SQLToDBer(db *sql.DB) DBer {
	return &sqlConverter{db: db}
}

// SQLConnToDBConner converts sql.Conn to an struct that implement DBConner interface
func SQLConnToDBConner(conn *sql.Conn) DBConner {
	return &sqlConnConverter{conn: conn}
}

// SQLStmtToDBStmter converts sql.Stmt to an struct that implement DBStmter interface
func SQLStmtToDBStmter(stmt *sql.Stmt) DBStmter {
	return &sqlStmtConverter{stmt: stmt}
}

// SQLTxToDBTxer converts sql.Tx to an struct that implement DBTxer interface
func SQLTxToDBTxer(tx *sql.Tx) DBTxer {
	return &sqlTxConverter{tx: tx}
}

// Conn convert the real implementation from Conn in database/sql in DBer.Conn
func (c *sqlConverter) Conn(ctx context.Context) (DBConner, error) {
	sqlConn, err := c.db.Conn(ctx)

	connlib := &sqlConnConverter{conn: sqlConn}

	return connlib, err
}

// PingContext convert the real implementation from PingContext in database/sql in DBer.PingContext
func (c *sqlConverter) PingContext(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

// Ping convert the real implementation from Ping in database/sql in DBer.Ping
func (c *sqlConverter) Ping() error {
	return c.db.Ping()
}

// Stats convert the real implementation from Stats in database/sql in DBer.Stats
func (c *sqlConverter) Stats() sql.DBStats {
	return c.db.Stats()
}

// Prepare convert the real implementation from Prepare in database/sql in DBer.Prepare
func (c *sqlConverter) Prepare(query string) (DBStmter, error) {
	stmt, err := c.db.Prepare(query)

	stmtlib := &sqlStmtConverter{stmt: stmt}

	return stmtlib, err
}

// PrepareContext convert the real implementation from PrepareContext in database/sql in DBer.PrepareContext
func (c *sqlConverter) PrepareContext(ctx context.Context, query string) (DBStmter, error) {
	stmt, err := c.db.PrepareContext(ctx, query)

	stmtlib := &sqlStmtConverter{stmt: stmt}

	return stmtlib, err
}

// BeginTx convert the real implementation from BeginTx in database/sql in DBer.BeginTx
func (c *sqlConverter) BeginTx(ctx context.Context, opts *sql.TxOptions) (DBTxer, error) {
	tx, err := c.db.BeginTx(ctx, opts)

	txlib := &sqlTxConverter{tx: tx}

	return txlib, err
}

// Query convert the real implementation from Query in database/sql/stmt in DBStmter.Query
func (c *sqlStmtConverter) Query(args ...interface{}) (DBRowser, error) {
	rows, err := c.stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rows, nil
}

// QueryContext convert the real implementation from QueryContext in database/sql/stmt in DBStmter.QueryContext
func (c *sqlStmtConverter) QueryContext(ctx context.Context, args ...interface{}) (DBRowser, error) {
	rows, err := c.stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rows, err
}

// Exec convert the real implementation from Exec in database/sql/stmt in DBStmter.Exec
func (c *sqlStmtConverter) Exec(args ...interface{}) (sql.Result, error) {
	return c.stmt.Exec(args...)
}

// ExecContext convert the real implementation from ExecContext in database/sql/stmt in DBStmter.ExecContext
func (c *sqlStmtConverter) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return c.stmt.ExecContext(ctx, args...)
}

// PrepareContext convert the real implementation from PrepareContext in database/sql/tx in DBTxer.PrepareContext
func (c *sqlTxConverter) PrepareContext(ctx context.Context, query string) (DBStmter, error) {
	stmt, err := c.tx.PrepareContext(ctx, query)

	stmtlib := &sqlStmtConverter{stmt: stmt}

	return stmtlib, err
}

func (c *sqlStmtConverter) Close() error {
	return c.stmt.Close()
}

func (c *sqlConnConverter) Close() error {
	return c.conn.Close()
}

func (c *sqlTxConverter) Commit() error {
	return c.tx.Commit()
}

func (c *sqlTxConverter) Rollback() error {
	return c.tx.Rollback()
}
