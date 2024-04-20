package directdebit

import (
	"encoding/json"
	"errors"
	"net/url"
	"reflect"
)

var ErrNonStringField = errors.New("non-string field with 'url' tag encountered")

// processField processes an individual struct field for conversion to URL values.
func processField(field reflect.Value, tag string, values *url.Values) error {
	switch field.Kind() {
	case reflect.String:
		values.Add(tag, field.String())
	case reflect.Struct:
		jsonBytes, err := json.Marshal(field.Interface())
		if err != nil {
			return err
		}
		encodedString := string(jsonBytes)
		values.Add(tag, encodedString)
	default:
		return ErrNonStringField
	}
	return nil
}

// StructToURLValues converts a struct (or nested struct) into url.Values based on tags.
func StructToURLValues(i interface{}) (url.Values, error) {
	values := url.Values{}
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tag := t.Field(i).Tag.Get("url")
		if tag == "" {
			continue
		}
		if err := processField(v.Field(i), tag, &values); err != nil {
			return nil, err
		}
	}

	return values, nil
}
