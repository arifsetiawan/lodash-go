// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lo

import (
	a "github.com/interactiv/datastruct/array"
)

// Chunk Creates an array of elements split into groups the length of size. If collection can’t be split evenly, the final chunk will be the remaining elements.
func Chunk(collection a.ArrayInterface, length int) a.ArrayInterface {
	result := a.New()
	if length <= 0 {
		return result
	}
	temp := collection.Slice()
	for temp.Length() > 0 {
		a := a.New()
		for i := 0; i < length && i <= temp.Length(); i++ {
			a.Push(temp.Shift())
		}
		result.Push(a)
	}
	return result
}

// Difference creates an array excluding all provided values
func Difference(array a.ArrayInterface, values a.ArrayInterface) a.ArrayInterface {
	return array.Filter(func(el interface{}, i int) bool {
		return values.IndexOf(el, 0) == -1
	})
}

// Without creates an array excluding all provided values
func Without(array a.ArrayInterface, values ...interface{}) a.ArrayInterface {
	return Difference(array, a.New(values...))
}

// Intersection creates an array of unique values in all provided arrays
func Intersection(arrays ...a.ArrayInterface) a.ArrayInterface {
	switch len(arrays) {
	case 0:
		return a.New()
	case 1:
		return Unique(arrays[0])
	case 2:
		u := Unique(arrays[1])
		return Unique(arrays[0]).Filter(func(el interface{}, i int) bool {
			return u.IndexOf(el, 0) >= 0
		})
	}
	return Intersection(append([]a.ArrayInterface{Intersection(arrays[0], arrays[1])}, arrays[2:]...)...)
}

// Xor creates an array that is the symmetric difference of the provided arrays.
func Xor(arrays ...a.ArrayInterface) a.ArrayInterface {
	switch len(arrays) {
	case 0:
		return a.New()
	case 1:
		return Unique(arrays[0])
	case 2:
		return Difference(Unique(arrays[0]), Unique(arrays[1])).Concat(Difference(Unique(arrays[1]), Unique(arrays[0])))

	}
	return Xor(append([]a.ArrayInterface{Xor(arrays[0], arrays[1])}, arrays[2:]...)...)
}

// IndexOf returns the index of a value in the array from its begining, returns -1 if not found
func IndexOf(array a.ArrayInterface, value interface{}, fromIndex int) int {
	return array.IndexOf(value, fromIndex)
}

// LastIndexOf returns the index of a value in the array from its end, returns -1 if not found
func LastIndexOf(array a.ArrayInterface, value interface{}, fromIndex int) int {
	return array.LastIndexOf(value, fromIndex)
}

// Union returns an array filled by all unique values of the arrays
func Union(arrays ...a.ArrayInterface) a.ArrayInterface {
	result := a.New()
	for _, array := range arrays {
		array.ForEach(func(el interface{}, i int) {
			if result.IndexOf(el, 0) == -1 {
				result.Push(el)
			}
		})
	}
	return result
}

// Unique filters remove duplicate values from an array
func Unique(array a.ArrayInterface) a.ArrayInterface {
	res := a.New()
	array.ForEach(func(val interface{}, i int) {
		if res.IndexOf(val, 0) == -1 {
			res.Push(val)
		}
	})
	return res
}

// First returns the first element of an array
func First(array a.ArrayInterface) interface{} {
	return array.At(0)
}

// Last returns the last element of an array
func Last(array a.ArrayInterface) interface{} {
	return array.At(array.Length() - 1)
}

func Zip(arrays ...a.ArrayInterface) a.ArrayInterface {
	result := a.New()
	switch len(arrays) {
	case 0:
		return result
	case 1:
		return arrays[0].Slice()
	default:
		arrays[0].ForEach(func(el interface{}, i int) {
			array := a.New()
			for _, val := range arrays {
				array.Push(val.At(i))
			}
			result.Push(array)
		})
	}
	return result
}

// Equal compares arrays elements , return true if arrays are equal
func Equal(arrays ...a.ArrayInterface) bool {
	switch len(arrays) {
	case 0, 1:
		return true
	case 2:
		return Difference(arrays[0], arrays[1]).Length() == 0
	default:
		return Equal(arrays[0], arrays[1]) && Equal(arrays[2:]...)
	}
}

// Compact remove nil values from array
func Compact(array a.ArrayInterface) a.ArrayInterface {
	return array.Filter(func(el interface{}, i int) bool {
		return el != nil
	})

}
