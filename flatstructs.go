package flatstructs

import (
	"reflect"
	"strings"
	"unicode"
)

var (
	Default = NewBuilder("key", "")
)

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

func (b *Builder) Keys(v interface{}) ([]string, error) {
	err := checkValue(v)
	if err != nil {
		return nil, err
	}

	return b.keys(v, []string{})
}

func (b *Builder) keys(v interface{}, prefix []string) ([]string, error) {
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
			subKeys, err = b.keys(
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

func (b *Builder) Values(v interface{}) ([]interface{}, error) {
	err := checkValue(v)
	if err != nil {
		return nil, err
	}

	return b.values(v)
}

func (b *Builder) values(v interface{}) ([]interface{}, error) {
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
			subValues, err = b.values(
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
	// FIXME: Find a proper way of filtering unexported fields.
	// XXX: http://coub.com/view/ru550
	return unicode.IsUpper(rune(field.Name[0]))
}

func indirectValue(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		return reflectValue.Elem()
	}
	return reflectValue
}

func indirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
		return reflectType.Elem()
	}
	return reflectType
}

//

func Keys(v interface{}) ([]string, error) {
	return Default.Keys(v)
}

func Values(v interface{}) ([]interface{}, error) {
	return Default.Values(v)
}

//

func NewBuilder(tag, keyDelimiter string) *Builder {
	return &Builder{tag, keyDelimiter}
}
