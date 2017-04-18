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
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestBuilderKeysNil(t *testing.T) {
	type Flat struct{}
	var (
		sample *Flat
	)

	keys, err := Keys(sample)
	if err == nil {
		t.Error("Nil as argument should be reported as ErrInvalid")
		return
	}

	if _, ok := err.(*ErrInvalid); !ok {
		t.Errorf(
			"Invalid error type, expected ErrInvalid, got '%T'",
			err,
		)
	}

	assert.Equal(
		t,
		([]string)(nil),
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNotStruct(t *testing.T) {
	sample := []string{}

	keys, err := Keys(&sample)
	if err == nil {
		t.Error("Nil as argument should be reported as ErrInvalidKind")
		return
	}

	if _, ok := err.(*ErrInvalidKind); !ok {
		t.Errorf(
			"Invalid error type, expected ErrInvalidKind, got '%T'",
			err,
		)
	}

	assert.Equal(
		t,
		([]string)(nil),
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysFlat(t *testing.T) {
	type Flat struct {
		Foo string
	}
	sample := Flat{"foo"}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"Foo"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysFlatMultiple(t *testing.T) {
	type Flat struct {
		Foo string
		Bar string
	}
	sample := Flat{"foo", "bar"}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"Foo", "Bar"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysFlatMultipleWithKeys(t *testing.T) {
	type Flat struct {
		Foo string `key:"foo"`
		Bar string
	}
	sample := Flat{"foo", "bar"}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"foo", "Bar"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysFlatUnexported(t *testing.T) {
	type Flat struct {
		foo string
		bar string
	}
	sample := Flat{"foo", "bar"}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNested(t *testing.T) {
	type Flat struct {
		Baz string `key:"baz"`
	}
	type Nested struct {
		Foo string `key:"foo"`
		Bar Flat
	}
	sample := Nested{"foo", Flat{"baz"}}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"foo", "Barbaz"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNestedPtr(t *testing.T) {
	type Flat struct {
		Baz string `key:"baz"`
	}
	type Nested struct {
		Foo string `key:"foo"`
		Bar *Flat
	}
	sample := Nested{"foo", &Flat{"baz"}}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"foo", "Barbaz"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNestedCustomDelimiter(t *testing.T) {
	type Flat struct {
		Baz string `key:"baz"`
	}
	type Nested struct {
		Foo string `key:"foo"`
		Bar *Flat
	}
	sample := Nested{"foo", &Flat{"baz"}}

	keys, err := NewBuilder("key", " -> ").Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"foo", "Bar -> baz"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNestedCustomKey(t *testing.T) {
	type Flat struct {
		Baz string `k:"baz"`
	}
	type Nested struct {
		Foo string `k:"foo"`
		Bar *Flat
	}
	sample := Nested{"foo", &Flat{"baz"}}

	keys, err := NewBuilder("k", "").Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"foo", "Barbaz"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNestedUnexported(t *testing.T) {
	type Flat struct {
		baz string
	}
	type Nested struct {
		Foo string
		Bar *Flat
		boo string
	}
	sample := Nested{"foo", &Flat{"baz"}, "boo"}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"Foo", "Bar"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNestedUnexportedMixed(t *testing.T) {
	type Flat struct {
		baz  string
		Booz string
	}
	type Nested struct {
		Foo string
		Bar *Flat
		boo string
	}
	sample := Nested{"foo", &Flat{"baz", "booz"}, "boo"}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"Foo", "BarBooz"},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNestedMixedTypes(t *testing.T) {
	type Flat struct {
		Baz  int
		Booz float64
		Goo  time.Time
		Zoo  []string
	}
	type Nested struct {
		Foo string
		Bar *Flat
	}
	sample := Nested{
		"foo",
		&Flat{
			0,
			0.1,
			time.Now(),
			[]string{"hello", "you"},
		},
	}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{
			"Foo",
			"BarBaz",
			"BarBooz",
			"BarGoo",
			"BarZoo",
		},
		keys,
		spew.Sdump(sample),
	)
}

func TestBuilderKeysNestedNil(t *testing.T) {
	type Flat struct {
		baz string
	}
	type Nested struct {
		Foo string
		Bar *Flat
		Baz string
	}
	sample := Nested{"foo", nil, "baz"}

	keys, err := Keys(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]string{"Foo", "Baz"},
		keys,
		spew.Sdump(sample),
	)
}

//

func TestBuilderValuesNil(t *testing.T) {
	type Flat struct{}
	var (
		sample *Flat
	)

	values, err := Values(sample)
	if err == nil {
		t.Error("Nil as argument should be reported as ErrInvalid")
		return
	}

	if _, ok := err.(*ErrInvalid); !ok {
		t.Errorf(
			"Invalid error type, expected ErrInvalid, got '%T'",
			err,
		)
	}

	assert.Equal(
		t,
		([]interface{})(nil),
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesNotStruct(t *testing.T) {
	sample := []string{}

	values, err := Values(&sample)
	if err == nil {
		t.Error("Nil as argument should be reported as ErrInvalidKind")
		return
	}

	if _, ok := err.(*ErrInvalidKind); !ok {
		t.Errorf(
			"Invalid error type, expected ErrInvalidKind, got '%T'",
			err,
		)
	}

	assert.Equal(
		t,
		([]interface{})(nil),
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesFlat(t *testing.T) {
	type Flat struct {
		Foo string
	}
	sample := Flat{"foo"}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo"},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesFlatMultiple(t *testing.T) {
	type Flat struct {
		Foo string
		Bar string
	}
	sample := Flat{"foo", "bar"}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo", "bar"},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesFlatMultipleWithEmptyValues(t *testing.T) {
	type Flat struct {
		Foo string
		Bar string
		Baz string
	}
	sample := Flat{
		Foo: "foo",
		Bar: "bar",
	}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo", "bar", ""},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesFlatUnexported(t *testing.T) {
	type Flat struct {
		foo string
		bar string
	}
	sample := Flat{"foo", "bar"}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesNested(t *testing.T) {
	type Flat struct {
		Baz string
	}
	type Nested struct {
		Foo string
		Bar Flat
	}
	sample := Nested{"foo", Flat{"baz"}}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo", "baz"},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesNestedPtr(t *testing.T) {
	type Flat struct {
		Baz string
	}
	type Nested struct {
		Foo string
		Bar *Flat
	}
	sample := Nested{"foo", &Flat{"baz"}}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo", "baz"},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesNestedUnexported(t *testing.T) {
	type Flat struct {
		baz string
	}
	type Nested struct {
		Foo string
		Bar *Flat
		boo string
	}
	sample := Nested{"foo", &Flat{"baz"}, "boo"}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo", Flat{"baz"}},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesNestedUnexportedMixed(t *testing.T) {
	type Flat struct {
		baz  string
		Booz string
	}
	type Nested struct {
		Foo string
		Bar *Flat
		boo string
	}
	sample := Nested{"foo", &Flat{"baz", "booz"}, "boo"}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo", "booz"},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesNestedMixedTypes(t *testing.T) {
	type Flat struct {
		Baz  int
		Booz float64
		Goo  time.Time
		Zoo  []string
	}
	type Nested struct {
		Foo string
		Bar *Flat
	}
	now := time.Now()
	sample := Nested{
		"foo",
		&Flat{
			0,
			0.1,
			now,
			[]string{"hello", "you"},
		},
	}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{
			"foo",
			0,
			0.1,
			now,
			[]string{"hello", "you"},
		},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderValuesNestedNil(t *testing.T) {
	type Flat struct {
		baz string
	}
	type Nested struct {
		Foo string
		Bar *Flat
		Baz string
	}
	sample := Nested{"foo", nil, "baz"}

	values, err := Values(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		[]interface{}{"foo", "baz"},
		values,
		spew.Sdump(sample),
	)
}

func TestBuilderMap(t *testing.T) {
	type Flat struct {
		baz  string
		Jazz string
	}
	type Nested struct {
		Foo string
		xyz string
		Bar *Flat
		Baz string
	}
	sample := Nested{
		"foo",
		"xyz",
		&Flat{"baz", "jazz"},
		"baz",
	}

	mapping, err := Map(&sample)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(
		t,
		map[string]interface{}{
			"Foo":     "foo",
			"BarJazz": "jazz",
			"Baz":     "baz",
		},
		mapping,
		spew.Sdump(sample),
	)
}
