package database

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

// Mock struct for database mock
type Mock struct {
	service

	patchBeginMap                         map[hash][]outputForBegin
	patchCommitMap                        map[hash][]outputForCommit
	patchRollbackMap                      map[hash][]outputForRollback
	patchSelectMap                        map[hash][]outputForSelect
	patchSelectUniqueValueMap             map[hash][]outputForSelectUniqueValue
	patchSelectUniqueValueNonEmptyMap     map[hash][]outputForSelectUniqueValueNonEmpty
	patchExecuteMap                       map[hash][]outputForExecute
	patchExecuteEnsuringOneAffectedRowMap map[hash][]outputForExecuteEnsuringOneAffectedRow
}

//

// ParseMockDBResultAffectedRows parse a mocked dbresult with affected rows
func ParseMockDBResultAffectedRows(affectedRows int64) *DBResult {
	dbr := DBResult{
		affectedRows: affectedRows,
	}

	// done
	return &dbr
}

// ParseMockDBResultFromJSON parse a mocked dbresult from a json
func ParseMockDBResultFromJSON(jsonDBR string) *DBResult {
	arrByte := []byte(jsonDBR)
	arrRowsMap := make([]map[string]interface{}, 0)

	err := json.Unmarshal(arrByte, &arrRowsMap)
	if err != nil {
		panic(fmt.Sprintf("Cannot unmarshal mock dbresult: %v", jsonDBR))
	}

	dbr := ParseMockDBResultFromArrRowsMap(arrRowsMap)

	// done
	return dbr
}

// ParseMockDBResultFromArrRowsMap parse a mocked dbresult from an array of maps
func ParseMockDBResultFromArrRowsMap(arrRowsMap []map[string]interface{}) *DBResult {
	rows := make(DBRowArray, 0)
	// iterate over rows
	for _, item := range arrRowsMap {
		columns := make(DBColumns)
		// iterate over columns
		for key, value := range item {
			column := DBColumn{
				name:  key,
				field: value,
			}
			columns[key] = column
		}

		// add row
		row := DBRow{
			columns: columns,
		}
		rows = append(rows, row)
	}

	// mocked dbresult
	dbr := DBResult{
		rows: DBRows{
			DBRowArray: rows,
		},
	}

	// done
	return &dbr
}

//

// NewMock returns a Mock struct for database mock functions
func NewMock() *Mock {
	patchBeginMap := make(map[hash][]outputForBegin)
	patchCommitMap := make(map[hash][]outputForCommit)
	patchRollbackMap := make(map[hash][]outputForRollback)
	patchSelectMap := make(map[hash][]outputForSelect)
	patchSelectUniqueValueMap := make(map[hash][]outputForSelectUniqueValue)
	patchSelectUniqueValueNonEmptyMap := make(map[hash][]outputForSelectUniqueValueNonEmpty)
	patchExecuteMap := make(map[hash][]outputForExecute)
	patchExecuteEnsuringOneAffectedRowMap := make(map[hash][]outputForExecuteEnsuringOneAffectedRow)
	databaseMock := &Mock{
		patchBeginMap:                         patchBeginMap,
		patchCommitMap:                        patchCommitMap,
		patchRollbackMap:                      patchRollbackMap,
		patchSelectMap:                        patchSelectMap,
		patchSelectUniqueValueMap:             patchSelectUniqueValueMap,
		patchSelectUniqueValueNonEmptyMap:     patchSelectUniqueValueNonEmptyMap,
		patchExecuteMap:                       patchExecuteMap,
		patchExecuteEnsuringOneAffectedRowMap: patchExecuteEnsuringOneAffectedRowMap,
	}

	// done
	return databaseMock
}

//

type inputForBegin struct {
	DBC *DBContext
}

type outputForBegin struct {
	dbc *DBContext
	err error
}

type inputForCommit struct {
	DBC *DBContext
}

type outputForCommit struct {
	err error
}

type inputForRollback struct {
	DBC *DBContext
}

type outputForRollback struct {
	err error
}

type inputForSelect struct {
	DBC       *DBContext
	Query     string
	ForUpdate bool
	ArrParams []interface{}
}

type outputForSelect struct {
	dbr *DBResult
	err error
}

type inputForSelectUniqueValue struct {
	DBC       *DBContext
	Query     string
	ForUpdate bool
	ArrParams []interface{}
}

type outputForSelectUniqueValue struct {
	dbr *DBRow
	err error
}

type inputForSelectUniqueValueNonEmpty struct {
	DBC       *DBContext
	Query     string
	ForUpdate bool
	ArrParams []interface{}
}

type outputForSelectUniqueValueNonEmpty struct {
	dbr *DBRow
	err error
}

type inputForExecute struct {
	DBC       *DBContext
	Query     string
	ArrParams []interface{}
}

type outputForExecute struct {
	dbr *DBResult
	err error
}

type inputForExecuteEnsuringOneAffectedRow struct {
	DBC       *DBContext
	Query     string
	ArrParams []interface{}
}

type outputForExecuteEnsuringOneAffectedRow struct {
	err error
}

type hash [16]byte

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}

// PatchBegin patch for Begin function
func (mock *Mock) PatchBegin(inputDBC *DBContext, outputDBC *DBContext, outputError error) {
	input := getInputForBegin(inputDBC)
	inputHash := toHash(input)
	output := getOutputForBegin(outputDBC, outputError)

	if _, exists := mock.patchBeginMap[inputHash]; !exists {
		arrOutputForBegin := make([]outputForBegin, 0)
		mock.patchBeginMap[inputHash] = arrOutputForBegin
	}
	mock.patchBeginMap[inputHash] = append(mock.patchBeginMap[inputHash], output)
}

func getInputForBegin(dbc *DBContext) inputForBegin {
	return inputForBegin{
		DBC: dbc,
	}
}

func getOutputForBegin(dbc *DBContext, err error) outputForBegin {
	return outputForBegin{
		dbc: dbc,
		err: err,
	}
}

// PatchCommit patch for Commit function
func (mock *Mock) PatchCommit(inputDBC *DBContext, outputError error) {
	input := getInputForCommit(inputDBC)
	inputHash := toHash(input)
	output := getOutputForCommit(outputError)

	if _, exists := mock.patchCommitMap[inputHash]; !exists {
		arrOutputForCommit := make([]outputForCommit, 0)
		mock.patchCommitMap[inputHash] = arrOutputForCommit
	}
	mock.patchCommitMap[inputHash] = append(mock.patchCommitMap[inputHash], output)
}

func getInputForCommit(dbc *DBContext) inputForCommit {
	return inputForCommit{
		DBC: dbc,
	}
}

func getOutputForCommit(err error) outputForCommit {
	return outputForCommit{
		err: err,
	}
}

// PatchRollback patch for Rollback function
func (mock *Mock) PatchRollback(inputDBC *DBContext, outputError error) {
	input := getInputForRollback(inputDBC)
	inputHash := toHash(input)
	output := getOutputForRollback(outputError)

	if _, exists := mock.patchRollbackMap[inputHash]; !exists {
		arrOutputForRollback := make([]outputForRollback, 0)
		mock.patchRollbackMap[inputHash] = arrOutputForRollback
	}
	mock.patchRollbackMap[inputHash] = append(mock.patchRollbackMap[inputHash], output)
}

func getInputForRollback(dbc *DBContext) inputForRollback {
	return inputForRollback{
		DBC: dbc,
	}
}

func getOutputForRollback(err error) outputForRollback {
	return outputForRollback{
		err: err,
	}
}

// PatchWithTransaction patch for WithTransaction function
func (mock *Mock) PatchWithTransaction(txFn func(dbc *DBContext) error, outputError error) {
	panic("To patch Database.WithTransaction patch Database.Begin, Database.Rollback and Database.Commit")
}

// PatchSelect patch for Select function
func (mock *Mock) PatchSelect(inputDBC *DBContext, inputQuery string, inputForUpdate bool,
	inputArrParams []interface{}, outputDBResult *DBResult, outputError error) {
	input := getInputForSelect(inputDBC, inputQuery, inputForUpdate, inputArrParams...)
	inputHash := toHash(input)
	output := getOutputForSelect(outputDBResult, outputError)

	if _, exists := mock.patchSelectMap[inputHash]; !exists {
		arrOutputForSelect := make([]outputForSelect, 0)
		mock.patchSelectMap[inputHash] = arrOutputForSelect
	}
	mock.patchSelectMap[inputHash] = append(mock.patchSelectMap[inputHash], output)
}

func getInputForSelect(dbc *DBContext, query string, forUpdate bool, arrParams ...interface{}) inputForSelect {
	return inputForSelect{
		DBC:       dbc,
		Query:     query,
		ForUpdate: forUpdate,
		ArrParams: arrParams,
	}
}

func getOutputForSelect(dbr *DBResult, err error) outputForSelect {
	return outputForSelect{
		dbr: dbr,
		err: err,
	}
}

// PatchExecute patch for Execute function
func (mock *Mock) PatchExecute(inputDBC *DBContext, inputQuery string, inputArrParams []interface{},
	outputDBResult *DBResult, outputError error) {
	input := getInputForExecute(inputDBC, inputQuery, inputArrParams...)
	inputHash := toHash(input)
	output := getOutputForExecute(outputDBResult, outputError)

	if _, exists := mock.patchExecuteMap[inputHash]; !exists {
		arrOutputForExecute := make([]outputForExecute, 0)
		mock.patchExecuteMap[inputHash] = arrOutputForExecute
	}
	mock.patchExecuteMap[inputHash] = append(mock.patchExecuteMap[inputHash], output)
}

func getInputForExecute(dbc *DBContext, query string, arrParams ...interface{}) inputForExecute {
	return inputForExecute{
		DBC:       dbc,
		Query:     query,
		ArrParams: arrParams,
	}
}

func getOutputForExecute(dbr *DBResult, err error) outputForExecute {
	return outputForExecute{
		dbr: dbr,
		err: err,
	}
}

// Begin mock for Begin function
func (mock *Mock) Begin(inDbc *DBContext) (outDbc *DBContext, err error) {
	input := getInputForBegin(inDbc)
	inputHash := toHash(input)
	arrOutputForBegin, exists := mock.patchBeginMap[inputHash]
	if !exists || len(arrOutputForBegin) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.Begin(dbc: %v)", inDbc))
	}

	output := arrOutputForBegin[0]
	arrOutputForBegin = arrOutputForBegin[1:]
	mock.patchBeginMap[inputHash] = arrOutputForBegin

	if output.err != nil {
		return nil, output.err
	}

	// done
	return output.dbc, nil
}

// Commit mock for Commit function
func (mock *Mock) Commit(dbc *DBContext) error {
	input := getInputForCommit(dbc)
	inputHash := toHash(input)
	arrOutputForCommit, exists := mock.patchCommitMap[inputHash]
	if !exists || len(arrOutputForCommit) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.Commit(dbc: %v)", dbc))
	}

	output := arrOutputForCommit[0]
	arrOutputForCommit = arrOutputForCommit[1:]
	mock.patchCommitMap[inputHash] = arrOutputForCommit

	if output.err != nil {
		return output.err
	}

	// done
	return nil
}

// Rollback mock for rollback
func (mock *Mock) Rollback(dbc *DBContext) error {
	input := getInputForRollback(dbc)
	inputHash := toHash(input)
	arrOutputForRollback, exists := mock.patchRollbackMap[inputHash]
	if !exists || len(arrOutputForRollback) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.Rollback(dbc: %v)", dbc))
	}

	output := arrOutputForRollback[0]
	arrOutputForRollback = arrOutputForRollback[1:]
	mock.patchRollbackMap[inputHash] = arrOutputForRollback

	if output.err != nil {
		return output.err
	}

	// done
	return nil
}

// WithTransaction mock for WithTransaction function
func (mock *Mock) WithTransaction(txFn func(dbc *DBContext) error) (err error) {
	txContext, err := mock.Begin(nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			// Rollbacks the transaction if there's a panic
			_ = mock.Rollback(txContext)
			panic(p)
		} else if err != nil {
			// Rollbacks the transaction if the txFn returns an error
			rollbackErr := mock.Rollback(txContext)
			if rollbackErr != nil {
				err = fmt.Errorf("error rollbacking transaction: %s, %s", rollbackErr, err)
			}
		} else {
			// Commits the transaction
			err = mock.Commit(txContext)
		}
	}()

	// Executes txFn passed as parameter
	err = txFn(txContext)
	return err
}

// Select mock for select
func (mock *Mock) Select(dbc *DBContext, query string, forUpdate bool, params ...interface{}) (*DBResult, error) {
	input := getInputForSelect(dbc, query, forUpdate, params...)
	inputHash := toHash(input)
	arrOutputForSelect, exists := mock.patchSelectMap[inputHash]
	if !exists || len(arrOutputForSelect) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.Select(dbc: %v, query: %s, forUpdate: %t, params: %v)",
			dbc, query, forUpdate, params))
	}

	output := arrOutputForSelect[0]
	arrOutputForSelect = arrOutputForSelect[1:]
	mock.patchSelectMap[inputHash] = arrOutputForSelect

	if output.err != nil {
		return nil, output.err
	}

	// done
	return output.dbr, nil
}

// Execute mock for execute
func (mock *Mock) Execute(dbc *DBContext, query string, params ...interface{}) (*DBResult, error) {
	input := getInputForExecute(dbc, query, params...)
	inputHash := toHash(input)
	arrOutputForExecute, exists := mock.patchExecuteMap[inputHash]
	if !exists || len(arrOutputForExecute) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.Execute(dbc: %v, query: %s, params: %v)", dbc, query, params))
	}

	output := arrOutputForExecute[0]
	arrOutputForExecute = arrOutputForExecute[1:]
	mock.patchExecuteMap[inputHash] = arrOutputForExecute

	if output.err != nil {
		return nil, output.err
	}

	// done
	return output.dbr, nil
}

// SelectUniqueValue

// PatchSelectUniqueValue patch for SelectUniqueValue function
func (mock *Mock) PatchSelectUniqueValue(inputDBC *DBContext, inputQuery string, inputForUpdate bool,
	inputArrParams []interface{}, outputDBRow *DBRow, outputError error) {
	input := getInputForSelectUniqueValue(inputDBC, inputQuery, inputForUpdate, inputArrParams)
	inputHash := toHash(input)
	output := getOutputForSelectUniqueValue(outputDBRow, outputError)

	if _, exists := mock.patchSelectUniqueValueMap[inputHash]; !exists {
		arrOutputForSelectUniqueValue := make([]outputForSelectUniqueValue, 0)
		mock.patchSelectUniqueValueMap[inputHash] = arrOutputForSelectUniqueValue
	}
	mock.patchSelectUniqueValueMap[inputHash] = append(mock.patchSelectUniqueValueMap[inputHash], output)
}

func getInputForSelectUniqueValue(dbc *DBContext, query string, forUpdate bool,
	arrParams []interface{}) inputForSelectUniqueValue {
	return inputForSelectUniqueValue{
		DBC:       dbc,
		Query:     query,
		ForUpdate: forUpdate,
		ArrParams: arrParams,
	}
}

func getOutputForSelectUniqueValue(dbr *DBRow, err error) outputForSelectUniqueValue {
	return outputForSelectUniqueValue{
		dbr: dbr,
		err: err,
	}
}

// SelectUniqueValue mock for SelectUniqueValue
func (mock *Mock) SelectUniqueValue(dbc *DBContext, query string, forUpdate bool,
	params ...interface{}) (*DBRow, error) {
	input := getInputForSelectUniqueValue(dbc, query, forUpdate, params)
	inputHash := toHash(input)
	arrOutputForSelectUniqueValue, exists := mock.patchSelectUniqueValueMap[inputHash]
	if !exists || len(arrOutputForSelectUniqueValue) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.SelectUniqueValue(dbc: %v, query: %s, forUpdate: %t, params: %v)",
			dbc, query, forUpdate, params))
	}

	output := arrOutputForSelectUniqueValue[0]
	arrOutputForSelectUniqueValue = arrOutputForSelectUniqueValue[1:]
	mock.patchSelectUniqueValueMap[inputHash] = arrOutputForSelectUniqueValue

	if output.err != nil {
		return nil, output.err
	}

	// done
	return output.dbr, nil
}

// SelectUniqueValueNonEmpty

// PatchSelectUniqueValueNonEmpty patch for SelectUniqueValueNonEmpty function
func (mock *Mock) PatchSelectUniqueValueNonEmpty(inputDBC *DBContext, inputQuery string, inputForUpdate bool,
	inputArrParams []interface{}, outputDBRow *DBRow, outputError error) {
	input := getInputForSelectUniqueValueNonEmpty(inputDBC, inputQuery, inputForUpdate, inputArrParams)
	inputHash := toHash(input)
	output := getOutputForSelectUniqueValueNonEmpty(outputDBRow, outputError)

	if _, exists := mock.patchSelectUniqueValueNonEmptyMap[inputHash]; !exists {
		arrOutputForSelectUniqueValueNonEmpty := make([]outputForSelectUniqueValueNonEmpty, 0)
		mock.patchSelectUniqueValueNonEmptyMap[inputHash] = arrOutputForSelectUniqueValueNonEmpty
	}
	mock.patchSelectUniqueValueNonEmptyMap[inputHash] = append(mock.patchSelectUniqueValueNonEmptyMap[inputHash], output)
}

func getInputForSelectUniqueValueNonEmpty(dbc *DBContext, query string, forUpdate bool,
	arrParams []interface{}) inputForSelectUniqueValueNonEmpty {
	return inputForSelectUniqueValueNonEmpty{
		DBC:       dbc,
		Query:     query,
		ForUpdate: forUpdate,
		ArrParams: arrParams,
	}
}

func getOutputForSelectUniqueValueNonEmpty(dbr *DBRow, err error) outputForSelectUniqueValueNonEmpty {
	return outputForSelectUniqueValueNonEmpty{
		dbr: dbr,
		err: err,
	}
}

// SelectUniqueValueNonEmpty mock for SelectUniqueValueNonEmpty
func (mock *Mock) SelectUniqueValueNonEmpty(dbc *DBContext, query string, forUpdate bool,
	params ...interface{}) (*DBRow, error) {
	input := getInputForSelectUniqueValueNonEmpty(dbc, query, forUpdate, params)
	inputHash := toHash(input)
	arrOutputForSelectUniqueValueNonEmpty, exists := mock.patchSelectUniqueValueNonEmptyMap[inputHash]
	if !exists || len(arrOutputForSelectUniqueValueNonEmpty) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.SelectUniqueValueNonEmpty(dbc: %v, query: %s, forUpdate: %t, params: %v)",
			dbc, query, forUpdate, params))
	}

	output := arrOutputForSelectUniqueValueNonEmpty[0]
	arrOutputForSelectUniqueValueNonEmpty = arrOutputForSelectUniqueValueNonEmpty[1:]
	mock.patchSelectUniqueValueNonEmptyMap[inputHash] = arrOutputForSelectUniqueValueNonEmpty

	if output.err != nil {
		return nil, output.err
	}

	// done
	return output.dbr, nil
}

// ExecuteEnsuringOneAffectedRow

// PatchExecuteEnsuringOneAffectedRow patch for ExecuteEnsuringOneAffectedRow function
func (mock *Mock) PatchExecuteEnsuringOneAffectedRow(inputDBC *DBContext, inputQuery string,
	inputArrParams []interface{}, outputError error) {
	input := getInputForExecuteEnsuringOneAffectedRow(inputDBC, inputQuery, inputArrParams)
	inputHash := toHash(input)
	output := getOutputForExecuteEnsuringOneAffectedRow(outputError)

	if _, exists := mock.patchExecuteEnsuringOneAffectedRowMap[inputHash]; !exists {
		outputForExecuteEnsuringOneAffectedRow := make([]outputForExecuteEnsuringOneAffectedRow, 0)
		mock.patchExecuteEnsuringOneAffectedRowMap[inputHash] = outputForExecuteEnsuringOneAffectedRow
	}
	mock.patchExecuteEnsuringOneAffectedRowMap[inputHash] =
		append(mock.patchExecuteEnsuringOneAffectedRowMap[inputHash], output)
}

func getInputForExecuteEnsuringOneAffectedRow(dbc *DBContext, query string,
	arrParams []interface{}) inputForExecuteEnsuringOneAffectedRow {
	return inputForExecuteEnsuringOneAffectedRow{
		DBC:       dbc,
		Query:     query,
		ArrParams: arrParams,
	}
}

func getOutputForExecuteEnsuringOneAffectedRow(err error) outputForExecuteEnsuringOneAffectedRow {
	return outputForExecuteEnsuringOneAffectedRow{
		err: err,
	}
}

// ExecuteEnsuringOneAffectedRow mock for ExecuteEnsuringOneAffectedRow
func (mock *Mock) ExecuteEnsuringOneAffectedRow(dbc *DBContext, query string,
	params ...interface{}) error {
	input := getInputForExecuteEnsuringOneAffectedRow(dbc, query, params)
	inputHash := toHash(input)
	arrOutputForExecuteEnsuringOneAffectedRow, exists := mock.patchExecuteEnsuringOneAffectedRowMap[inputHash]
	if !exists || len(arrOutputForExecuteEnsuringOneAffectedRow) == 0 {
		panic(fmt.Sprintf("Mock not available for Database.ExecuteEnsuringOneAffectedRow(dbc: %v, query: %s, params: %v)",
			dbc, query, params))
	}

	output := arrOutputForExecuteEnsuringOneAffectedRow[0]
	arrOutputForExecuteEnsuringOneAffectedRow = arrOutputForExecuteEnsuringOneAffectedRow[1:]
	mock.patchExecuteEnsuringOneAffectedRowMap[inputHash] = arrOutputForExecuteEnsuringOneAffectedRow

	if output.err != nil {
		return output.err
	}

	// done
	return nil
}
