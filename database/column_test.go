package database_test

import (
	"strconv"
	"testing"

	"github.com/FlatDigital/core-go-toolkit/v2/database"
	"github.com/stretchr/testify/assert"
)

const (
	col1 string = "col1"
	col2 string = "col2"
	col3 string = "col3"

	valueString string = "valueString"
)

func Test_GetColumnName_Success(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := valueString

	column := database.NewColumn(name, value)

	// when
	columnName := column.GetColumnName()

	// then
	ass.NotNil(columnName)
	ass.Equal(name, columnName)
}

func Test_GetColumnName_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	column := database.DBColumn{}

	// when
	columnName := column.GetColumnName()

	// then
	ass.Empty(columnName)
}

func Test_GetRawValue_Success(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := valueString

	column := database.NewColumn(name, value)

	// when
	raw := column.GetRawValue()

	// then
	ass.NotNil(raw)
	ass.Equal(value, raw)
}

func Test_GetBuffer_Success(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := []byte(valueString)

	column := database.NewColumn(name, value)

	// when
	bytes, err := column.GetBuffer()

	// then
	ass.NotNil(bytes)
	ass.Nil(err)
	ass.Equal(value, bytes)
}

func Test_GetBuffer_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	column := database.DBColumn{}

	// when
	bytes, err := column.GetBuffer()

	// then
	ass.Nil(bytes)
	ass.Nil(err)
}

func Test_GetBuffer_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := valueString

	column := database.NewColumn(name, value)

	// when
	bytes, err := column.GetBuffer()

	// then
	ass.Nil(bytes)
	ass.NotNil(err)
}

func Test_GetIn64_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	column := database.DBColumn{}

	// when
	result, err := column.GetInt64()

	// then
	ass.Nil(result)
	ass.Nil(err)
}

func Test_GetInt64_String_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := true

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetInt64()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}

func Test_GetInt64_ParseInt_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := valueString

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetInt64()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}

func Test_GetUIn64_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	column := database.DBColumn{}

	// when
	result, err := column.GetUInt64()

	// then
	ass.Nil(result)
	ass.Nil(err)
}

func Test_GetUInt64_String_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := true

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetUInt64()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}

func Test_GetUInt64_Less_Zero_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := -1

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetUInt64()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}

func Test_GetUInt64_With_Int64(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := int64(999999999999999999)

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetUInt64()

	// then
	ass.NotNil(result)
	ass.Nil(err)
}

func Test_GetFloat64_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	column := database.DBColumn{}

	// when
	result, err := column.GetFloat64()

	// then
	ass.Nil(result)
	ass.Nil(err)
}

func Test_GetFloat64_String_Success(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := "123.9"
	valueFloat, _ := strconv.ParseFloat(value, 64)

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetFloat64()

	// then
	ass.NotNil(result)
	ass.Nil(err)
	ass.Equal(valueFloat, *result)
}

func Test_GetFloat64_String_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := true

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetFloat64()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}

func Test_GetFloat64_ParseInt_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := valueString

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetFloat64()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}

func Test_String_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	column := database.DBColumn{}

	// when
	result, err := column.GetString()

	// then
	ass.Nil(result)
	ass.Nil(err)
}

func Test_GetString_UInt(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := []uint8(valueString)

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetString()

	// then
	ass.NotNil(result)
	ass.Nil(err)
	ass.Equal(value, []uint8(*result))
}

func Test_GetString_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := true

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetString()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}

func Test_GetBool_Success(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := int64(1)

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetBool()

	var resultInt int64
	if *result {
		resultInt = 1
	}

	// then
	ass.NotNil(result)
	ass.Nil(err)
	ass.Equal(value, resultInt)
}

func Test_GetBool_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	column := database.DBColumn{}

	// when
	result, err := column.GetBool()

	// then
	ass.Nil(result)
	ass.Nil(err)
}

func Test_GetBool_Error(t *testing.T) {
	// given
	ass := assert.New(t)
	name := col1
	value := valueString

	column := database.NewColumn(name, value)

	// when
	result, err := column.GetBool()

	// then
	ass.Nil(result)
	ass.NotNil(err)
}
