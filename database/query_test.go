package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	insertStmtIncomingTrans string = `INSERT INTO incoming_transaction VALUES `
	selectStmtVersion       string = `SELECT * FROM version WHERE version = ?`
	selectStmtIncomingTrans string = "SELECT t.id as transaction_id, it.operation_id " +
		"FROM incoming_transaction it " +
		"INNER JOIN transaction t ON t.operation_id = it.operation_id " +
		"WHERE it.operation_id IN (-1000);"
	endChar          string = `;`
	palceholder1000         = "-1000"
	valuePlaceholder        = `(?,?)`
)

func Test_QueryBeginEndSingleValueCompleteSuccess(t *testing.T) {
	ass := assert.New(t)
	begin := insertStmtIncomingTrans
	end := endChar
	value := valuePlaceholder
	params := []interface{}{10, 15}
	q, err := NewQueryBeginEnd(begin, end, value)
	ass.NotNil(q)
	ass.Nil(err)
	err = q.AddParams(params)
	ass.Nil(err)
	statement, err := q.Statement()
	ass.Equal(begin+value+end, statement)
	ass.Nil(err)
	outputParams, err := q.Params()
	ass.Equal(params, outputParams)
	ass.Nil(err)
	_, _, err = q.GetStatementAndParams()
	ass.Nil(err)
}

func Test_NewQueryBeginEndMustWorkSuccess(t *testing.T) {
	ass := assert.New(t)
	begin := insertStmtIncomingTrans
	end := endChar
	value := valuePlaceholder
	q := NewQueryBeginEndMustWork(begin, end, value)
	ass.NotNil(q)
}

func Test_NewQueryBeginEndMustWorkPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the build of the query does not returned a panic")
		}
	}()

	// We force a panic
	begin := insertStmtIncomingTrans
	end := endChar
	value := `(now())` // We sent any ? in order to get test panicking
	NewQueryBeginEndMustWork(begin, end, value)
}

func Test_QueryPlaceholderSingleValueCompleteSuccess(t *testing.T) {
	ass := assert.New(t)
	base := selectStmtIncomingTrans
	placeholder := palceholder1000
	value := "?"
	params := []interface{}{10}
	q, err := NewQueryPlaceholder(base, placeholder, value)
	ass.NotNil(q)
	ass.Nil(err)
	err = q.AddParams(params)
	ass.Nil(err)
	q.ForUpdate(true)
	statement, err := q.Statement()
	ass.Equal("SELECT t.id as transaction_id, it.operation_id FROM incoming_transaction it "+
		"INNER JOIN transaction t ON t.operation_id = it.operation_id "+
		"WHERE it.operation_id IN (?) FOR UPDATE;", statement)
	ass.Nil(err)
	outputParams, err := q.Params()
	ass.Equal(params, outputParams)
	ass.Nil(err)
}

func Test_NewQueryPlaceholderMustWorkSuccess(t *testing.T) {
	ass := assert.New(t)
	base := selectStmtIncomingTrans
	placeholder := palceholder1000
	value := "?"
	q := NewQueryPlaceholderMustWork(base, placeholder, value)
	ass.NotNil(q)
}

func Test_NewQueryPlaceholderMustWorkPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the build of the query does not returned a panic")
		}
	}()

	// We force a panic
	base := selectStmtIncomingTrans
	placeholder := palceholder1000
	value := "" // we send an empty value to get it panicking
	NewQueryPlaceholderMustWork(base, placeholder, value)
}

func Test_QueryPlaceholderWithoutValues(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryPlaceholder(`SELECT
		t.id as transaction_id,
		it.operation_id
	  FROM incoming_transaction it
		INNER JOIN transaction t ON t.operation_id = it.operation_id
	  WHERE
		it.operation_id IN (-1000)
	  ;`, palceholder1000, "")
	ass.Nil(q)
	ass.Error(err)
	ass.EqualError(err, "unable to find any ? in the value placeholder")
}

func Test_QueryPlaceholderAppearsMultipleTimes(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryPlaceholder(`SELECT
		t.id as transaction_id,
		it.operation_id,
		-1000 as error
	  FROM incoming_transaction it
		INNER JOIN transaction t ON t.operation_id = it.operation_id
	  WHERE
		it.operation_id IN (-1000)
	  ;`, palceholder1000, "?")
	ass.Nil(q)
	ass.Error(err)
	ass.EqualError(err, "the placeholder appears 2 times into base query (only 1 is allowed)")
}

func Test_QueryBeginEndValuesDifferentLength(t *testing.T) {
	ass := assert.New(t)
	begin := insertStmtIncomingTrans
	end := endChar
	value := valuePlaceholder
	params := []interface{}{10, 15, 17}
	q, err := NewQueryBeginEnd(begin, end, value)
	ass.NotNil(q)
	ass.Nil(err)
	err = q.AddParams(params)
	ass.Error(err)
	ass.EqualError(err, "the allowed length for this query is 2, but you have sent a 3 params")
}

func Test_QueryBeginEndMultipleParams(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryBeginEnd(insertStmtIncomingTrans, ";", valuePlaceholder)
	ass.NotNil(q)
	ass.Nil(err)
	err = q.AddParams([]interface{}{10, "world"})
	ass.Nil(err)
	err = q.AddParams([]interface{}{"hello", 18})
	ass.Nil(err)
}

func Test_QueryBeginEndWithoutParams(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryBeginEnd(insertStmtIncomingTrans, ";", valuePlaceholder)
	ass.NotNil(q)
	ass.Nil(err)
	statement, err := q.Statement()
	ass.Equal("", statement)
	ass.Error(err)
	ass.EqualError(err, "you haven't added any params to this query")
	_, err = q.Params()
	ass.Error(err)
	ass.EqualError(err, "you haven't added any params to this query")
	_, _, err = q.GetStatementAndParams()
	ass.EqualError(err, "unable to obtain statement or params for query "+
		"(statement error: you haven't added any params to this query / params "+
		"error: you haven't added any params to this query)")
}

func Test_QueryBeginEndWithoutValues(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryBeginEnd(insertStmtIncomingTrans, endChar, ``)
	ass.Nil(q)
	ass.Error(err)
	ass.EqualError(err, "unable to find any ? in the value placeholder")
}

func Test_QueryBeginEndEmptyEnd(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryBeginEnd(insertStmtIncomingTrans, ``, valuePlaceholder)
	ass.NotNil(q)
	ass.Nil(err)
	ass.Equal(";", q.stmtEnd)
}

func Test_QueryBeginEndForUpdate(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryBeginEnd(insertStmtIncomingTrans, endChar, valuePlaceholder)
	ass.NotNil(q)
	ass.Nil(err)
	ass.Equal(false, q.IsForUpdate())
	q.ForUpdate(true)
	ass.Equal(true, q.IsForUpdate())
	q.ForUpdate(false)
	ass.Equal(false, q.IsForUpdate())
}

func Test_QueryBeginEndInvalidQueryType(t *testing.T) {
	ass := assert.New(t)
	q, err := NewQueryBeginEnd(insertStmtIncomingTrans, ";", valuePlaceholder)
	ass.NotNil(q)
	ass.Nil(err)
	err = q.AddParams([]interface{}{10, "world"})
	ass.Nil(err)
	q.queryType = "InvalidQueryType"
	stmt, err := q.Statement()
	ass.Equal("", stmt)
	ass.Error(err)
	ass.EqualError(err, "unimplemented queryType: InvalidQueryType")
}

func Test_NewQueryPlainSuccess(t *testing.T) {
	ass := assert.New(t)
	stmt := selectStmtVersion
	q, err := NewQueryPlain(stmt)
	ass.NotNil(q)
	ass.Nil(err)
}

func Test_NewQueryPlainFail(t *testing.T) {
	ass := assert.New(t)
	stmt := `SELECT * FROM version`
	q, err := NewQueryPlain(stmt)
	ass.Nil(q)
	ass.Error(err)
	ass.EqualError(err, "unable to find any ? in the value placeholder")
}

func Test_QueryPlainFailWithMultipleParams(t *testing.T) {
	ass := assert.New(t)
	stmt := selectStmtVersion
	q, err := NewQueryPlain(stmt)
	ass.NotNil(q)
	ass.Nil(err)
	err = q.AddParams([]interface{}{"1.0.0"})
	ass.Nil(err)
	err = q.AddParams([]interface{}{"1.0.0"})
	ass.Error(err)
	ass.EqualError(err, "unable to add more than a group of params to a plain query")
}

func Test_QueryPlainStatementSuccess(t *testing.T) {
	ass := assert.New(t)
	stmt := selectStmtVersion
	q, err := NewQueryPlain(stmt)
	ass.NotNil(q)
	ass.Nil(err)
	err = q.AddParams([]interface{}{"1.0.0"})
	ass.Nil(err)
	s, err := q.Statement()
	ass.NotEmpty(s)
	ass.Nil(err)
}

func Test_NewQueryPlainMustWorkSuccess(t *testing.T) {
	ass := assert.New(t)
	stmt := selectStmtVersion
	q := NewQueryPlainMustWork(stmt)
	ass.NotNil(q)
}

func Test_NewQueryPlainMustWorkPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("the build of the query does not returned a panic")
		}
	}()

	// We force a panic
	NewQueryPlainMustWork(`SELECT * FROM version`) // We don't send any ? to get test panicking
}
