package sqlmock

import "database/sql"

// SQLRowsMock mock struct
type SQLRowsMock struct {
	patchRowsColumns []outputForRowsColumns
	patchRowsClose   []outputForRowsClose
	patchRowsNext    []outputForRowsNext
	patchRowsScan    map[hash][]outputForRowsScan
}

// NewRowsMockService return mock for database/sql
func NewRowsMockService() *SQLRowsMock {
	sqlRowsMock := SQLRowsMock{
		patchRowsColumns: make([]outputForRowsColumns, 0),
		patchRowsClose:   make([]outputForRowsClose, 0),
		patchRowsNext:    make([]outputForRowsNext, 0),
		patchRowsScan:    map[hash][]outputForRowsScan{},
	}

	return &sqlRowsMock
}

type (
	outputForRowsColumns struct {
		cols        []string
		outputError error
	}

	outputForRowsClose struct {
		outputError error
	}

	outputForRowsNext struct {
		outputRet bool
	}

	inputForRowsScan struct {
		Dest []interface{}
	}

	outputForRowsScan struct {
		outputError error
	}
)

// PatchColumns patches the funcion Columns
func (mock *SQLRowsMock) PatchColumns(cols []string, outputErr error) {
	output := outputForRowsColumns{
		cols:        cols,
		outputError: outputErr,
	}

	mock.patchRowsColumns = append(mock.patchRowsColumns, output)
}

// Columns mocks the real implementation of Columns for the database/sql/rows
func (mock *SQLRowsMock) Columns() ([]string, error) {
	if len(mock.patchRowsColumns) == 0 {
		panic("Mock not available for SQLRowsMock.Columns")
	}

	output := mock.patchRowsColumns[0]
	// dequeue
	mock.patchRowsColumns = mock.patchRowsColumns[1:]

	return output.cols, output.outputError
}

// PatchClose patches the funcion Close
func (mock *SQLRowsMock) PatchClose(outputErr error) {
	output := outputForRowsClose{
		outputError: outputErr,
	}

	mock.patchRowsClose = append(mock.patchRowsClose, output)
}

// Close mocks the real implementation of Close for the database/sql/rows
func (mock *SQLRowsMock) Close() error {
	if len(mock.patchRowsClose) == 0 {
		panic("Mock not available for SQLRowsMock.Close")
	}

	output := mock.patchRowsClose[0]
	// dequeue
	mock.patchRowsClose = mock.patchRowsClose[1:]

	return output.outputError
}

// PatchNext patches the funcion Next
func (mock *SQLRowsMock) PatchNext(outputRet bool) {
	output := outputForRowsNext{
		outputRet: outputRet,
	}

	mock.patchRowsNext = append(mock.patchRowsNext, output)
}

// Next mocks the real implementation of Next for the database/sql/rows
func (mock *SQLRowsMock) Next() bool {
	if len(mock.patchRowsNext) == 0 {
		panic("Mock not available for SQLRowsMock.Next")
	}

	output := mock.patchRowsNext[0]
	// dequeue
	mock.patchRowsNext = mock.patchRowsNext[1:]

	return output.outputRet
}

// PatchScan patches the funcion Scan
func (mock *SQLRowsMock) PatchScan(outputDest []interface{}, outputErr error) {
	input := inputForRowsScan{
		Dest: outputDest,
	}
	hash := toHash(input)

	output := outputForRowsScan{
		outputError: outputErr,
	}

	mock.patchRowsScan[hash] = append(mock.patchRowsScan[hash], output)
}

// Scan mocks the real implementation of Scan for the database/sql/rows
func (mock *SQLRowsMock) Scan(dest ...interface{}) error {
	inputStruct := inputForRowsScan{
		Dest: dest,
	}
	hash := toHash(inputStruct)

	mocksArr, isPresent := mock.patchRowsScan[hash]
	if !isPresent || len(mocksArr) == 0 {
		panic("Mock not available for SQLRowsMock.Scan")
	}

	output := mocksArr[0]
	// dequeue
	mock.patchRowsScan[hash] = mocksArr[1:]

	return output.outputError
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
