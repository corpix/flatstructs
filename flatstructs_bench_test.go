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
)

func BenchmarkBuilderKeysFlat(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	flat := &Flat{}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Keys(flat)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkBuilderKeysFlatWithData(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	flat := &Flat{1, 2, 3, 4}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Keys(flat)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkBuilderKeysNested(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	type Nested struct {
		Foo *Flat
		Bar *Flat
		Baz Flat
		Daz Flat
	}
	nested := &Nested{
		&Flat{1, 2, 3, 4},
		&Flat{1, 2, 3, 4},
		Flat{1, 2, 3, 4},
		Flat{1, 2, 3, 4},
	}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Keys(nested)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkBuilderKeysNestedWithData(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	type Nested struct {
		Foo Flat
		Bar Flat
		Baz Flat
		Daz Flat
	}
	nested := &Nested{
		Foo: Flat{1, 2, 3, 4},
		Bar: Flat{1, 2, 3, 4},
		Baz: Flat{1, 2, 3, 4},
		Daz: Flat{1, 2, 3, 4},
	}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Keys(nested)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkBuilderKeysNestedHard(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	type Nested struct {
		Foo *Nested
		Bar *Nested
		Baz *Nested
		Daz *Nested
	}
	nested := &Nested{
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
	}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Keys(nested)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

//

func BenchmarkBuilderValuesFlat(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	flat := &Flat{1, 2, 3, 4}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Values(flat)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkBuilderValuesNested(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	type Nested struct {
		Foo *Flat
		Bar *Flat
		Baz Flat
		Daz Flat
	}
	nested := &Nested{
		&Flat{1, 2, 3, 4},
		&Flat{1, 2, 3, 4},
		Flat{1, 2, 3, 4},
		Flat{1, 2, 3, 4},
	}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Values(nested)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkBuilderValuesNestedHard(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	type Nested struct {
		Foo *Nested
		Bar *Nested
		Baz *Nested
		Daz *Nested
	}
	nested := &Nested{
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
		&Nested{
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
			&Nested{
				&Nested{},
				&Nested{},
				&Nested{},
				&Nested{},
			},
		},
	}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Values(nested)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkBuilderMapNested(b *testing.B) {
	type Flat struct {
		Foo          int
		Bar          int
		Baz          int
		LongerBurger int
	}
	type Nested struct {
		Foo *Flat
		Bar *Flat
		Baz *Flat
		Daz *Flat
	}
	nested := &Nested{
		&Flat{1, 2, 3, 4},
		&Flat{1, 2, 3, 4},
		&Flat{1, 2, 3, 4},
		&Flat{1, 2, 3, 4},
	}
	var (
		err error
	)
	for k := 0; k < b.N; k++ {
		_, err = Map(nested)
		if err != nil {
			b.Error(err)
			return
		}
	}
}
