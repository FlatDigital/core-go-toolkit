package database

import (
	"fmt"
	"strconv"
)

// DBColumn it's a column
type DBColumn struct {
	name  string
	field interface{}
}

// DBColumns it's a map of columns
type DBColumns map[string]DBColumn

//

// NewColumn create a new DBColumn
func NewColumn(name string, field interface{}) *DBColumn {
	return &DBColumn{
		name:  name,
		field: field,
	}
}

// GetColumnName returns column name
func (dbc *DBColumn) GetColumnName() string {
	return dbc.name
}

// GetRawValue returns the value of the field
func (dbc *DBColumn) GetRawValue() interface{} {
	return dbc.field
}

// GetBuffer returns the value of the field in byte array type
func (dbc *DBColumn) GetBuffer() ([]byte, error) {
	if dbc.field == nil {
		return nil, nil
	}
	bytesValue, ok := dbc.field.([]byte)
	if !ok {
		return nil, fmt.Errorf("'%s' invalid type, value '%v'", dbc.name, dbc.field)
	}
	return bytesValue, nil
}

// GetInt64 returns the value of the field in int64 type
func (dbc *DBColumn) GetInt64() (*int64, error) {
	if dbc.field == nil {
		return nil, nil
	}
	var int64Value int64
	var err error
	int64Value, ok := dbc.field.(int64)
	if !ok {
		strValue, ok := dbc.field.(string)
		if !ok {
			return nil, fmt.Errorf("'%s' invalid type, value '%v'", dbc.name, dbc.field)
		}
		int64Value, err = strconv.ParseInt(strValue, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("'%s' invalid type, value '%v'", dbc.name, dbc.field)
		}
	}
	return &int64Value, nil
}

// GetUInt64 returns the value of the field in uint64 type
func (dbc *DBColumn) GetUInt64() (*uint64, error) {
	if dbc.field == nil {
		return nil, nil
	}
	var uint64Value uint64
	uint64Value, ok := dbc.field.(uint64)
	if !ok {
		int64Ptr, err := dbc.GetInt64()
		if err != nil {
			return nil, err
		}
		// safe cast using interface{}
		int64Value := *int64Ptr
		if int64Value < 0 {
			return nil, fmt.Errorf("'%s' invalid type, value '%v'", dbc.name, dbc.field)
		}
		uint64Value = uint64(int64Value)
	}
	return &uint64Value, nil
}

// GetFloat64 returns the value of the field in float64 type
func (dbc *DBColumn) GetFloat64() (*float64, error) {
	if dbc.field == nil {
		return nil, nil
	}

	var err error

	var float64Value float64
	float64Value, ok := dbc.field.(float64)
	if !ok {
		var strValue string
		strValue, ok = dbc.field.(string)
		if !ok {
			bytesValue, ok := dbc.field.([]uint8)
			if !ok {
				return nil, fmt.Errorf("'%s' invalid type, value '%v'", dbc.name, dbc.field)
			}
			strValue = string(bytesValue)
		}
		float64Value, err = strconv.ParseFloat(strValue, 64)
		if err != nil {
			return nil, fmt.Errorf("'%s' invalid type, value '%s'", dbc.name, dbc.field)
		}

	}
	return &float64Value, nil
}

// GetString returns the value of the field in string type
func (dbc *DBColumn) GetString() (*string, error) {
	if dbc.field == nil {
		return nil, nil
	}
	stringValue, ok := dbc.field.(string)
	if !ok {
		bufferValue, ok := dbc.field.([]uint8)
		if !ok {
			// fmt.Printf("GETSTRING()    Raw(%v) String(%v) Buffer(%v)", dbc.field, dbc.field.(string), dbc.field.([]uint8))
			return nil, fmt.Errorf("'%s' invalid type, value '%v'", dbc.name, dbc.field)
		}
		stringValue = string(bufferValue)
	}
	return &stringValue, nil
}

// GetBool returns the value of the field in bool type
func (dbc *DBColumn) GetBool() (*bool, error) {
	if dbc.field == nil {
		return nil, nil
	}
	var boolValue bool
	boolValue, ok := dbc.field.(bool)
	if !ok {
		intValue, err := dbc.GetInt64()
		if err != nil {
			return nil, err
		}
		if *intValue != int64(0) {
			boolValue = true
		}
	}
	return &boolValue, nil
}
