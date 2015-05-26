package array

import (
	"fmt"

	"reflect"
)

// Reduce reduces an array to a single value
// CAN PANIC! if array isn't iterable
func Reduce(array interface{}, reducer func(result interface{}, element interface{}, index int) interface{}, initial interface{}) interface{} {
	arrayValue := reflect.ValueOf(array)
	switch arrayValue.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:

		for i := 0; i < arrayValue.Len(); i++ {
			initial = reducer(initial, arrayValue.Index(i).Interface(), i)
		}
		return initial
	default:
		panic(fmt.Sprintf("%v is not iterable", array))
	}
}

// IndexOf returns the index of element if element is found. Returns -1 if element is not found
// CAN PANIC!
func IndexOf(array interface{}, element interface{}) int {
	if !IsArray(array) {
		panic(fmt.Sprintf("%v is not iterable", array))
	}
	arrayValue := reflect.ValueOf(array)
	for i := 0; i < arrayValue.Len(); i++ {
		if arrayValue.Index(i).Interface() == element {
			return i
		}
	}
	return -1
}

// IsArray returns true if value is iterable
func IsArray(value interface{}) bool {
	arrayValue := reflect.ValueOf(value)
	switch arrayValue.Kind() {
	case reflect.Array, reflect.Slice, reflect.String:
		return true
	default:
		return false
	}

}
