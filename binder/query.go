package binder

import (
	"fmt"
	"reflect"
)

func (b *Binder) BindQuery(v interface{}) error {
	params := b.Request.URL.Query()

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
		queryValue := params.Get(fieldName)
		if queryValue == "" {
			continue
		}

		if err := setFieldValue(fieldVal, queryValue, delimiter); err != nil {
			return fmt.Errorf("error setting field %s: %v", field.Name, err)
		}
	}

	return nil
}
