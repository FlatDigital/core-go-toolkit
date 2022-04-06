package converter_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/FlatDigital/flat-go-toolkit/src/api/libs/database/converter"
	"github.com/stretchr/testify/assert"
)

func Test_SQLToDBer(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}

	// when
	dber := converter.SQLToDBer(sql)

	// then
	ass.NotNil(dber)
}

func Test_SQLConnToDBConner(t *testing.T) {
	// given
	ass := assert.New(t)

	conn := &sql.Conn{}

	// when
	conner := converter.SQLConnToDBConner(conn)

	// then
	ass.NotNil(conner)
}

func Test_SQLTxToDBTxer(t *testing.T) {
	// given
	ass := assert.New(t)

	tx := &sql.Tx{}

	// when
	txer := converter.SQLTxToDBTxer(tx)

	// then
	ass.NotNil(txer)
}

func Test_SQLStmtToDBStmter(t *testing.T) {
	// given
	ass := assert.New(t)

	stmt := &sql.Stmt{}

	// when
	stmter := converter.SQLStmtToDBStmter(stmt)

	// then
	ass.NotNil(stmter)
}

func Test_Prepare(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}
	dber := converter.SQLToDBer(sql)

	// when
	ass.Panics(func() {
		dber.Prepare("")
	})
}

func Test_PrepareContext(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}
	dber := converter.SQLToDBer(sql)
	ctx := context.Background()

	// when
	ass.Panics(func() {
		dber.PrepareContext(ctx, "")
	})
}

func Test_Query(t *testing.T) {
	// given
	ass := assert.New(t)

	stmt := &sql.Stmt{}
	stmter := converter.SQLStmtToDBStmter(stmt)

	// when
	ass.Panics(func() {
		stmter.Query("")
	})
}

func Test_QueryContext(t *testing.T) {
	// given
	ass := assert.New(t)

	stmt := &sql.Stmt{}
	stmter := converter.SQLStmtToDBStmter(stmt)
	ctx := context.Background()

	// when
	ass.Panics(func() {
		stmter.QueryContext(ctx, "")
	})
}

func Test_Tx_PrepareContext(t *testing.T) {
	// given
	ass := assert.New(t)

	tx := &sql.Tx{}
	txer := converter.SQLTxToDBTxer(tx)
	ctx := context.Background()

	// when
	ass.Panics(func() {
		txer.PrepareContext(ctx, "")
	})
}

func Test_Conn(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}
	dber := converter.SQLToDBer(sql)
	ctx := context.Background()

	// when
	ass.Panics(func() {
		dber.Conn(ctx)
	})
}

func Test_PingContext(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}
	dber := converter.SQLToDBer(sql)
	ctx := context.Background()

	// when
	ass.Panics(func() {
		dber.PingContext(ctx)
	})
}

func Test_Ping(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}
	dber := converter.SQLToDBer(sql)

	// when
	ass.Panics(func() {
		dber.Ping()
	})
}

func Test_Stats(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}
	dber := converter.SQLToDBer(sql)

	// when
	stats := dber.Stats()

	ass.NotNil(stats)
}

func Test_BeginTx(t *testing.T) {
	// given
	ass := assert.New(t)

	sql := &sql.DB{}
	dber := converter.SQLToDBer(sql)
	ctx := context.Background()

	// when
	ass.Panics(func() {
		dber.BeginTx(ctx, nil)
	})
}

func Test_Stmt_Exec(t *testing.T) {
	// given
	ass := assert.New(t)

	stmt := &sql.Stmt{}
	stmter := converter.SQLStmtToDBStmter(stmt)

	// when
	ass.Panics(func() {
		stmter.Exec()
	})
}

func Test_Stmt_ExecContext(t *testing.T) {
	// given
	ass := assert.New(t)

	stmt := &sql.Stmt{}
	stmter := converter.SQLStmtToDBStmter(stmt)
	ctx := context.Background()

	// when
	ass.Panics(func() {
		stmter.ExecContext(ctx)
	})
}

func Test_Stmt_Close(t *testing.T) {
	// given
	ass := assert.New(t)

	stmt := &sql.Stmt{}
	stmter := converter.SQLStmtToDBStmter(stmt)

	// when
	ass.Panics(func() {
		stmter.Close()
	})
}

func Test_Conn_Close(t *testing.T) {
	// given
	ass := assert.New(t)

	conn := &sql.Conn{}
	conner := converter.SQLConnToDBConner(conn)

	// when
	ass.Panics(func() {
		conner.Close()
	})
}

func Test_Tx_Commit(t *testing.T) {
	// given
	ass := assert.New(t)

	tx := &sql.Tx{}
	txer := converter.SQLTxToDBTxer(tx)

	// when
	ass.Panics(func() {
		txer.Commit()
	})
}

func Test_Tx_Rollback(t *testing.T) {
	// given
	ass := assert.New(t)

	tx := &sql.Tx{}
	txer := converter.SQLTxToDBTxer(tx)

	// when
	ass.Panics(func() {
		txer.Rollback()
	})
}
