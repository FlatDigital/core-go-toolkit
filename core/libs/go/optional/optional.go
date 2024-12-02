package domain

import (
	"encoding/json"
	"time"
)

const (
	// YYYYMMDD Time Layout YYYY-MM-DD: 2022-03-23
	YYYYMMDD = "2006-01-02"

	// YYYYMMDDHHMMSS Time Layout YYYY-MM-DD: 2006-01-02 15:04:05
	YYYYMMDDHHMMSS = "2006-01-02 15:04:05"
)

type (
	OptionalInt64 struct {
		Value int64
		Valid bool
		Set   bool
	}

	OptionalUint64 struct {
		Value uint64
		Valid bool
		Set   bool
	}

	OptionalFloat64 struct {
		Value float64
		Valid bool
		Set   bool
	}

	OptionalString struct {
		Value string
		Valid bool
		Set   bool
	}

	OptionalBool struct {
		Value bool
		Valid bool
		Set   bool
	}

	OptionalDate struct {
		Value string
		Valid bool
		Set   bool
	}

	OptionalStringArray struct {
		Value []string
		Valid bool
		Set   bool
	}

	OptionalUint64Array struct {
		Value []uint64
		Valid bool
		Set   bool
	}

	OptionalMapStringInterface struct {
		Value map[string]interface{}
		Valid bool
		Set   bool
	}
)

func (i *OptionalInt64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp int64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalInt64) GetForInt64Pointer(to *int64) *int64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalInt64) GetForInt64(to int64) int64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return 0
	}
	return to
}

func (i *OptionalUint64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp uint64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalUint64) GetForUint64Pointer(to *uint64) *uint64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalUint64) GetForUint64(to uint64) uint64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return 0
	}
	return to
}

func (i *OptionalFloat64) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp float64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalFloat64) GetForFloat64Pointer(to *float64) *float64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalFloat64) GetForFloat64(to float64) float64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return 0
	}
	return to
}

func (i *OptionalString) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalString) GetForStringPointer(to *string) *string {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalString) GetForString(to string) string {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return ""
	}
	return to
}

func (i *OptionalBool) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp bool
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalBool) GetForBoolPointer(to *bool) *bool {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalBool) GetForBool(to bool) bool {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return false
	}
	return to
}

func (i *OptionalDate) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Verify date format is valid
	theTime, err := time.Parse(YYYYMMDD, temp)
	if err != nil {
		return err
	}

	// Apply expected time format
	temp = theTime.Format(YYYYMMDD)
	i.Value = temp
	i.Valid = true

	return nil
}

func (i *OptionalDate) GetForDatePointer(to *string) *string {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalDate) GetForDate(to string) string {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return ""
	}
	return to
}

func (i *OptionalStringArray) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp []string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalStringArray) GetForStringArrayPointer(to *[]string) *[]string {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalStringArray) GetForStringArray(to []string) []string {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return []string{}
	}
	return to
}

func (i *OptionalUint64Array) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp []uint64
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalUint64Array) GetForStringArrayPointer(to *[]uint64) *[]uint64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalUint64Array) GetForStringArray(to []uint64) []uint64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return []uint64{}
	}
	return to
}

func (i *OptionalMapStringInterface) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp map[string]interface{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

func (i *OptionalMapStringInterface) GetForStringArrayPointer(to *map[string]interface{}) *map[string]interface{} {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *OptionalMapStringInterface) GetForStringArray(to map[string]interface{}) map[string]interface{} {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return map[string]interface{}{}
	}
	return to
}
