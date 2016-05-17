// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lo

import (
	a "github.com/interactiv/datastruct/array"
	"reflect"
	"testing"
)

type Expectation struct {
	value interface{}
	test  *testing.T
}

func (e *Expectation) toEqual(val interface{}) {
	if e.value != val {
		e.test.Errorf("%+v should equal %+v", e.value, val)
	}
}

func Expect(val interface{}, t *testing.T) *Expectation {
	return &Expectation{val, t}
}

func TestChunk(t *testing.T) {
	type fixture struct {
		array    a.ArrayInterface
		length   int
		expected a.ArrayInterface
	}
	fixtures := a.New(&fixture{
		a.New(1, 2, 3, 4, 5),
		3,
		a.New(a.New(1, 2, 3), a.New(4, 5)),
	})
	fixtures.ForEach(func(el interface{}, i int) {
		fix := el.(*fixture)
		result := Chunk(fix.array, fix.length)
		t.Logf("%+v", result)
		fix.expected.ForEach(func(chunk interface{}, i int) {
			chunk.(a.ArrayInterface).ForEach(func(val interface{}, j int) {
				Expect(result.At(i).(a.ArrayInterface).At(j), t).toEqual(val)
			})
		})

	})

}

func TestWithout(t *testing.T) {
	arr := a.New(1, 2, 3)
	res := Without(arr, 1, 2)
	Expect(res.At(0), t).toEqual(3)
}

func TestDifference(t *testing.T) {
	type fixture struct {
		array1   a.ArrayInterface
		array2   a.ArrayInterface
		expected a.ArrayInterface
	}
	fixtures := a.New(
		&fixture{
			a.New(1, 2, 3),
			a.New(4, 2),
			a.New(1, 3),
		},
		&fixture{
			a.New(4, 2),
			a.New(1, 2, 3),
			a.New(4),
		},
		&fixture{
			a.New(1, 2),
			a.New(4, 2),
			a.New(1),
		},
	)
	fixtures.ForEach(func(el interface{}, i int) {
		fix := el.(*fixture)
		result := Difference(fix.array1, fix.array2)
		t.Log(result)
		fix.expected.ForEach(func(el interface{}, i int) {
			Expect(result.At(i), t).toEqual(el)
		})
	})
}

func TestUnion(t *testing.T) {
	type fixture struct {
		arrays   []a.ArrayInterface
		expected a.ArrayInterface
	}
	fixtures := a.New(
		&fixture{
			[]a.ArrayInterface{a.New(1, 2), a.New(4, 2), a.New(2, 1)},
			a.New(1, 2, 4),
		},
	)
	fixtures.ForEach(func(el interface{}, i int) {
		fix := el.(*fixture)
		result := Union(fix.arrays...)
		fix.expected.ForEach(func(el interface{}, i int) {
			Expect(result.At(i), t).toEqual(el)
		})
	})
}

func TestUniq(t *testing.T) {
	type fixture struct {
		array    a.ArrayInterface
		expected a.ArrayInterface
	}
	fixtures := a.New(&fixture{
		a.New(5, 2, 3, 4, 5, 2, 6, 1),
		a.New(5, 2, 3, 4, 6, 1),
	})
	fixtures.ForEach(func(fix interface{}, i int) {
		fixt := fix.(*fixture)
		res := Unique(fixt.array)
		fixt.expected.ForEach(func(val interface{}, i int) {
			Expect(val, t).toEqual(res.At(i))
		})
	})
}

func TestIntersection(t *testing.T) {
	type fixture struct {
		args     []a.ArrayInterface
		expected a.ArrayInterface
	}
	fixtures := a.New(
		&fixture{
			[]a.ArrayInterface{a.New(1, 2), a.New(4, 2), a.New(2, 1)},
			a.New(2),
		},
		&fixture{
			[]a.ArrayInterface{a.New(1, 2, 3, 1, 5, 2)},
			a.New(1, 2, 3, 5),
		},
		&fixture{
			[]a.ArrayInterface{},
			a.New(),
		},
	)
	fixtures.ForEach(func(val interface{}, i int) {
		fix := val.(*fixture)
		res := Intersection(fix.args...)
		fix.expected.ForEach(func(val interface{}, i int) {
			Expect(res.At(i), t).toEqual(val)
		})
	})
}

func TestXor(t *testing.T) {
	type fixture struct {
		args     []a.ArrayInterface
		expected a.ArrayInterface
	}
	fixtures := []*fixture{
		&fixture{
			[]a.ArrayInterface{a.New(1, 2), a.New(4, 2)},
			a.New(1, 4),
		},
	}
	for _, fixture := range fixtures {
		res := Xor(fixture.args...)
		//t.Log(res)
		fixture.expected.ForEach(func(val interface{}, i int) {
			Expect(res.At(i), t).toEqual(val)
		})
	}
}

func TestZip(t *testing.T) {
	//	t.Skip()
	type fixture struct {
		arguments []a.ArrayInterface
		expected  a.ArrayInterface
	}

	fixtures := []*fixture{
		&fixture{
			[]a.ArrayInterface{
				a.New("fred", "barney"),
				a.New(30, 40),
				a.New(true, false),
			},
			a.New(
				a.New("fred", 30, true),
				a.New("barney", 40, false),
			),
		},
	}
	for _, fix := range fixtures {
		ZipFunction := reflect.ValueOf(Zip)
		var values []reflect.Value

		for _, argument := range fix.arguments {
			values = append(values, reflect.ValueOf(argument))
		}

		result := ZipFunction.Call(values)[0]
		fix.expected.ForEach(func(el interface{}, i int) {
			expected := el.(a.ArrayInterface)
			result := result.Interface().(a.ArrayInterface).At(i).(a.ArrayInterface)
			Expect(Equal(result, expected), t).toEqual(true)
		})
	}
}

func TestEqual(t *testing.T) {
	type fixture struct {
		arrays   []a.ArrayInterface
		expected bool
	}
	sample := &struct{ i int }{1}
	number := 2
	fixtures := []*fixture{
		&fixture{
			[]a.ArrayInterface{a.New("a", "b"), a.New(1, "b"), a.New("a", "b")},
			false,
		},
		&fixture{
			[]a.ArrayInterface{},
			true,
		},
		&fixture{
			[]a.ArrayInterface{a.New("a", "b")},
			true,
		},
		&fixture{
			[]a.ArrayInterface{a.New("a", "b"), a.New("a", "b")},
			true,
		},
		&fixture{
			[]a.ArrayInterface{a.New(1, 2, 3), a.New(1, 2, 3), a.New(1, 2, 3)},
			true,
		},
		&fixture{
			[]a.ArrayInterface{a.New(1, nil, 3), a.New(1, nil, 3)},
			true,
		},
		&fixture{
			[]a.ArrayInterface{a.New(sample), a.New(sample)},
			true,
		},
		&fixture{
			[]a.ArrayInterface{a.New(*sample), a.New(*sample)},
			true,
		},
		&fixture{
			[]a.ArrayInterface{a.New(&number), a.New(&number)},
			true,
		},
	}
	for _, fix := range fixtures {
		Expect(Equal(fix.arrays...), t).toEqual(fix.expected)
	}
}
