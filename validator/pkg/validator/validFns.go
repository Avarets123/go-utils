package validator

import (
	"fmt"
)

type ValidFn func(v any, fieldType string) error

func ValidValueFn[T any](v any, fieldType string) error {

	_, ok := v.(T)
	if ok {
		return nil
	}

	return fmt.Errorf("%+v : is not type %s", v, fieldType)

}

func StrValidFn[T any](v any) error {
	_, ok := v.(T)
	if ok {
		return nil
	}

	return fmt.Errorf("%+v : is not string", v)

}

var MapValidFns map[string]ValidFn = map[string]ValidFn{
	"string": ValidValueFn[string],
	"int":    ValidValueFn[int],
	"float":  ValidValueFn[float64],
}
