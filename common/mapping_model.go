package common

import (
	"errors"
	"reflect"
)

func MappingModel(src interface{}, dest interface{}) error {
	srcFields := reflect.ValueOf(src)
	destFields := reflect.ValueOf(dest)

	if srcFields.Kind() != reflect.Ptr || destFields.Kind() != reflect.Ptr {
		return ErrNotPointer
	}

	srcFields = srcFields.Elem()
	destFields = destFields.Elem()

	for i := 0; i < srcFields.NumField(); i++ {
		srcFieldName := srcFields.Type().Field(i).Name
		srcFieldValue := srcFields.Field(i)
		destFieldName := destFields.FieldByName(srcFieldName)

		if !srcFieldValue.IsNil() && destFieldName.IsValid() {
			if srcFieldValue.Kind() == destFieldName.Kind() {
				destFieldName.Set(srcFieldValue)
				continue
			}

			if srcFieldValue.Kind() == reflect.Ptr && destFieldName.Kind() != reflect.Ptr {
				destFieldName.Set(reflect.ValueOf(srcFieldValue.Elem().Interface()))
				continue
			}

			if srcFieldValue.Kind() != reflect.Ptr && destFieldName.Kind() == reflect.Ptr {
				destFieldName.Set(reflect.ValueOf(&srcFieldValue))
				continue
			}

			return ErrCannotMapping
		}
	}

	return nil
}

var (
	ErrNotPointer    = errors.New("not pointer when mapping model")
	ErrCannotMapping = errors.New("cannot mapping model")
)
