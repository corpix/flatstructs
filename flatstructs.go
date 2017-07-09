package flatstructs

import (
	"reflect"
	"strings"
)

var (
	// Default is a Builder with the default parameters.
	Default = NewBuilder("key", "")
)

// Builder is a flat structure builder.
// It could be parameterized with tag name
// to look up struct field names in and
// a keyDelimiter which delimits key parts
// when converting nested struct into flat.
type Builder struct {
	Tag          string
	KeyDelimiter string
}

func (b *Builder) fieldName(field reflect.StructField) string {
	name := field.Tag.Get(b.Tag)
	if name == "" {
		return field.Name
	}
	return name
}

// Keys creates a flat slice of keys from a nested structure exported fields.
func (b *Builder) Keys(v interface{}) ([]string, error) {
	err := checkValue(v)
	if err != nil {
		return nil, err
	}

	return b.toKeys(v, []string{})
}

// toKeys, see Keys().
func (b *Builder) toKeys(v interface{}, prefix []string) ([]string, error) {
	var (
		reflectType  reflect.Type  = indirectType(reflect.TypeOf(v))
		reflectValue reflect.Value = indirectValue(reflect.ValueOf(v))
		field        reflect.StructField
		fieldValue   reflect.Value
		key          string
		keys         []string
		err          error
	)

	if !reflectValue.IsValid() {
		return nil, NewErrInvalid(v)
	}

	err = checkStruct(reflectValue)
	if err != nil {
		return nil, err
	}

	keys = []string{}

	for n := 0; n < reflectValue.NumField(); n++ {
		field = reflectType.Field(n)
		if !isStructFieldExported(field) {
			continue
		}

		fieldValue = indirectValue(reflectValue.Field(n))
		if !fieldValue.CanAddr() {
			continue
		}

		key = b.fieldName(field)
		switch fieldValue.Kind() {
		case reflect.Struct:
			var (
				subKeys []string
			)
			subKeys, err = b.toKeys(
				fieldValue.Addr().Interface(),
				append(prefix, key),
			)
			if err != nil {
				return nil, err
			}

			if len(subKeys) == 0 {
				keys = append(
					keys,
					strings.Join(
						append(prefix, key),
						b.KeyDelimiter,
					),
				)
			} else {
				for _, v := range subKeys {
					keys = append(
						keys,
						v,
					)
				}
			}
		default:
			keys = append(
				keys,
				strings.Join(
					append(prefix, key),
					b.KeyDelimiter,
				),
			)
		}
	}

	return keys, nil
}

// Values creates a flat slice of values from a nested structure exported fields.
func (b *Builder) Values(v interface{}) ([]interface{}, error) {
	err := checkValue(v)
	if err != nil {
		return nil, err
	}

	return b.toValues(v)
}

// toValues, see Values().
func (b *Builder) toValues(v interface{}) ([]interface{}, error) {
	var (
		reflectType  reflect.Type  = indirectType(reflect.TypeOf(v))
		reflectValue reflect.Value = indirectValue(reflect.ValueOf(v))
		field        reflect.StructField
		fieldValue   reflect.Value
		values       []interface{}
		err          error
	)

	if !reflectValue.IsValid() {
		return nil, NewErrInvalid(v)
	}

	err = checkStruct(reflectValue)
	if err != nil {
		return nil, err
	}

	values = []interface{}{}

	for n := 0; n < reflectValue.NumField(); n++ {
		field = reflectType.Field(n)
		if !isStructFieldExported(field) {
			continue
		}

		fieldValue = indirectValue(reflectValue.Field(n))
		if !fieldValue.CanAddr() {
			continue
		}

		if !fieldValue.CanInterface() {
			values = append(
				values,
				nil,
			)
			continue
		}

		switch fieldValue.Kind() {
		case reflect.Struct:
			var (
				subValues []interface{}
			)
			subValues, err = b.toValues(
				fieldValue.Addr().Interface(),
			)
			if err != nil {
				return nil, err
			}

			if len(subValues) == 0 {
				values = append(
					values,
					fieldValue.Interface(),
				)
			} else {
				for _, v := range subValues {
					values = append(
						values,
						v,
					)
				}
			}
		default:
			values = append(
				values,
				fieldValue.Interface(),
			)
		}
	}

	return values, nil
}

// Map creates a map with Keys(): Values() from a nested structure.
func (b *Builder) Map(v interface{}) (map[string]interface{}, error) {
	err := checkValue(v)
	if err != nil {
		return nil, err
	}

	return b.toMap(v)
}

// toMap, see Map().
func (b *Builder) toMap(v interface{}) (map[string]interface{}, error) {
	var (
		result = make(map[string]interface{})
		keys   []string
		values []interface{}
		err    error
	)

	keys, err = b.Keys(v)
	if err != nil {
		return nil, err
	}
	values, err = b.Values(v)
	if err != nil {
		return nil, err
	}

	for k, _ := range keys {
		result[keys[k]] = values[k]
	}

	return result, nil
}

func checkValue(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return NewErrPtrRequired(v)
	}

	return nil
}

func checkStruct(reflectValue reflect.Value) error {
	if reflectValue.Kind() != reflect.Struct {
		return NewErrInvalidKind(
			reflect.Struct,
			reflectValue.Kind(),
		)
	}

	return nil
}

func isStructFieldExported(field reflect.StructField) bool {
	// From reflect docs:
	// PkgPath is the package path that qualifies a lower case (unexported)
	// field name. It is empty for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	return field.PkgPath == ""
}

func indirectValue(reflectValue reflect.Value) reflect.Value {
	if reflectValue.Kind() == reflect.Ptr {
		return reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) reflect.Type {
	if reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		return reflectType.Elem()
	}
	return reflectType
}

//

// Keys creates a flat slice of keys from a nested structure exported fields.
// It uses Default Builder.
func Keys(v interface{}) ([]string, error) {
	return Default.Keys(v)
}

// Values creates a flat slice of values from a nested structure exported fields.
// It uses Default Builder.
func Values(v interface{}) ([]interface{}, error) {
	return Default.Values(v)
}

// Map creates a map with Keys(): Values() from a nested structure.
// It uses Default Builder.
func Map(v interface{}) (map[string]interface{}, error) {
	return Default.Map(v)
}

//

// NewBuilder creates new flat struct builder with
// parameterized tag name and keyDelimiter which delimits key parts
// when nested struct converted to flat.
func NewBuilder(tag, keyDelimiter string) *Builder {
	return &Builder{tag, keyDelimiter}
}
