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
