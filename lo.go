// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lo

import (
	"fmt"
	"reflect"
)

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

// Reduce returns a value from an array by applying a reducer function to each element of an array. That value can be anything
func Reduce(collection Collection, function Function, initial Any, resultPointer Pointer) error {
	collectionValue := reflect.Indirect(reflect.ValueOf(collection))
	if collectionValue.Kind() != reflect.Slice && collectionValue.Kind() != reflect.Array {
		return NotASlice("Collection '%v' is not a slice")
	}
	collectionType := collectionValue.Type()
	if !IsFunction(function) {
		return NotAFunction("Value '%v' is not a function", function)
	}
	functionValue := reflect.ValueOf(function)
	functionType := functionValue.Type()
	returnValue := reflect.ValueOf(initial)
	numberInParameters := functionType.NumIn()
	numberOutParameters := functionType.NumOut()
	if numberInParameters < 2 || numberInParameters > 4 {
		return IncorrectInputParameterArity("Incorrect input parameter arity, should >=2 or <=4, got %d", numberInParameters)
	}
	if numberOutParameters == 0 || numberOutParameters > 2 {
		return IncorrectOutputParameterArity("Incorrect output paramater arity, should be > 0 or <=2, got %d ", numberOutParameters)
	}

	if numberInParameters == 2 {
		if rt, ft := returnValue.Type(), functionType.In(0); rt != ft {
			return IncorrectInputParameterType("Illegal parameter type for first parameter in function '%s',should be '#%v', got '#%v' ", functionType, rt, ft)
		}
		if cet, ft := collectionType.Elem(), functionType.In(1); cet != ft || !cet.AssignableTo(ft) {
			return IncorrectInputParameterType("Illegal parameter type for second parameter in function '%s',should be '#%v', got '#%v' ", functionType, cet, ft)
		}
	}
	if numberInParameters == 3 {
		if indexType, parameterType := reflect.TypeOf(0), functionType.In(2); indexType != parameterType || !indexType.AssignableTo(parameterType) {
			return IncorrectInputParameterType("Illegal parameter type for third parameter in function '%s',should be '#%v', got '#%v' ", functionType, indexType, parameterType)

		}
	}
	if numberInParameters == 4 {
		if parameterType := functionType.In(3); collectionType != parameterType || !collectionType.AssignableTo(parameterType) {
			return IncorrectInputParameterType("Illegal parameter type for fourth parameter in function '%s',should be '#%v', got '#%v' ", functionType, collectionType, parameterType)

		}
	}

	if numberOutParameters == 1 {
		if rt, firstOutputParameterType := returnValue.Type(), functionType.Out(0); !rt.AssignableTo(firstOutputParameterType) {
			return IncorrectOutputParameterType("Incorrect output parameter type for first output parameter in function '%s', should be '#%v', got '#%v'", functionType, rt, firstOutputParameterType)
		}
	}

	if numberOutParameters == 2 {
		if outParameter2Type, errorType := functionType.Out(1), reflect.TypeOf((*error)(nil)).Elem(); !outParameter2Type.Implements(errorType) {
			return IncorrectOutputParameterType("Incorrect ouput paramater type for second output paramater in function '%s', should be '#%v', got '#%v'", functionType, errorType, outParameter2Type)
		}
	}
	// Do iterate over the collection an reduce it
	for i := 0; i < collectionValue.Len(); i++ {
		var returnValues []reflect.Value
		switch {
		case numberInParameters == 2:
			returnValues = functionValue.Call([]reflect.Value{returnValue, collectionValue.Index(i)})
		case numberInParameters == 3:
			returnValues = functionValue.Call([]reflect.Value{returnValue, collectionValue.Index(i), reflect.ValueOf(i)})
		case numberInParameters == 4:
			returnValues = functionValue.Call([]reflect.Value{returnValue, collectionValue.Index(i), reflect.ValueOf(i), collectionValue})
		}
		// on error return
		if numberOutParameters == 2 && !returnValues[1].IsNil() {
			return returnValues[1].Interface().(error)
		}
		returnValue = returnValues[0]
	}
	resultPointerValue := reflect.ValueOf(resultPointer)
	if resultPointerValue.Kind() != reflect.Ptr {
		return NotAPointer("Result '%v' not a pointer", resultPointer)
	}
	if resultPointerValue.Elem().Type().AssignableTo(returnValue.Type()) {
		resultPointerValue.Elem().Set(returnValue)
	} else if returnValue.Type().ConvertibleTo(resultPointerValue.Elem().Type()) {
		resultPointerValue.Elem().Set(returnValue.Convert(resultPointerValue.Elem().Type()))
	} else {
		return NotAssignable("Can't assign value '%v' to pointer of value '%v'", returnValue.Interface(), resultPointer)
	}
	return nil
}

type Collection interface{}

// Function is purely semantic
type Function interface{}

// IsFunction returns true if f is a function
func IsFunction(f Function) bool {
	if reflect.ValueOf(f).Kind() != reflect.Func {
		return false
	}
	return true
}

// Any is purely semantic
type Any interface{}

// Pointer is purely semantic
type Pointer interface{}

// NotASlice returns a NotASliceError
func NotASlice(format string, arguments ...interface{}) NotASliceError {
	return NotASliceError(fmt.Sprintf(format, arguments...))
}

// NotASliceError signals that a value isn't a slice or an array
type NotASliceError string

// Error returns a string
func (err NotASliceError) Error() string {
	return string(err)
}

func NotAFunction(format string, arguments ...interface{}) NotAFunctionError {
	return NotAFunctionError(fmt.Sprintf(format, arguments...))
}

type NotAFunctionError string

// Error returns a string
func (err NotAFunctionError) Error() string {
	return string(err)
}

func NotAPointer(format string, arguments ...interface{}) NotAPointerError {
	return NotAPointerError(fmt.Sprintf(format, arguments...))
}

// NotAPointerError is an error used to signal that a argument
// passed to a function is not a pointer
type NotAPointerError string

// Error returns a string
func (err NotAPointerError) Error() string {
	return string(err)
}
func NotAssignable(format string, arguments ...interface{}) NotAssignableError {
	return NotAssignableError(fmt.Sprintf(format, arguments...))
}

type NotAssignableError string

// Error returns a string
func (err NotAssignableError) Error() string {
	return string(err)
}

// IncorrectInputParameterType returns an IncorrectInputParameterTypeError
func IncorrectInputParameterType(format string, arguments ...interface{}) IncorrectInputParameterTypeError {
	return IncorrectInputParameterTypeError(fmt.Sprintf(format, arguments...))
}

// IncorrectInputParameterTypeError represent a error triggered when an input parameter passed to a function
// is not the right type
type IncorrectInputParameterTypeError string

func (err IncorrectInputParameterTypeError) Error() string {
	return string(err)
}

// IncorrectOutputParameterType returns an IncorrectOutputParameterTypeError
func IncorrectOutputParameterType(format string, arguments ...interface{}) IncorrectOutputParameterTypeError {
	return IncorrectOutputParameterTypeError(fmt.Sprintf(format, arguments...))
}

// IncorrectOutputParameterTypeError represent a error triggered when a value type is not of the same type
// as an output paramater type
type IncorrectOutputParameterTypeError string

func (err IncorrectOutputParameterTypeError) Error() string {
	return string(err)
}

func IncorrectInputParameterArity(format string, arguments ...interface{}) IncorrectInputParameterArityError {
	return IncorrectInputParameterArityError(fmt.Sprintf(format, arguments...))
}

type IncorrectInputParameterArityError string

func (err IncorrectInputParameterArityError) Error() string {
	return string(err)
}

func IncorrectOutputParameterArity(format string, arguments ...interface{}) IncorrectOutputParameterArityError {
	return IncorrectOutputParameterArityError(fmt.Sprintf(format, arguments...))
}

type IncorrectOutputParameterArityError string

func (err IncorrectOutputParameterArityError) Error() string {
	return string(err)
}

type ErrorInterface interface {
	Error() string
}
