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
	Int64 struct {
		Value int64
		Valid bool
		Set   bool
	}

	Uint64 struct {
		Value uint64
		Valid bool
		Set   bool
	}

	Float64 struct {
		Value float64
		Valid bool
		Set   bool
	}

	String struct {
		Value string
		Valid bool
		Set   bool
	}

	Bool struct {
		Value bool
		Valid bool
		Set   bool
	}

	Date struct {
		Value string
		Valid bool
		Set   bool
	}

	DateTime struct {
		Value time.Time
		Valid bool
		Set   bool
	}

	StringArray struct {
		Value []string
		Valid bool
		Set   bool
	}

	Uint64Array struct {
		Value []uint64
		Valid bool
		Set   bool
	}

	MapStringInterface struct {
		Value map[string]interface{}
		Valid bool
		Set   bool
	}
)

func (i *Int64) UnmarshalJSON(data []byte) error {
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

func (i *Int64) GetForInt64Pointer(to *int64) *int64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *Int64) GetForInt64(to int64) int64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return 0
	}
	return to
}

func (i *Uint64) UnmarshalJSON(data []byte) error {
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

func (i *Uint64) GetForUint64Pointer(to *uint64) *uint64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *Uint64) GetForUint64(to uint64) uint64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return 0
	}
	return to
}

func (i *Float64) UnmarshalJSON(data []byte) error {
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

func (i *Float64) GetForFloat64Pointer(to *float64) *float64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *Float64) GetForFloat64(to float64) float64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return 0
	}
	return to
}

func (i *String) UnmarshalJSON(data []byte) error {
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

func (i *String) GetForStringPointer(to *string) *string {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *String) GetForString(to string) string {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return ""
	}
	return to
}

func (i *Bool) UnmarshalJSON(data []byte) error {
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

func (i *Bool) GetForBoolPointer(to *bool) *bool {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *Bool) GetForBool(to bool) bool {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return false
	}
	return to
}

func (i *Date) UnmarshalJSON(data []byte) error {
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

func (i *Date) GetForDatePointer(to *string) *string {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *Date) GetForDate(to string) string {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return ""
	}
	return to
}

func (i *DateTime) UnmarshalJSON(data []byte) error {
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
	theTime, err := time.Parse(YYYYMMDDHHMMSS, temp)
	if err != nil {
		return err
	}

	// Apply expected time format
	temp = theTime.Format(YYYYMMDDHHMMSS)
	i.Value = theTime
	i.Valid = true

	return nil
}

func (i *DateTime) GetForDateTimePointer(to *time.Time) *time.Time {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *DateTime) GetForDateTime(to time.Time) time.Time {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return time.Time{}
	}
	return to
}

func (i *StringArray) UnmarshalJSON(data []byte) error {
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

func (i *StringArray) GetForStringArrayPointer(to *[]string) *[]string {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *StringArray) GetForStringArray(to []string) []string {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return []string{}
	}
	return to
}

func (i *Uint64Array) UnmarshalJSON(data []byte) error {
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

func (i *Uint64Array) GetForStringArrayPointer(to *[]uint64) *[]uint64 {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *Uint64Array) GetForStringArray(to []uint64) []uint64 {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return []uint64{}
	}
	return to
}

func (i *MapStringInterface) UnmarshalJSON(data []byte) error {
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

func (i *MapStringInterface) GetForStringArrayPointer(to *map[string]interface{}) *map[string]interface{} {
	if i.Set {
		if i.Valid {
			value := i.Value
			return &value
		}
		return nil
	}
	return to
}

func (i *MapStringInterface) GetForStringArray(to map[string]interface{}) map[string]interface{} {
	if i.Set {
		if i.Valid {
			return i.Value
		}
		return map[string]interface{}{}
	}
	return to
}
