package dash

import (
	d "github.com/interactiv/datastruct"
)

func Chunk(collection d.ArrayInterface, length int) d.ArrayInterface {
	result := d.NewArray()
	if length <= 0 {
		return result
	}
	temp := collection.Slice()
	for temp.Length() > 0 {
		a := d.NewArray()
		for i := 0; i < length && i <= temp.Length(); i++ {
			a.Push(temp.Shift())
		}
		result.Push(a)
	}
	return result
}

func Difference(array d.ArrayInterface, values d.ArrayInterface) d.ArrayInterface {
	if values.Length() > array.Length() {
		values, array = array, values
	}
	return array.Filter(func(el interface{}, i int) bool {
		return values.IndexOf(el, 0) == -1
	})
}

// Intersection creates an array of unique values in all provided arrays
func Intersection(arrays ...d.ArrayInterface) d.ArrayInterface {
	switch len(arrays) {
	case 0:
		return d.NewArray()
	case 1:
		return Unique(arrays[0])
	case 2:
		u := Unique(arrays[1])
		return Unique(arrays[0]).Filter(func(el interface{}, i int) bool {
			return u.IndexOf(el, 0) >= 0
		})
	}
	return Intersection(append([]d.ArrayInterface{Intersection(arrays[0], arrays[1])}, arrays[2:]...)...)
}

// Xor creates an array that is the symmetric difference of the provided arrays.
func Xor(arrays ...d.ArrayInterface) d.ArrayInterface {
	switch len(arrays) {
	case 0:
		return d.NewArray()
	case 1:
		return Unique(arrays[0].Slice())
	case 2:
		return Difference(Unique(arrays[0]), Unique(arrays[1]))

	}
	return Xor(append([]d.ArrayInterface{Xor(arrays[0], arrays[1])}, arrays[2:]...)...)
}
func IndexOf(array d.ArrayInterface, value interface{}, fromIndex int) int {
	return array.IndexOf(value, fromIndex)
}

func LastIndexOf(array d.ArrayInterface, value interface{}, fromIndex int) int {
	return array.LastIndexOf(value, fromIndex)
}

func Union(arrays ...d.ArrayInterface) d.ArrayInterface {
	result := d.NewArray()
	for _, array := range arrays {
		array.ForEach(func(el interface{}, i int) {
			if result.IndexOf(el, 0) == -1 {
				result.Push(el)
			}
		})
	}
	return result
}

func Unique(array d.ArrayInterface) d.ArrayInterface {
	res := d.NewArray()
	array.ForEach(func(val interface{}, i int) {
		if res.IndexOf(val, 0) == -1 {
			res.Push(val)
		}
	})
	return res
}

func First(array d.ArrayInterface) interface{} {
	return array.At(0)
}

func Last(array d.ArrayInterface) interface{} {
	return array.At(array.Length() - 1)
}

/*


type DashStruct struct {
	*datastruct.Array
}



func Dash(array interface{}) DashInterface {
	return &DashStruct{datastruct.NewArrayFrom(array)}
}
*/
