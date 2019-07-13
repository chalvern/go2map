package go2map

import (
	"errors"
)

// FlatMap flatMap
type FlatMap map[string]interface{}

// GetStringOf return string value
func (fm FlatMap) GetStringOf(path string) (targetValue string, err error) {
	rawValue := fm[path]
	if rawValue == nil {
		return "", errors.New(path + " not found")
	}

	targetValue, ok := rawValue.(string)
	if !ok {
		return "", errors.New(path + " value is not string")
	}
	return
}
