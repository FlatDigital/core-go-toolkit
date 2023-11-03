package sqlmock

import "github.com/stretchr/testify/mock"

// SQLResultMock mock struct
type SQLResultMock struct {
	mock.Mock
}

// NewResultMockService return mock for database/sql
func NewResultMockService() *SQLResultMock {
	return &SQLResultMock{
		Mock: mock.Mock{},
	}
}

// PatchRowsAffected patches the funcion RowsAffected
func (mock *SQLResultMock) PatchRowsAffected(rowsNum int64, outputErr error) {
	mock.On("RowsAffected").Return(rowsNum, outputErr).Once()
}

// RowsAffected mocks the real implementation of RowsAffected for the database/sql/rows
func (mock *SQLResultMock) RowsAffected() (int64, error) {
	args := mock.Called()
	rowsNum, _ := args.Get(0).(int64)
	err, _ := args.Get(1).(error)
	return rowsNum, err
}

// LastInsertId mocks the real implementation of LastInsertId for the database/sql/result
func (mock *SQLResultMock) LastInsertId() (int64, error) {
	panic("TODO: Implement mock for sql.stmt.LastInsertId")
}
