package go2map

import (
	"errors"
)

// FlatMap flatMap
type FlatMap map[string]interface{}

// MustGetStringOf path's value always returned
func (fm FlatMap) MustGetStringOf(path string) (targetValue string) {
	targetValue, err := fm.GetStringOf(path)
	if err != nil {
		targetValue = err.Error()
	}
	return
}

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

// GetIntOf return string value
func (fm FlatMap) GetIntOf(path string) (targetValue int, err error) {
	rawValue := fm[path]
	if rawValue == nil {
		return 0, errors.New(path + " not found")
	}

	targetValue, ok := rawValue.(int)
	if !ok {
		return 0, errors.New(path + " value is not int")
	}
	return
}
