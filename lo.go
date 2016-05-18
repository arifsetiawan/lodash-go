// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lo

// Chunk Creates an array of elements split into groups the length of size. If collection can’t be split evenly, the final chunk will be the remaining elements.
func Chunk(collection []interface{}, length int) []interface{} {
	result := []interface{}{}
	if length <= 0 {
		return result
	}
	temp := collection[:]
	for len(temp) > 0 {
		a := []interface{}{}
		for i := 0; i < length && i <= len(temp); i++ {
			// a.Push(temp.Shift())
			a = append(a, temp[:1]...)
			temp = temp[1:]
		}
		result = append(result, a)
	}
	return result
}

// Difference creates an array excluding all provided values
func Difference(array []interface{}, values []interface{}) []interface{} {
	return Filter(array, func(el interface{}, i int, array []interface{}) bool {
		return IndexOf(values, el, 0) == -1
	})
}

// Without creates an array excluding all provided values
func Without(array []interface{}, values ...interface{}) []interface{} {
	return Difference(array, values)
}

// Intersection creates an array of unique values in all provided arrays
func Intersection(arrays ...[]interface{}) []interface{} {
	switch len(arrays) {
	case 0:
		return []interface{}{}
	case 1:
		return Unique(arrays[0])
	case 2:
		u := Unique(arrays[1])
		a := Unique(arrays[0])
		return Filter(a, func(el interface{}, i int, a []interface{}) bool {
			return IndexOf(u, el, 0) >= 0
		})
	}
	return Intersection(append([][]interface{}{Intersection(arrays[0], arrays[1])}, arrays[2:]...)...)
}

// Xor creates an array that is the symmetric difference of the provided arrays.
func Xor(arrays ...[]interface{}) []interface{} {
	switch len(arrays) {
	case 0:
		return []interface{}{}
	case 1:
		return Unique(arrays[0])
	case 2:
		return append(Difference(Unique(arrays[0]), Unique(arrays[1])), Difference(Unique(arrays[1]), Unique(arrays[0]))...)

	}
	return Xor(append([][]interface{}{Xor(arrays[0], arrays[1])}, arrays[2:]...)...)
}

// LastIndexOf returns the index of a value in the array from its end, returns -1 if not found
func LastIndexOf(array []interface{}, value interface{}, fromIndex int) int {
	switch {
	case fromIndex >= (len(array) - 1):
	case fromIndex < 0:
	default:
		for i := len(array) - 1; i >= 0; i-- {
			if value == array[i] {
				return i
			}
		}
	}
	return -1
}

// Union returns an array filled by all unique values of the arrays
func Union(arrays ...[]interface{}) []interface{} {
	result := []interface{}{}
	for _, array := range arrays {
		ForEach(array, func(el interface{}, i int, array []interface{}) {
			if IndexOf(result, el, 0) == -1 {
				result = append(result, el)
			}
		})
	}
	return result
}

// Unique filters remove duplicate values from an array
func Unique(array []interface{}) []interface{} {
	res := []interface{}{}
	ForEach(array, func(val interface{}, i int, array []interface{}) {
		if IndexOf(res, val, 0) == -1 {
			res = append(res, val)
		}
	})
	return res
}

// First returns the first element of an array
func First(array []interface{}) interface{} {
	return array[0]
}

// Last returns the last element of an array
func Last(array []interface{}) interface{} {
	return array[len(array)-1]
}

// Zip zips an array
func Zip(arrays ...[]interface{}) []interface{} {
	result := []interface{}{}
	switch len(arrays) {
	case 0:
		return result
	case 1:
		return arrays[0][:]
	default:
		ForEach(arrays[0], func(el interface{}, i int, array []interface{}) {
			a := []interface{}{}
			for val := range arrays {
				a = append(array, val)
			}
			result = append(result, a)
		})
	}
	return result
}

// Equal compares arrays elements , return true if arrays are equal
func Equal(arrays ...[]interface{}) bool {
	switch len(arrays) {
	case 0, 1:
		return true
	case 2:
		return len(Difference(arrays[0], arrays[1])) == 0
	default:
		return Equal(arrays[0], arrays[1]) && Equal(arrays[2:]...)
	}
}

// Compact remove nil values from array
func Compact(array []interface{}) []interface{} {
	return Filter(array, func(el interface{}, i int, array []interface{}) bool {
		return el != nil
	})

}

// Filter iterates over elements of collection
// returning an array of all elements predicate returns truthy for.
// The predicate is invoked with three arguments: (value, index|key, collection).

func Filter(collection []interface{}, predicate func(interface{}, int, []interface{}) bool) []interface{} {
	result := []interface{}{}
	for index, value := range collection {
		if predicate(value, index, collection) {
			result = append(result, value)
		}
	}
	return result
}

// IndexOf returns the index of a value in the array from its begining, returns -1 if not found

func IndexOf(collection []interface{}, element interface{}, index int) int {
	if index >= len(collection) {
		return -1
	}
	if index < 0 {
		index = 0
	}
	for i := index; i < len(collection); i++ {
		if collection[i] == element {
			return i
		}
	}
	return -1
}

// ForEach executes handler on each element of the collection
func ForEach(collection []interface{}, handler func(interface{}, int, []interface{})) {
	for index, value := range collection {
		handler(value, index, collection)
	}
}