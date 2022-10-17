package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
)

// Map define a type map with
type Map map[string]interface{}

//

var (
	// ErrEmptyKey empty key
	ErrEmptyKey = errors.New("empty key")
	// ErrKeyNotFound key not found
	ErrKeyNotFound = errors.New("key not found")
	// ErrInvalidType invalid type
	ErrInvalidType = errors.New("invalid type")
	// ErrNilValueFound err nil value found
	ErrNilValueFound = errors.New("nil value found")
)

//

func (m Map) getRawValue(key string) (interface{}, error) {
	if key == "" {
		return nil, ErrEmptyKey
	}
	rawValue, ok := m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	if rawValue == nil {
		return nil, ErrNilValueFound
	}
	return rawValue, nil
}

// GetArrayString returns the value of the field as a array of strings
func (m Map) GetArrayString(key string) ([]string, error) {
	if key == "" {
		return nil, ErrEmptyKey
	}
	rawArray, ok := m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	if rawArray == nil {
		return make([]string, 0), nil
	}

	if arrayString, ok := rawArray.([]string); ok {
		return arrayString, nil
	}

	arrayString, ok := rawArray.([]interface{})
	if !ok {
		return nil, ErrInvalidType
	}
	arrayValue := make([]string, len(arrayString))
	for idx, e := range arrayString {
		s, ok := e.(string)
		if !ok {
			return nil, ErrInvalidType
		}
		arrayValue[idx] = s
	}
	return arrayValue, nil
}

// GetMap returns the value of the field in a map type
func (m Map) GetMap(key string) (Map, error) {
	rawValue, err := m.getRawValue(key)
	if err != nil {
		return nil, err
	}
	mapValue, ok := rawValue.(Map)
	if !ok {
		mapInterfaceValue, ok := rawValue.(map[string]interface{})
		if !ok {
			return nil, ErrInvalidType
		}
		mapValue = Map(mapInterfaceValue)
	}
	return mapValue, nil
}

// GetBool returns the value of the field in boolean type
func (m Map) GetBool(key string) (bool, error) {
	rawValue, err := m.getRawValue(key)
	if err != nil {
		return false, err
	}

	boolValue, ok := rawValue.(bool)
	if !ok {
		return false, ErrInvalidType
	}
	return boolValue, nil
}

// GetString returns the value of the field in string type
func (m Map) GetString(key string) (string, error) {
	rawValue, err := m.getRawValue(key)
	if err != nil {
		return "", err
	}

	stringValue, ok := rawValue.(string)
	if !ok {
		bufferValue, ok := rawValue.([]uint8)
		if !ok {
			return "", ErrInvalidType
		}
		stringValue = string(bufferValue)
	}
	return stringValue, nil
}

// GetFloat64 returns the value of the field in float64 type
func (m Map) GetFloat64(key string) (float64, error) {
	rawValue, err := m.getRawValue(key)
	if err != nil {
		return float64(0), err
	}

	float64Value, ok := rawValue.(float64)
	if !ok {
		return float64(0), ErrInvalidType
	}
	return float64Value, nil
}

// GetInt64 returns the value of the field in uint64 type
func (m Map) GetInt64(key string) (int64, error) {
	rawValue, err := m.getRawValue(key)
	if err != nil {
		return int64(0), err
	}

	int64Value, ok := rawValue.(int64)
	if !ok {
		return int64(0), ErrInvalidType
	}
	return int64Value, nil
}

// GetBuffer returns the value of the field in byte array type
func (m Map) GetBuffer(key string) ([]byte, error) {
	rawValue, err := m.getRawValue(key)
	if err != nil {
		return nil, err
	}

	bytesValue, ok := rawValue.([]byte)
	if !ok {
		return nil, ErrInvalidType
	}
	return bytesValue, nil
}

// GetArrayMap returns the value of the field as a array of map
func (m Map) GetArrayMap(key string) ([]Map, error) {
	if key == "" {
		return nil, ErrEmptyKey
	}
	rawArray, ok := m[key]
	if !ok {
		return nil, ErrKeyNotFound
	}
	if rawArray == nil {
		return make([]Map, 0), nil
	}

	if arrayEntityMap, ok := rawArray.([]Map); ok {
		return arrayEntityMap, nil
	}

	arrayInterface, ok := rawArray.([]interface{})
	if !ok {
		return nil, ErrInvalidType
	}
	arrayValue := make([]Map, len(arrayInterface))
	for idx, e := range arrayInterface {
		em, ok := e.(map[string]interface{})
		if !ok {
			return nil, ErrInvalidType
		}
		arrayValue[idx] = em
	}
	return arrayValue, nil
}

// GetFloat get the value of the field in float type
func (m Map) GetFloat(key string) float64 {
	value, _ := m.getFloat(key)
	return value
}

// Gets a float and a boolean indicating success
func (m Map) getFloat(key string) (float64, error) {
	if key == "" {
		return 0, ErrEmptyKey
	}
	rawValue, ok := m[key]
	if !ok || rawValue == nil {
		return 0, ErrKeyNotFound
	}
	float64Value, ok := rawValue.(float64)
	if !ok {
		return 0, ErrInvalidType
	}
	return float64Value, nil
}

// GetInt returns the value of the field in int type
func (m Map) GetInt(key string) (int, error) {
	floatValue, err := m.getFloat(key)
	if err != nil {
		return int(0), err
	}
	return int(floatValue), nil
}

// SetValue Sets the specified value into the specified key. Key cannot be zero value
// Value can be nil
func (m Map) SetValue(key string, value interface{}) error {
	if key == "" {
		return ErrEmptyKey
	}
	m[key] = value
	return nil
}

// Clone creates a deep copy of the Map
func (m Map) Clone() (Map, error) {
	var buf bytes.Buffer
	gob.Register([]Map{})
	gob.Register(map[string]interface{}{})
	gob.Register([]interface{}{})
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(map[string]interface{}(m))
	if err != nil {
		return nil, err
	}
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}

// DeleteValue removes the specified key/value pair from the map
func (m Map) DeleteValue(key string) error {
	_, ok := m[key]
	if !ok {
		return ErrKeyNotFound
	}
	delete(m, key)
	return nil
}

//

// ByteArrayToMap returns a map from a byte array
func ByteArrayToMap(byteArray []byte) (*Map, error) {
	dataMap := make(Map)
	err := json.Unmarshal(byteArray, &dataMap)
	if err != nil {
		// fallback to array of maps
		arrayDataMap := make([]Map, 0)
		err := json.Unmarshal(byteArray, &arrayDataMap)
		if err != nil {
			return nil, err
		}
		// populate the dataMap with the array in a root_array tag
		_ = dataMap.SetValue("root", arrayDataMap)
	}

	return &dataMap, nil
}
