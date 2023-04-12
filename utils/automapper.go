package utils

import (
	"fmt"
	"reflect"
)

const TAG = "mapper"

func findTag(field reflect.StructField, tag string) (string, error) {
	t := field.Tag
	if val, ok := t.Lookup(tag); ok {
		return val, nil
	}
	return "", fmt.Errorf("Tag provided does not define a mapper tag")
}

func findFieldByTag(value reflect.Value, tag, tagValue string) (string, error) {
	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		tagVal, err := findTag(field, tag)
		if err != nil {
			return "", err
		}

		if tagVal == tagValue {
			return field.Name, nil
		}
	}

	return "", fmt.Errorf("Src does not contain tag value '%s'", tagValue)
}

func Map(src, dest interface{}) {
	valSrc := reflect.ValueOf(src)
	if valSrc.Kind() == reflect.Pointer {
		valSrc = valSrc.Elem()
	}

	if valSrc.Kind() != reflect.Struct {
		return
	}

	valDest := reflect.ValueOf(dest)
	if valDest.Kind() != reflect.Pointer {
		valDest = valDest.Elem()
	}

	if valDest.Elem().Kind() != reflect.Struct {
		return
	}

	typeSrc := valSrc.Type()
	for i := 0; i < valSrc.NumField(); i++ {
		fieldSrc := typeSrc.Field(i)
		tag, err := findTag(fieldSrc, TAG)
		if err != nil {
			continue
		}

		fieldName, err := findFieldByTag(valDest, TAG, tag)
		if err != nil {
			continue
		}

		// value := valSrc.Field(i).Interface()
		valDest.FieldByName(fieldName)
	}
}
