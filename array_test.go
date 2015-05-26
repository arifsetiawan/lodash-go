package array_test

import (
	"testing"

	"github.com/interactiv/array"
	"github.com/interactiv/expect"
)

func TestReduce(t *testing.T) {
	e := expect.New(t)
	list := []int{1, 2, 3}
	result := array.Reduce(list, func(res interface{}, elem interface{}, i int) interface{} {
		return res.(int) + elem.(int)
	}, 0)
	e.Expect(result).ToEqual(6)
	str := "hello"
	strReducer := func(str interface{}, elem interface{}, i int) interface{} {
		return string(elem.(uint8)) + str.(string)
	}
	e.Expect(array.Reduce(str, strReducer, "")).ToEqual("olleh")
}

func TestIndexOf(t *testing.T) {
	e := expect.New(t)
	var list = []string{"a", "b", "c"}
	e.Expect(array.IndexOf(list, "c")).ToEqual(2)
	e.Expect(array.IndexOf(list, "d")).ToEqual(-1)
	e.Expect(func() {
		array.IndexOf(1, "bar")
	}).ToPanic()
	type Foo struct {
		Bar, Baz int
	}
	var array2 = []*Foo{&Foo{1, 2}, &Foo{3, 4}}
	e.Expect(array.IndexOf(array2, array2[1])).ToEqual(1)
}
