package flatstructs

import (
	"fmt"
	"reflect"
)

type ErrInvalid struct {
	v interface{}
}

func (e *ErrInvalid) Error() string {
	return fmt.Sprintf(
		"Reflect reports this value is invalid '%#v'",
		e.v,
	)
}

func NewErrInvalid(v interface{}) error {
	return &ErrInvalid{v}
}

//

type ErrInvalidKind struct {
	expected reflect.Kind
	got      reflect.Kind
}

func (e *ErrInvalidKind) Error() string {
	return fmt.Sprintf(
		"Expected '%s' kind, got '%s'",
		e.expected,
		e.got,
	)
}

func NewErrInvalidKind(expected, got reflect.Kind) error {
	return &ErrInvalidKind{expected, got}
}

//

type ErrPtrRequired struct {
	v interface{}
}

func (e *ErrPtrRequired) Error() string {
	return fmt.Sprintf(
		"A pointer to the value '%#v' is required, not the value itself",
		e.v,
	)
}

func NewErrPtrRequired(v interface{}) error {
	return &ErrPtrRequired{v}
}
