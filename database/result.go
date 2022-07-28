package database

// DBResult it's a database result struct
type DBResult struct {
	affectedRows int64
	lastInsertId int64
	rows         DBRows
}

//

// AffectedRows returns the quantity of affected rows
func (dbr *DBResult) AffectedRows() int64 {
	return dbr.affectedRows
}

// LastInsertID returns the last inserted id
func (dbr *DBResult) LastInsertID() int64 {
	return dbr.lastInsertId
}

// GetRows returns an array with the rows of the resultset
func (dbr *DBResult) GetRows() []DBRow {
	return dbr.rows.DBRowArray
}
