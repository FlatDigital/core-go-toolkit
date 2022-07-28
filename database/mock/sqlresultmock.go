package sqlmock

// SQLResultMock mock struct
type SQLResultMock struct {
	patchResultRowsAffected []outputForResultRowsAffected
	patchResultLastInsertId []outputForResultLastInsertId
}

// NewResultMockService return mock for database/sql
func NewResultMockService() *SQLResultMock {
	sqlResultMock := SQLResultMock{
		patchResultRowsAffected: make([]outputForResultRowsAffected, 0),
	}

	return &sqlResultMock
}

type (
	outputForResultRowsAffected struct {
		rowsNum     int64
		outputError error
	}

	outputForResultLastInsertId struct {
		lastInsertId int64
		outputError  error
	}
)

// PatchRowsAffected patches the function RowsAffected
func (mock *SQLResultMock) PatchRowsAffected(rowsNum int64, outputErr error) {
	output := outputForResultRowsAffected{
		rowsNum:     rowsNum,
		outputError: outputErr,
	}

	mock.patchResultRowsAffected = append(mock.patchResultRowsAffected, output)
}

// PatchLastInsertId patches the function PatchLastInsertId
func (mock *SQLResultMock) PatchLastInsertId(lastInsertId int64, outputErr error) {
	output := outputForResultLastInsertId{
		lastInsertId: lastInsertId,
		outputError:  outputErr,
	}

	mock.patchResultLastInsertId = append(mock.patchResultLastInsertId, output)
}

// RowsAffected mocks the real implementation of RowsAffected for the database/sql/rows
func (mock *SQLResultMock) RowsAffected() (int64, error) {
	if len(mock.patchResultRowsAffected) == 0 {
		panic("Mock not available for SQLResultMock.RowsAffected")
	}

	output := mock.patchResultRowsAffected[0]
	// dequeue
	mock.patchResultRowsAffected = mock.patchResultRowsAffected[1:]

	return output.rowsNum, output.outputError
}

// LastInsertId mocks the real implementation of LastInsertId for the database/sql/rows
func (mock *SQLResultMock) LastInsertId() (int64, error) {
	if len(mock.patchResultLastInsertId) == 0 {
		panic("Mock not available for SQLResultMock.LastInsertId")
	}

	output := mock.patchResultLastInsertId[0]
	// dequeue
	mock.patchResultLastInsertId = mock.patchResultLastInsertId[1:]

	return output.lastInsertId, output.outputError
}
