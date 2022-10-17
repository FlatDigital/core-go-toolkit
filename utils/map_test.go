package utils

import (
	"encoding/gob"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	ass := assert.New(t)

	var contentFloat float64 = 1.1
	var contentFloatZero float64 = 1.1
	var contentInt int64 = 7

	m := make(Map)
	m["bool"] = true
	m["string"] = "string"
	m["float64"] = contentFloat
	m["map"] = make(Map)
	m["map-interface"] = map[string]interface{}{}
	m["nil"] = nil
	m["buffer"] = []byte(`buffer_value`)
	m["buffer-uint8"] = []byte(`buffer_value_uint8`)
	m["int64"] = contentInt

	// empty
	emptyMap, err := m.GetMap("")
	ass.Nil(emptyMap)
	ass.NotNil(err)
	ass.Equal(ErrEmptyKey, err)

	emptyBool, err := m.GetBool("")
	ass.False(emptyBool)
	ass.NotNil(err)
	ass.Equal(ErrEmptyKey, err)

	emptyString, err := m.GetString("")
	ass.Empty(emptyString)
	ass.NotNil(err)
	ass.Equal(ErrEmptyKey, err)

	emptyFloat, err := m.GetFloat64("")
	ass.Zero(emptyFloat)
	ass.NotNil(err)
	ass.Equal(ErrEmptyKey, err)

	// invalid key
	invalidKeyMap, err := m.GetMap("invalid")
	ass.Nil(invalidKeyMap)
	ass.NotNil(err)
	ass.Equal(ErrKeyNotFound, err)

	invalidKeyBool, err := m.GetBool("invalid")
	ass.False(invalidKeyBool)
	ass.NotNil(err)
	ass.Equal(ErrKeyNotFound, err)

	invalidKeyString, err := m.GetString("invalid")
	ass.Empty(invalidKeyString)
	ass.NotNil(err)
	ass.Equal(ErrKeyNotFound, err)

	invalidKeyFloat, err := m.GetFloat64("invalid")
	ass.Zero(invalidKeyFloat)
	ass.NotNil(err)
	ass.Equal(ErrKeyNotFound, err)

	invalidKeyBuffer, err := m.GetBuffer("invalid")
	ass.Zero(invalidKeyBuffer)
	ass.NotNil(err)
	ass.Equal(ErrKeyNotFound, err)

	invalidKeyInt64, err := m.GetInt64("invalid")
	ass.Zero(invalidKeyInt64)
	ass.NotNil(err)
	ass.Equal(ErrKeyNotFound, err)

	// invalid type
	var invalidTypeMap Map
	var invalidTypeBool bool
	var invalidTypeString string
	var invalidTypeFloat float64
	var invalidTypeBuffer []byte
	var invalidTypeInt64 int64

	invalidTypeMap, err = m.GetMap("bool")
	ass.Empty(invalidTypeMap)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeMap, err = m.GetMap("string")
	ass.Empty(invalidTypeMap)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeMap, err = m.GetMap("float64")
	ass.Empty(invalidTypeMap)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)

	invalidTypeBool, err = m.GetBool("map")
	ass.False(invalidTypeBool)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeBool, err = m.GetBool("string")
	ass.False(invalidTypeBool)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeBool, err = m.GetBool("float64")
	ass.False(invalidTypeBool)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)

	invalidTypeString, err = m.GetString("map")
	ass.Empty(invalidTypeString)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeString, err = m.GetString("bool")
	ass.Empty(invalidTypeString)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeString, err = m.GetString("float64")
	ass.Empty(invalidTypeString)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)

	invalidTypeFloat, err = m.GetFloat64("map")
	ass.Zero(invalidTypeFloat)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeFloat, err = m.GetFloat64("bool")
	ass.Zero(invalidTypeFloat)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeFloat, err = m.GetFloat64("string")
	ass.Zero(invalidTypeFloat)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)

	invalidTypeBuffer, err = m.GetBuffer("map")
	ass.Zero(invalidTypeBuffer)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeBuffer, err = m.GetBuffer("bool")
	ass.Zero(invalidTypeBuffer)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeBuffer, err = m.GetBuffer("string")
	ass.Zero(invalidTypeBuffer)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)

	invalidTypeInt64, err = m.GetInt64("map")
	ass.Zero(invalidTypeInt64)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeInt64, err = m.GetInt64("bool")
	ass.Zero(invalidTypeInt64)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)
	invalidTypeInt64, err = m.GetInt64("string")
	ass.Zero(invalidTypeInt64)
	ass.NotNil(err)
	ass.Equal(ErrInvalidType, err)

	// Err nil value
	nilValue, err := m.GetMap("nil")
	ass.NotNil(err)
	ass.Equal(ErrNilValueFound, err)
	ass.Equal(ErrNilValueFound, err)
	var nilValueMap Map
	ass.Equal(nilValueMap, nilValue)

	// success
	validMap, err := m.GetMap("map")
	ass.NotNil(validMap)
	ass.IsType(Map{}, validMap)
	ass.Nil(err)
	validBool, err := m.GetBool("bool")
	ass.True(validBool)
	ass.IsType(true, validBool)
	ass.Nil(err)
	validString, err := m.GetString("string")
	ass.NotEmpty(validString)
	ass.Equal("string", validString)
	ass.IsType("", validString)
	ass.Nil(err)
	validFloat, err := m.GetFloat64("float64")
	ass.NotZero(validFloat)
	ass.Equal(contentFloat, validFloat)
	ass.IsType(contentFloatZero, validFloat)
	ass.Nil(err)
	validBuffer, err := m.GetBuffer("buffer")
	ass.Nil(err)
	ass.IsType([]byte(``), validBuffer)
	ass.Nil(err)
	validInt64, err := m.GetInt64("int64")
	ass.Nil(err)
	ass.IsType(contentInt, validInt64)
	ass.Nil(err)
	validBufferUint8, err := m.GetString("buffer-uint8")
	ass.Nil(err)
	ass.IsType(string(`buffer_value_uint8`), validBufferUint8)
	ass.Nil(err)
	validMapInterface, err := m.GetMap("map-interface")
	ass.Nil(err)
	ass.IsType(Map{}, validMapInterface)
	ass.Nil(err)
}

func TestMapType_GetArrayMap(t *testing.T) {
	ass := assert.New(t)

	m := make(Map)
	arrayMap := make([]Map, 1)
	am := make(Map)
	am["test"] = "test"
	arrayMap[0] = am
	m["key"] = arrayMap

	arrayMapValue, err := m.GetArrayMap("key")
	ass.Nil(err)
	ass.Equal(1, len(arrayMapValue))
	ass.Equal(am, arrayMapValue[0])

	_, err = m.GetArrayMap("")
	ass.NotNil(err)
	ass.Equal(ErrEmptyKey, err)

	_, err = m.GetArrayMap("nf")
	ass.NotNil(err)

	m["key"] = nil
	_, err = m.GetArrayMap("key")
	ass.Nil(err)

	arrayString := make([]string, 0)
	m["key"] = arrayString
	_, err = m.GetArrayMap("key")
	ass.NotNil(err)

	arrayInterface := `{"key":[{"test":"test"},{"test":"test"}]}`
	json.Unmarshal([]byte(arrayInterface), &m)
	arrayMapValue, err = m.GetArrayMap("key")
	ass.Nil(err)
	ass.Equal(2, len(arrayMapValue))
	ass.Equal(am, arrayMapValue[0])
}

func TestMap_SetValue(t *testing.T) {
	ass := assert.New(t)

	// Given
	m := make(Map)

	// When
	m.SetValue("key", "value")

	// Then
	val, err := m.GetString("key")
	ass.Nil(err)
	ass.NotNil(val)
	ass.Equal("value", val)
}

func TestMap_Clone(t *testing.T) {
	ass := assert.New(t)

	// Given
	m1 := make(Map)
	m1.SetValue("key1", "value1")
	m1.SetValue("key2", "value2")
	identifications1 := make([]Map, 0)
	identification1 := make(Map)
	identification1.SetValue("buffer", "abc123")
	identifications1 = append(identifications1, identification1)
	m1.SetValue("identifications", identifications1)

	// When
	m2, err := m1.Clone()

	// Then
	ass.Nil(err)
	key1m1, _ := m1.GetString("key1")
	key1m2, _ := m2.GetString("key1")
	ass.Equal(key1m1, key1m2)
	key2m1, _ := m1.GetString("key2")
	key2m2, _ := m2.GetString("key2")
	ass.Equal(key2m1, key2m2)

	m1.SetValue("key1", "none")
	value, _ := m2.GetString("key1")
	ass.Equal(value, "value1")

	identificationsM2, _ := m2.GetArrayMap("identifications")
	for _, identification := range identificationsM2 {
		buffer, _ := identification.GetString("buffer")
		ass.Equal("abc123", buffer)
		identification.SetValue("buffer", "def345")
		buffer, _ = identification.GetString("buffer")
		ass.Equal("def345", buffer)
	}

	identificationsM1, _ := m1.GetArrayMap("identifications")
	for _, identification := range identificationsM1 {
		buffer, _ := identification.GetString("buffer")
		ass.Equal("abc123", buffer)
	}
}

func TestMap_Clone_Fail(t *testing.T) {
	ass := assert.New(t)

	// Given
	gob.Register([]interface{}{})
	m1 := make(Map)
	m1.SetValue("key1", "value1")
	m1.SetValue("key2", "value2")
	identifications1 := make([]Map, 0)
	identification1 := make(Map)
	identification1.SetValue("buffer", time.Now())
	identifications1 = append(identifications1, identification1)
	m1.SetValue("identifications", identifications1)

	// When
	_, err := m1.Clone()

	// Then
	ass.NotNil(err)
	ass.Equal("gob: type not registered for interface: time.Time", err.Error())
}

func TestEntityMapType_GetInt(t *testing.T) {
	ass := assert.New(t)
	var m Map

	idExample := `{"key":434}`
	json.Unmarshal([]byte(idExample), &m)
	var key int = 434
	id, _ := m.GetInt("key")
	ass.Equal(key, id)
	idError, err := m.GetInt("not_valid_key")
	ass.Equal(int(0), idError)
	ass.NotNil(err)
	ass.Equal("key not found", err.Error())

	idErrorEmpty, err := m.GetInt("")
	ass.Equal(int(0), idErrorEmpty)
	ass.NotNil(err)
	ass.Equal("empty key", err.Error())
}

func TestByteArrayToMap_SingleMap(t *testing.T) {
	ass := assert.New(t)

	bytes := []byte(`{"key":434}`)

	m, err := ByteArrayToMap(bytes)
	ass.Nil(err)

	value, _ := (*m).getRawValue("key")
	var singleMap float64 = 434
	ass.Equal(value, singleMap)
}

func TestByteArrayToMap_ArrayMap(t *testing.T) {
	ass := assert.New(t)

	bytes := []byte(`[{"key":434},{"key2":565}]`)

	m, err := ByteArrayToMap(bytes)
	ass.Nil(err)

	value, _ := (*m).GetArrayMap("root")
	ass.IsType([]Map{}, value)

	ass.Equal(len(value), 2)
	var singleMap float64 = 434
	ass.Equal(value[0].GetFloat("key"), singleMap)
	singleMap = 565
	ass.Equal(value[1].GetFloat("key2"), singleMap)
}

func TestByteArrayToMap_Error(t *testing.T) {
	ass := assert.New(t)

	bytes := []byte(`{"key`)

	_, err := ByteArrayToMap(bytes)
	ass.NotNil(err)
	ass.Equal("unexpected end of JSON input", err.Error())
}
