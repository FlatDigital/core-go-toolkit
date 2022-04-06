package database

import (
	"fmt"
	"reflect"

	// package imported for mysql driver for database/sql
	_ "github.com/go-sql-driver/mysql"
)

// DBRow it's a row
type DBRow struct {
	columns DBColumns
}

// DBRowArray it's an array of row
type DBRowArray []DBRow

// DBRows contains the result of a query
type DBRows struct {
	DBRowArray
}

//

// NewRow create a new row
func NewRow(columns DBColumns) *DBRow {
	return &DBRow{
		columns: columns,
	}
}

// Equals returns if both db rows contain
func (dbr *DBRow) Equals(row *DBRow) bool {
	return reflect.DeepEqual(dbr.columns, row.columns)
}

// GetColumnByName retrieves a column by it's name
func (dbr *DBRow) GetColumnByName(name string) (*DBColumn, error) {
	dbc, ok := dbr.columns[name]
	if !ok {
		return nil, fmt.Errorf("'%s' key not found", name)
	}
	return &dbc, nil
}

// GetBufferByName retrieves a column by it's name
func (dbr *DBRow) GetBufferByName(name string) ([]byte, error) {
	dbc, err := dbr.GetColumnByName(name)
	if err != nil {
		return nil, err
	}
	val, err := dbc.GetBuffer()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetBufferByNameRequired retrieves a non-empty column by it's name
func (dbr *DBRow) GetBufferByNameRequired(name string) ([]byte, error) {
	val, err := dbr.GetBufferByName(name)
	if err != nil {
		return []byte{}, err
	}
	if len(val) == 0 {
		return []byte{}, fmt.Errorf("'%s' required field returned empty value", name)
	}
	return val, nil
}

// GetInt64ByName retrieves a column by it's name
func (dbr *DBRow) GetInt64ByName(name string) (*int64, error) {
	dbc, err := dbr.GetColumnByName(name)
	if err != nil {
		return nil, err
	}
	val, err := dbc.GetInt64()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetInt64ByNameRequired retrieves a non-empty column by it's name
func (dbr *DBRow) GetInt64ByNameRequired(name string) (int64, error) {
	val, err := dbr.GetInt64ByName(name)
	if err != nil {
		return 0, err
	}
	if val == nil { // Here we don't compare against zero-value
		return 0, fmt.Errorf("'%s' required field returned empty value", name)
	}
	return *val, nil
}

// GetUInt64ByName retrieves a column by it's name
func (dbr *DBRow) GetUInt64ByName(name string) (*uint64, error) {
	dbc, err := dbr.GetColumnByName(name)
	if err != nil {
		return nil, err
	}
	val, err := dbc.GetUInt64()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetUInt64ByNameRequired retrieves a non-empty column by it's name
func (dbr *DBRow) GetUInt64ByNameRequired(name string) (uint64, error) {
	val, err := dbr.GetUInt64ByName(name)
	if err != nil {
		return 0, err
	}
	if val == nil { // Here we don't compare against zero-value
		return 0, fmt.Errorf("'%s' required field returned empty value", name)
	}
	return *val, nil
}

// GetFloat64ByName retrieves a column by it's name
func (dbr *DBRow) GetFloat64ByName(name string) (*float64, error) {
	dbc, err := dbr.GetColumnByName(name)
	if err != nil {
		return nil, err
	}
	val, err := dbc.GetFloat64()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetFloat64ByNameRequired retrieves a non-empty column by it's name
func (dbr *DBRow) GetFloat64ByNameRequired(name string) (float64, error) {
	val, err := dbr.GetFloat64ByName(name)
	if err != nil {
		return 0, err
	}
	if val == nil { // Here we don't compare against zero-value
		return 0, fmt.Errorf("'%s' required field returned empty value", name)
	}
	return *val, nil
}

// GetStringByName retrieves a column by it's name
func (dbr *DBRow) GetStringByName(name string) (*string, error) {
	dbc, err := dbr.GetColumnByName(name)
	if err != nil {
		return nil, err
	}
	val, err := dbc.GetString()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetStringByNameRequired retrieves a non-empty column by it's name
func (dbr *DBRow) GetStringByNameRequired(name string) (string, error) {
	val, err := dbr.GetStringByName(name)
	if err != nil {
		return "", err
	}
	if val == nil || *val == "" {
		return "", fmt.Errorf("'%s' required field returned empty value", name)
	}
	return *val, nil
}

// GetBoolByName retrieves a column by it's name
func (dbr *DBRow) GetBoolByName(name string) (*bool, error) {
	dbc, err := dbr.GetColumnByName(name)
	if err != nil {
		return nil, err
	}
	val, err := dbc.GetBool()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// GetBoolByNameRequired retrieves a non-empty column by it's name
func (dbr *DBRow) GetBoolByNameRequired(name string) (bool, error) {
	val, err := dbr.GetBoolByName(name)
	if err != nil {
		return false, err
	}
	if val == nil { // Here we don't compare against zero-value
		return false, fmt.Errorf("'%s' required field returned empty value", name)
	}
	return *val, nil
}
