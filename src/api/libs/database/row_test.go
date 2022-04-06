package database_test

import (
	"testing"

	"github.com/FlatDigital/flat-go-toolkit/src/api/libs/database"
	"github.com/stretchr/testify/assert"
)

func Test_GetColumnByName_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := "val1"
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetColumnByName("fake")

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetBufferByName_Nil(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBufferByName(colName1)

	// then
	ass.NotNil(colRes)
	ass.Nil(err)
}

func Test_GetBufferByName_NoCol_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBufferByName("fake")

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetBufferByName_Type_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBufferByName(colName2)

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetColumnByNameRequired_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBufferByNameRequired(colName1)

	// then
	ass.NotNil(colRes)
	ass.Nil(err)
}

func Test_GetColumnByNameRequired_NoCol_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBufferByNameRequired("fake")

	// then
	ass.NotNil(colRes)
	ass.NotNil(err)
	ass.Equal([]byte{}, colRes)
}

func Test_GetColumnByNameRequired_Empty_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte{}
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBufferByNameRequired(colName1)

	// then
	ass.NotNil(colRes)
	ass.NotNil(err)
	ass.Equal([]byte{}, colRes)
}

func Test_GetInt64ByName_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetInt64ByName("fakecol")

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetInt64ByName_Type_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetInt64ByName(colName1)

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetInt64ByNameRequiered_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetInt64ByNameRequired("fakecol")

	// then
	ass.Equal(int64(0), colRes)
	ass.NotNil(err)
}

func Test_GetUInt64ByName_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetUInt64ByName("fakecol")

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetUInt64ByName_Type_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetUInt64ByName(colName1)

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetUInt64ByNameRequiered_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetUInt64ByNameRequired("fakecol")

	// then
	ass.Equal(uint64(0), colRes)
	ass.NotNil(err)
}

func Test_GetFloat64ByName_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetFloat64ByName("fakecol")

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetFloat64ByName_Type_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetFloat64ByName(colName1)

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetFloat64ByNameRequiered_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetFloat64ByNameRequired("fakecol")

	// then
	ass.Equal(float64(0), colRes)
	ass.NotNil(err)
}

func Test_GetStringByName_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetStringByName("fakecol")

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetStringByName_Type_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := 13123123
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetStringByName(colName1)

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetStringByNameRequiered_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetStringByNameRequired("fakecol")

	// then
	ass.Equal("", colRes)
	ass.NotNil(err)
}

func Test_GetBoolByName_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := int64(1)
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBoolByName(colName1)

	var resultInt int64
	if *colRes {
		resultInt = 1
	}

	// then
	ass.NotNil(colRes)
	ass.Nil(err)
	ass.Equal(colValue1, resultInt)
}

func Test_GetBoolByName_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBoolByName("fakecol")

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetBoolByName_Type_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := 13123123
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBoolByName(colName1)

	// then
	ass.Nil(colRes)
	ass.NotNil(err)
}

func Test_GetBoolByNameRequiered_No_Col_Error(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := []byte("val1")
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBoolByNameRequired("fakecol")

	// then
	ass.Equal(false, colRes)
	ass.NotNil(err)
}

func Test_GetBoolByNameRequired_Success(t *testing.T) {
	// given
	ass := assert.New(t)

	colName1 := col1
	colName2 := col2
	colName3 := col3
	colValue1 := int64(1)
	colValue2 := 2
	colValue3 := float64(3)

	column1 := database.NewColumn(colName1, colValue1)
	column2 := database.NewColumn(colName2, colValue2)
	column3 := database.NewColumn(colName3, colValue3)

	columns := make(database.DBColumns)
	columns[colName1] = *column1
	columns[colName2] = *column2
	columns[colName3] = *column3

	row := database.NewRow(columns)

	// when
	colRes, err := row.GetBoolByNameRequired(colName1)

	var resultInt int64
	if colRes {
		resultInt = 1
	}

	// then
	ass.NotNil(colRes)
	ass.Nil(err)
	ass.Equal(colValue1, resultInt)
}
