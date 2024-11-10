package library

import "reflect"

// CopyAndRemoveFieldsByKeys menyalin struct dan membuang field yang sesuai dengan keys yang diberikan.
func CopyAndRemoveFieldsByKeys(src interface{}, keysToRemove []string) interface{} {
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Struct {
		return nil // Pastikan input adalah struct
	}

	// Membuat salinan dari struct dengan tipe yang sama
	dst := reflect.New(srcVal.Type()).Elem()

	// Mengonversi keysToRemove menjadi map untuk pencarian yang lebih cepat
	keysMap := make(map[string]struct{}, len(keysToRemove))
	for _, key := range keysToRemove {
		keysMap[key] = struct{}{}
	}

	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Field(i)
		fieldType := srcVal.Type().Field(i)

		// Hanya menyalin field jika nama field tidak ada dalam keysToRemove
		if _, found := keysMap[fieldType.Name]; !found {
			dst.FieldByName(fieldType.Name).Set(field)
		}
	}

	return dst.Interface()
}
