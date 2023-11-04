package sqlmock

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

// SQLRowsMock mock struct
type SQLRowsMock struct {
	mock.Mock
}

// NewRowsMockService return mock for database/sql
func NewRowsMockService() *SQLRowsMock {
	return &SQLRowsMock{
		Mock: mock.Mock{},
	}
}

// PatchColumns patches the funcion Columns
func (mock *SQLRowsMock) PatchColumns(cols []string, outputErr error) {
	mock.On("Columns").Return(cols, outputErr).Once()
}

// Columns mocks the real implementation of Columns for the database/sql/rows
func (mock *SQLRowsMock) Columns() ([]string, error) {
	args := mock.Called()
	cols, _ := args.Get(0).([]string)
	err, _ := args.Get(1).(error)
	return cols, err
}

// PatchClose patches the funcion Close
func (mock *SQLRowsMock) PatchClose(outputErr error) {
	mock.On("Close").Return(outputErr).Once()
}

// Close mocks the real implementation of Close for the database/sql/rows
func (mock *SQLRowsMock) Close() error {
	args := mock.Called()
	err, _ := args.Get(0).(error)
	return err
}

// PatchNext patches the funcion Next
func (mock *SQLRowsMock) PatchNext(outputRet bool) {
	mock.On("Next").Return(outputRet).Once()
}

// Next mocks the real implementation of Next for the database/sql/rows
func (mock *SQLRowsMock) Next() bool {
	args := mock.Called()
	ret, _ := args.Get(0).(bool)
	return ret
}

// PatchScan patches the funcion Scan
func (mock *SQLRowsMock) PatchScan(outputDest []interface{}, outputErr error) {
	mock.On("Scan", outputDest).Return(outputErr).Once()
}

// Scan mocks the real implementation of Scan for the database/sql/rows
func (mock *SQLRowsMock) Scan(dest ...interface{}) error {
	args := mock.Called(dest)
	err, _ := args.Get(0).(error)
	return err
}

// ColumnTypes mocks the real implementation of ColumnTypes for the database/sql/rows
func (mock *SQLRowsMock) ColumnTypes() ([]*sql.ColumnType, error) {
	panic("TODO: Implement mock for sql.rows.ColumnTypes")
}

// Err mocks the real implementation of Err for the database/sql/rows
func (mock *SQLRowsMock) Err() error {
	panic("TODO: Implement mock for sql.rows.Err")
}

// NextResultSet mocks the real implementation of NextResultSet for the database/sql/rows
func (mock *SQLRowsMock) NextResultSet() bool {
	panic("TODO: Implement mock for sql.rows.NextResultSet")
}
