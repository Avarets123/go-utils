package validator

import (
	"fmt"
	"io"
	"reflect"
	"strings"
)

const validTag string = "valid"
const required string = "required"

type ValidatorMethods interface {
	ValidFromReader(io.Reader, any) ValidatorError
	ValidStruct(data any) ValidatorError
}

type Validator struct {
	errors ValidatorError
}

func New() *Validator {
	return &Validator{
		errors: make(ValidatorError),
	}
}

func (v *Validator) ValidStruct(data any) ValidatorError {

	refV := reflect.ValueOf(data)
	refT := reflect.TypeOf(data)

	for i := range refT.NumField() {

		fieldData := refT.Field(i)

		tagValue := fieldData.Tag.Get(validTag)
		if tagValue == "" {
			continue
		}

		tagValueSplited := strings.Split(tagValue, ",")

		dataType := tagValueSplited[0]

		isSlice := refV.Field(i).Kind() == reflect.Slice
		isDataRequired := false

		if len(tagValueSplited) > 1 {
			isDataRequired = tagValueSplited[1] == required
		}

		isValueZero := refV.Field(i).IsZero()

		fieldName := refT.Field(i).Name

		if isDataRequired && isValueZero {
			errStr := fmt.Sprintf("Field \"%s\" must not be empty!", fieldName)
			v.errors[fieldName] = append(v.errors[fieldName], errStr)
			continue
		}

		if isValueZero {
			continue
		}

		fieldValue := refV.Field(i).Interface()

		if isSlice {
			errStrSlice := v.validArray(fieldValue, dataType)
			if len(errStrSlice) == 0 {
				continue
			}
			v.errors[fieldName] = append(v.errors[fieldName], errStrSlice...)
			continue
		}

		validFn, ok := MapValidFns[dataType]
		if !ok {
			errStr := fmt.Sprintf("Field \"%s\" has unsupproted type %s!", fieldName, dataType)
			v.errors[fieldName] = append(v.errors[fieldName], errStr)
			continue
		}

		err := validFn(fieldValue, dataType)
		if err != nil {
			v.errors[fieldName] = append(v.errors[fieldName], err.Error())
			continue
		}

	}

	return v.errors

}

func (v *Validator) validArray(data any, fieldType string) []string {

	_, fieldType, _ = strings.Cut(fieldType, "[]")

	errs := []string{}

	refData := reflect.ValueOf(data)

	dataSlice := make([]any, refData.Len())

	for i := range refData.Len() {
		dataSlice[i] = refData.Index(i).Interface()

	}

	validFn := MapValidFns[fieldType]

	for i, v := range dataSlice {
		err := validFn(v, fieldType)
		if err != nil {
			errs = append(errs, fmt.Sprintf("Elem[%d]: %+v", i, v))
			continue
		}

	}

	return errs
}
