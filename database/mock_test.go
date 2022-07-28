package database_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/database"
	"github.com/stretchr/testify/assert"
)

const (
	selectStmt string = "SELECT 1 AS one;"
	updateStmt string = "UPDATE one SET two = ? WHERE three = ?;"
)

func Test_Mock_Database_BeginWithNoErrorMocked_ShouldProcess_WithoutError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockService.PatchBegin(dbc, dbc, nil)

	outDBC, err := mockService.Begin(dbc)

	// Then
	assertions.Nil(err)
	assertions.NotNil(outDBC)
}

func Test_Mock_Database_BeginWithNoErrorMockedCalledTwice_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockService.PatchBegin(dbc, dbc, nil)

	mockService.Begin(dbc)

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.Begin(dbc: %v)", dbc),
		func() { mockService.Begin(dbc) })
}

func Test_Mock_Database_BeginWithErrorMocked_ShouldProcess_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockedError := errors.New("test error")
	mockService.PatchBegin(dbc, nil, mockedError)

	outDBC, err := mockService.Begin(dbc)

	// Then
	assertions.NotNil(err)
	assertions.EqualError(mockedError, err.Error())
	assertions.Nil(outDBC)
}

func Test_Mock_Database_BeginWithoutMock_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.Begin(dbc: %v)", dbc),
		func() { mockService.Begin(dbc) })
}

//

func Test_Mock_Database_CommitWithNoErrorMocked_ShouldProcess_WithoutError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockService.PatchCommit(dbc, nil)

	err := mockService.Commit(dbc)

	// Then
	assertions.Nil(err)
}

func Test_Mock_Database_CommitWithNoErrorMockedCalledTwice_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockService.PatchCommit(dbc, nil)

	mockService.Commit(dbc)

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.Commit(dbc: %v)", dbc),
		func() { mockService.Commit(dbc) })
}

func Test_Mock_Database_CommitWithErrorMocked_ShouldProcess_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockedError := errors.New("test error")
	mockService.PatchCommit(dbc, mockedError)

	err := mockService.Commit(dbc)

	// Then
	assertions.NotNil(err)
	assertions.EqualError(mockedError, err.Error())
}

func Test_Mock_Database_CommitWithoutMock_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.Commit(dbc: %v)", dbc),
		func() { mockService.Commit(dbc) })
}

//

func Test_Mock_Database_RollbackWithNoErrorMocked_ShouldProcess_WithoutError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockService.PatchRollback(dbc, nil)

	err := mockService.Rollback(dbc)

	// Then
	assertions.Nil(err)
}

func Test_Mock_Database_RollbackWithNoErrorMockedCalledTwice_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockService.PatchRollback(dbc, nil)

	mockService.Rollback(dbc)

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.Rollback(dbc: %v)", dbc),
		func() { mockService.Rollback(dbc) })
}

func Test_Mock_Database_RollbackWithErrorMocked_ShouldProcess_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	mockedError := errors.New("test error")
	mockService.PatchRollback(dbc, mockedError)

	err := mockService.Rollback(dbc)

	// Then
	assertions.NotNil(err)
	assertions.EqualError(mockedError, err.Error())
}

func Test_Mock_Database_RollbackWithoutMock_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.Rollback(dbc: %v)", dbc),
		func() { mockService.Rollback(dbc) })
}

//

func Test_Mock_Database_WithTransactionWithNoErrorMocked_ShouldProcess_WithoutError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When

	txFn := func(dbc *database.DBContext) error { return nil }
	mockService.PatchWithTransaction(txFn, nil)

	err := mockService.WithTransaction(txFn)

	// Then
	assertions.Nil(err)
}

func Test_Mock_Database_WithTransactionWithNoErrorMockedCalledTwice_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	txFn := func(dbc *database.DBContext) error { return nil }
	mockService.PatchWithTransaction(txFn, nil)

	mockService.WithTransaction(txFn)

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.WithTransaction(txFn: ...)"),
		func() { mockService.WithTransaction(txFn) })
}

func Test_Mock_Database_WithTransactionWithErrorMocked_ShouldProcess_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	txFn := func(txFn *database.DBContext) error { return nil }
	mockedError := errors.New("test error")
	mockService.PatchWithTransaction(txFn, mockedError)

	err := mockService.WithTransaction(txFn)

	// Then
	assertions.NotNil(err)
	assertions.EqualError(mockedError, err.Error())
}

func Test_Mock_Database_WithTransactionWithoutMock_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	txFn := func(txFn *database.DBContext) error { return nil }

	// Then
	assertions.PanicsWithValue(fmt.Sprintf("Mock not available for Database.WithTransaction(txFn: ...)"),
		func() { mockService.WithTransaction(txFn) })
}

//

func Test_Mock_Database_SelectWithNoErrorMockedWithoutParams_ShouldProcess_WithoutError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When

	// without params
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}
	mockedDBR := database.ParseMockDBResultFromJSON(`[{"one":1}]`)

	mockService.PatchSelect(dbc, query, forUpdate, params, mockedDBR, nil)

	dbrWithoutParams, err := mockService.Select(dbc, query, forUpdate, params...)

	// Then

	// without params
	assertions.Nil(err)
	assertions.NotNil(dbrWithoutParams)
	assertions.Len(dbrWithoutParams.GetRows(), 1)
	val, _ := dbrWithoutParams.GetRows()[0].GetFloat64ByNameRequired("one")
	field := int64(val)
	assertions.Equal(int64(1), field)

	// When

	// with params
	query = "SELECT ? AS one, ? AS two;"
	params = []interface{}{"1", uint64(2)}

	arrRowsMap := make([]map[string]interface{}, 0)
	columnMap := make(map[string]interface{})
	columnMap["one"] = "1"
	columnMap["two"] = uint64(2)
	arrRowsMap = append(arrRowsMap, columnMap)
	mockedDBR = database.ParseMockDBResultFromArrRowsMap(arrRowsMap)
	mockService.PatchSelect(dbc, query, forUpdate, params, mockedDBR, nil)

	dbrWithParams, err := mockService.Select(dbc, query, forUpdate, params...)

	// Then

	// with params
	assertions.Nil(err)
	assertions.NotNil(dbrWithParams)
	assertions.Len(dbrWithParams.GetRows(), 1)
	valString, _ := dbrWithParams.GetRows()[0].GetStringByNameRequired("one")
	assertions.Equal("1", valString)
	valInt, _ := dbrWithParams.GetRows()[0].GetUInt64ByNameRequired("two")
	assertions.Equal(uint64(2), valInt)
}

func Test_Mock_Database_SelectWithNoErrorMockedCalledTwice_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}
	mockedDBR := database.ParseMockDBResultFromJSON(`[{"one":1}]`)

	mockService.PatchSelect(dbc, query, forUpdate, params, mockedDBR, nil)

	mockService.Select(dbc, query, forUpdate, params...)

	// Then
	assertions.PanicsWithValue("Mock not available for Database.Select(dbc: &{<nil> 0 <nil> <nil>}, query: "+
		"SELECT 1 AS one;, forUpdate: false, params: [])",
		func() { mockService.Select(dbc, query, forUpdate, params...) })
}

func Test_Mock_Database_SelectWithErrorMocked_ShouldProcess_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}
	mockedError := errors.New("test error")
	mockService.PatchSelect(dbc, query, forUpdate, params, nil, mockedError)

	dbr, err := mockService.Select(dbc, query, forUpdate, params...)

	// Then
	assertions.Nil(dbr)
	assertions.NotNil(err)
	assertions.EqualError(mockedError, err.Error())
}

func Test_Mock_Database_SelectWithoutMock_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}

	// Then
	assertions.PanicsWithValue("Mock not available for Database.Select(dbc: &{<nil> 0 <nil> <nil>}, query: "+
		"SELECT 1 AS one;, forUpdate: false, params: [])",
		func() { mockService.Select(dbc, query, forUpdate, params...) })
}

//

func Test_Mock_Database_ExecuteWithNoErrorMockedWithoutParams_ShouldProcess_WithoutError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := updateStmt
	params := []interface{}{"2", 3}
	mockedDBR := database.ParseMockDBResult(int64(1), int64(0))

	mockService.PatchExecute(dbc, query, params, mockedDBR, nil)

	dbr, err := mockService.Execute(dbc, query, params...)

	// Then
	assertions.NotNil(dbr)
	assertions.Equal(int64(1), dbr.AffectedRows())
	assertions.Nil(err)
}

func Test_Mock_Database_ExecuteWithNoErrorMockedCalledTwice_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := updateStmt
	params := []interface{}{"2", 3}
	mockedDBR := database.ParseMockDBResult(int64(1), int64(0))

	mockService.PatchExecute(dbc, query, params, mockedDBR, nil)

	mockService.Execute(dbc, query, params...)

	// Then
	assertions.PanicsWithValue("Mock not available for Database.Execute(dbc: &{<nil> 0 <nil> <nil>}, query: "+
		"UPDATE one SET two = ? WHERE three = ?;, params: [2 3])",
		func() { mockService.Execute(dbc, query, params...) })
}

func Test_Mock_Database_ExecuteWithErrorMocked_ShouldProcess_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := updateStmt
	params := []interface{}{"2", 3}
	mockedError := errors.New("test error")

	mockService.PatchExecute(dbc, query, params, nil, mockedError)

	dbr, err := mockService.Execute(dbc, query, params...)

	// Then
	assertions.Nil(dbr)
	assertions.NotNil(err)
	assertions.EqualError(mockedError, err.Error())
}

func Test_Mock_Database_ExecuteWithoutMock_ShouldPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := updateStmt
	params := []interface{}{"2", 3}

	// Then
	assertions.PanicsWithValue("Mock not available for Database.Execute(dbc: &{<nil> 0 <nil> <nil>}, query: "+
		"UPDATE one SET two = ? WHERE three = ?;, params: [[2 3]])",
		func() { mockService.Execute(dbc, query, params) })
}

func Test_Mock_SelectUniqueValue_WithoutError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}
	mockedDBR := database.ParseMockDBResultFromJSON(`[{"one":"1"}]`).GetRows()[0]

	mockService.PatchSelectUniqueValue(dbc, query, forUpdate, params, &mockedDBR, nil)
	dbRow, err := mockService.SelectUniqueValue(dbc, query, forUpdate, params...)
	one, nre := dbRow.GetInt64ByNameRequired("one")
	assertions.Nil(nre)

	// Then
	assertions.NotNil(dbRow)
	assertions.Equal(int64(1), one)
	assertions.Nil(err)
}

func Test_Mock_SelectUniqueValue_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}

	mockService.PatchSelectUniqueValue(dbc, query, forUpdate, params, nil, errors.New("forced for test"))
	dbRow, err := mockService.SelectUniqueValue(dbc, query, forUpdate, params...)

	// Then
	assertions.Nil(dbRow)
	assertions.NotNil(err)
	assertions.EqualError(err, "forced for test")
}

func Test_Mock_SelectUniqueValue_WithErrorFirstThenSuccess(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}
	// Error
	mockService.PatchSelectUniqueValue(dbc, query, forUpdate, params, nil, errors.New("forced for test"))
	// Success
	mockedDBR := database.ParseMockDBResultFromJSON(`[{"one":"1"}]`).GetRows()[0]
	mockService.PatchSelectUniqueValue(dbc, query, forUpdate, params, &mockedDBR, nil)
	dbRow1, err1 := mockService.SelectUniqueValue(dbc, query, forUpdate, params...)
	dbRow2, err2 := mockService.SelectUniqueValue(dbc, query, forUpdate, params...)
	one, _ := dbRow2.GetInt64ByNameRequired("one")

	// Then
	assertions.Nil(dbRow1)
	assertions.NotNil(err1)
	assertions.EqualError(err1, "forced for test")
	assertions.NotNil(dbRow2)
	assertions.Nil(err2)
	assertions.Equal(int64(1), one)
}

func Test_Mock_SelectUniqueValue_WithPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// Then
	assertions.PanicsWithValue("Mock not available for Database.SelectUniqueValue(dbc: &{<nil> 0 <nil> <nil>}, query: "+
		"SELECT 1 AS one;, forUpdate: false, params: [[]])",
		func() {
			dbc := &database.DBContext{}
			query := selectStmt
			forUpdate := false
			params := []interface{}{}
			mockService.SelectUniqueValue(dbc, query, forUpdate, params)
		})
}

func Test_Mock_SelectUniqueValue_ReturnsTwoRows_ForcedForTest_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := selectStmt
	forUpdate := false
	params := []interface{}{}
	mockService.PatchSelectUniqueValue(dbc, query, forUpdate, params, nil,
		errors.New("unexpected records size, expected 1 but was: 2"))
	dbRow, err := mockService.SelectUniqueValue(dbc, query, forUpdate, params...)

	// Then
	assertions.NotNil(err)
	assertions.EqualError(err, "unexpected records size, expected 1 but was: 2")
	assertions.Nil(dbRow)
}

func Test_Mock_ExecuteEnsuringOneAffectedRow_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := updateStmt
	params := []interface{}{"2", 3}

	mockService.PatchExecuteEnsuringOneAffectedRow(dbc, query, params, errors.New("forced for test"))
	err := mockService.ExecuteEnsuringOneAffectedRow(dbc, query, params...)

	// Then
	assertions.NotNil(err)
	assertions.EqualError(err, "forced for test")
}

func Test_Mock_ExecuteEnsuringOneAffectedRow_WithErrorFirstThenSuccess(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := updateStmt
	params := []interface{}{"2", 3}
	// Error
	mockService.PatchExecuteEnsuringOneAffectedRow(dbc, query, params, errors.New("forced for test"))
	// Success
	mockService.PatchExecuteEnsuringOneAffectedRow(dbc, query, params, nil)

	err1 := mockService.ExecuteEnsuringOneAffectedRow(dbc, query, params...)
	err2 := mockService.ExecuteEnsuringOneAffectedRow(dbc, query, params...)

	// Then
	assertions.NotNil(err1)
	assertions.EqualError(err1, "forced for test")
	assertions.Nil(err2)
}

func Test_Mock_ExecuteEnsuringOneAffectedRow_WithPanic(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// Then
	assertions.PanicsWithValue("Mock not available for Database.ExecuteEnsuringOneAffectedRow(dbc: "+
		"&{<nil> 0 <nil> <nil>}, query: UPDATE one SET two = ? WHERE three = ?;, params: [2 3])",
		func() {
			dbc := &database.DBContext{}
			query := updateStmt
			params := []interface{}{"2", 3}
			mockService.ExecuteEnsuringOneAffectedRow(dbc, query, params...)
		})
}

func Test_Mock_ExecuteEnsuringOneAffectedRow_ReturnsTwoRows_ForcedForTest_WithError(t *testing.T) {
	// Given
	assertions, mockService := buildMockDependencies(t)

	// When
	dbc := &database.DBContext{}
	query := updateStmt
	params := []interface{}{"2", 3}
	mockService.PatchExecuteEnsuringOneAffectedRow(dbc, query, params, errors.New("Unable to insert or update: 2"))

	err := mockService.ExecuteEnsuringOneAffectedRow(dbc, query, params...)

	// Then
	assertions.NotNil(err)
	assertions.EqualError(err, "Unable to insert or update: 2")
}

//

func buildMockDependencies(t *testing.T) (*assert.Assertions, *database.Mock) {
	assertions := assert.New(t)
	service := database.NewMock()

	return assertions, service
}
