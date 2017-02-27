// The MIT License (MIT)
//
// Copyright © 2017 Dmitry Moskowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
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
