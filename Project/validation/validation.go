package validation

import (
	"main/library"
	"net/http"
	"reflect"
)

// Fungsi Validation untuk memvalidasi nilai dari field yang ditentukan dalam berbagai tipe data
func Validation(w http.ResponseWriter, target interface{}, keys []string) bool {
	value := reflect.ValueOf(target)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if value.Kind() != reflect.Struct {
		library.StrucToJson(w, library.BadRequest)
		return false
	}

	for _, key := range keys {
		field := value.FieldByName(key)

		if !field.IsValid() {
			library.StrucToJson(w, library.BadRequest)
			return false
		}

		switch field.Kind() {
		case reflect.String:
			if field.String() == "" {
				library.StrucToJson(w, library.BadRequest)
				return false
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() == 0 {
				library.StrucToJson(w, library.BadRequest)
				return false
			}
		case reflect.Float32, reflect.Float64:
			if field.Float() == 0.0 {
				library.StrucToJson(w, library.BadRequest)
				return false
			}
		case reflect.Bool:
			if !field.Bool() {
				library.StrucToJson(w, library.BadRequest)
				return false
			}
		default:
			library.StrucToJson(w, library.BadRequest)
			return false
		}
	}

	return true
}
