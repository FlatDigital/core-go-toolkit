package secrets

import (
	"errors"
	"os"
)

// Secrets secrets service interface
type Secrets interface {
	// Get returns a secret
	Get(string) (string, error)
}

var (
	// ErrEmptyKey empty key error
	ErrEmptyKey = errors.New("empty key")

	// ErrEmptyValue empty value error
	ErrEmptyValue = errors.New("empty value")
)

// NewService returns a kvs service interface
func NewService() Secrets {
	return &service{}
}

// service struct
type service struct {
}

// Get returns a secret
// Return ErrEmptyKey if the key == ""
// Return ErrEmptyValue if the value == ""
func (service *service) Get(key string) (string, error) {
	if err := validateKey(key); err != nil {
		return "", err
	}

	secret := os.Getenv(key)
	if err := validateValue(secret); err != nil {
		return "", err
	}

	// done
	return secret, nil
}

//

// validateKey return ErrEmptyKey if the key passed by argument is ""
func validateKey(key string) error {
	if key == "" {
		return ErrEmptyKey
	}

	// done
	return nil
}

// validateValue return ErrEmptyValue if the value is ""
func validateValue(value string) error {
	if value == "" {
		return ErrEmptyValue
	}

	// done
	return nil
}
