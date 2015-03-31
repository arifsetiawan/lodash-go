package dash

import (
	d "github.com/interactiv/datastruct"
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
		array    d.ArrayInterface
		length   int
		expected d.ArrayInterface
	}
	fixtures := d.NewArray(&fixture{
		d.NewArray(1, 2, 3, 4, 5),
		3,
		d.NewArray(d.NewArray(1, 2, 3), d.NewArray(4, 5)),
	})
	fixtures.ForEach(func(el interface{}, i int) {
		fix := el.(*fixture)
		result := Chunk(fix.array, fix.length)
		t.Logf("%+v", result)
		fix.expected.ForEach(func(chunk interface{}, i int) {
			chunk.(d.ArrayInterface).ForEach(func(val interface{}, j int) {
				Expect(result.At(i).(d.ArrayInterface).At(j), t).toEqual(val)
			})
		})

	})

}
func TestDifference(t *testing.T) {
	type fixture struct {
		array1   d.ArrayInterface
		array2   d.ArrayInterface
		expected d.ArrayInterface
	}
	fixtures := d.NewArray(
		&fixture{
			d.NewArray(1, 2, 3),
			d.NewArray(4, 2),
			d.NewArray(1, 3),
		},
		&fixture{
			d.NewArray(4, 2),
			d.NewArray(1, 2, 3),
			d.NewArray(1, 3),
		},
	)
	fixtures.ForEach(func(el interface{}, i int) {
		fix := el.(*fixture)
		result := Difference(fix.array1, fix.array2)
		fix.expected.ForEach(func(el interface{}, i int) {
			Expect(result.At(i), t).toEqual(el)
		})
	})
}

func TestUnion(t *testing.T) {
	type fixture struct {
		arrays   []d.ArrayInterface
		expected d.ArrayInterface
	}
	fixtures := d.NewArray(
		&fixture{
			[]d.ArrayInterface{d.NewArray(1, 2), d.NewArray(4, 2), d.NewArray(2, 1)},
			d.NewArray(1, 2, 4),
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
		array    d.ArrayInterface
		expected d.ArrayInterface
	}
	fixtures := d.NewArray(&fixture{
		d.NewArray(5, 2, 3, 4, 5, 2, 6, 1),
		d.NewArray(5, 2, 3, 4, 6, 1),
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
		args     []d.ArrayInterface
		expected d.ArrayInterface
	}
	fixtures := d.NewArray(
		&fixture{
			[]d.ArrayInterface{d.NewArray(1, 2), d.NewArray(4, 2), d.NewArray(2, 1)},
			d.NewArray(2),
		},
		&fixture{
			[]d.ArrayInterface{d.NewArray(1, 2, 3, 1, 5, 2)},
			d.NewArray(1, 2, 3, 5),
		},
		&fixture{
			[]d.ArrayInterface{},
			d.NewArray(),
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
		args     []d.ArrayInterface
		expected d.ArrayInterface
	}
	fixtures := []*fixture{
		&fixture{
			[]d.ArrayInterface{d.NewArray(1, 2), d.NewArray(4, 2)},
			d.NewArray(1, 4),
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
