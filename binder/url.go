package binder

import (
	"fmt"
	"reflect"
)

func (b *Binder) BindParams(pattern string, v interface{}) error {
	params, err := parseURL(pattern, b.Request.URL.Path)
	if err != nil {
		return err
	}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("out must be a pointer to a struct")
	}
	val = val.Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldVal := val.Field(i)
		tag := field.Tag.Get(structQueryTag)
		if tag == "" || !fieldVal.CanSet() {
			continue
		}

		fieldName, delimiter := parseTag(tag)
		queryValue, ok := params[fieldName]
		if !ok {
			continue
		}
		if queryValue == "" {
			continue
		}

		if err := setFieldValue(fieldVal, queryValue, delimiter); err != nil {
			return fmt.Errorf("error setting field %s: %v", field.Name, err)
		}
	}

	return nil
}
