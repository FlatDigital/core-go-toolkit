package secrets

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"sync"
)

type (
	// Mock struct for secrets mock
	Mock struct {
		mux sync.Mutex

		patchGetMap map[hash][]outputForGet
	}

	inputForGet struct {
		InputKey string
	}

	outputForGet struct {
		outputValue string
		outputError error
	}

	hash [16]byte
)

// NewMock secrets Mock
func NewMock() *Mock {
	patchGetMap := make(map[hash][]outputForGet)

	secretsMock := &Mock{
		patchGetMap: patchGetMap,
	}
	return secretsMock
}

// PatchGet patch for Get function
func (mock *Mock) PatchGet(inputKey string, outputValue string, outputError error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := getInputForGet(inputKey)
	inputHash := toHash(input)
	output := getOutputForGet(outputValue, outputError)

	if _, exists := mock.patchGetMap[inputHash]; !exists {
		arrOutputForGet := make([]outputForGet, 0)
		mock.patchGetMap[inputHash] = arrOutputForGet
	}
	mock.patchGetMap[inputHash] = append(mock.patchGetMap[inputHash], output)
}

func getInputForGet(inputKey string) inputForGet {
	return inputForGet{
		InputKey: inputKey,
	}
}

func getOutputForGet(outputValue string, outputError error) outputForGet {
	return outputForGet{
		outputValue: outputValue,
		outputError: outputError,
	}
}

func toHash(input interface{}) hash {
	jsonBytes, _ := json.Marshal(input)
	return md5.Sum(jsonBytes)
}

// Get mock for Get function
func (mock *Mock) Get(key string) (string, error) {
	mock.mux.Lock()
	defer mock.mux.Unlock()

	input := getInputForGet(key)
	inputHash := toHash(input)
	arrOutputForGet, exists := mock.patchGetMap[inputHash]
	if !exists || len(arrOutputForGet) == 0 {
		panic(fmt.Sprintf("Mock not available for Secrets.Get(Key: %s)", key))
	}

	output := arrOutputForGet[0]
	arrOutputForGet = arrOutputForGet[1:]
	mock.patchGetMap[inputHash] = arrOutputForGet

	if output.outputError != nil {
		return "", output.outputError
	}

	// return last output error
	return output.outputValue, nil
}
