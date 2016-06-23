// Copyright 2015 mparaiso<mparaiso@online.fr>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package lo

import (
	"fmt"
	"reflect"
)

// Chunk Creates an array of elements split into groups the length of size. If collection can’t be split evenly, the final chunk will be the remaining elements.
//func Chunk(collection []interface{}, length int) []interface{} {
//	result := []interface{}{}
//	if length <= 0 {
//		return result
//	}
//	temp := collection[:]
//	for len(temp) > 0 {
//		a := []interface{}{}
//		for i := 0; i < length && i <= len(temp); i++ {
//			// a.Push(temp.Shift())
//			a = append(a, temp[:1]...)
//			temp = temp[1:]
//		}
//		result = append(result, a)
//	}
//	return result
//}

// Difference creates an array excluding all provided values
func Difference(in Collection, values Collection, out Collection) error {
	return Filter(in, func(el interface{}, i int, in Collection) (bool, error) {
		index, err := IndexOf(values, el, 0)
		return index == -1, err
	}, out)
}

// Intersection creates an array of unique values in all provided arrays
func Intersection(in Collection, values Collection, out Collection) error {
	if err := Unique(in, &in); err != nil {
		return err
	}
	if err := Unique(values, &values); err != nil {
		return err
	}
	return Filter(in, func(el interface{}, i int) (bool, error) {
		index, err := IndexOf(values, el, 0)
		return index >= 0, err
	}, out)
}

// Xor creates an array that is the symmetric difference of the provided arrays.
func Xor(in Collection, values Collection, out Pointer) error {
	var unionResult interface{}
	if err := Union(in, values, &unionResult); err != nil {
		return err
	}
	var intersectionResult interface{}
	if err := Intersection(in, values, &intersectionResult); err != nil {
		return err
	}
	return Difference(unionResult, intersectionResult, out)
}

// LastIndexOf returns the index of a value in the array from its end, returns -1 if not found
//func LastIndexOf(array []interface{}, value interface{}, fromIndex int) int {
//	switch {
//	case fromIndex >= (len(array) - 1):
//	case fromIndex < 0:
//	default:
//		for i := len(array) - 1; i >= 0; i-- {
//			if value == array[i] {
//				return i
//			}
//		}
//	}
//	return -1
//}

// Union outputs a collection that holds unique values of in and values collections
func Union(in Collection, values Collection, out Pointer) error {

	if !IsCollection(in) {
		return NotACollection("Value %v is not a collection", in)
	}
	if !IsCollection(values) {
		return NotACollection("Value %v is not a collection", values)
	}
	if !IsPointer(out) {
		return NotPointer("Value %v is not a pointer", out)
	}
	inVal := reflect.ValueOf(in)
	valuesVal := reflect.ValueOf(values)
	outVal := reflect.ValueOf(out)
	if a, b := inVal.Type(), valuesVal.Type(); a != b {
		return NotAssignable("Collection of type '%v' doesn't match type  ", a, b)
	}
	if a, c := inVal.Type(), outVal.Elem().Type(); !a.AssignableTo(c) {
		return NotAssignable("Collection of type '%v' is not assignable to output of type '%v'", a, c)
	}
	for i := 0; i < valuesVal.Len(); i++ {
		inVal = reflect.Append(inVal, valuesVal.Index(i))
	}
	if err := Unique(inVal.Interface(), out); err != nil {
		return err
	}
	return nil
}

// Unique filters remove duplicate values from an array
func Unique(in Collection, out Pointer) error {
	if !IsCollection(in) {
		return NotACollection("Value '%v' is not a collection", in)
	}
	if !IsPointer(out) {
		return NotPointer("Value '%v' is not a pointer", out)
	}
	inValue := reflect.ValueOf(in)
	inType := inValue.Type()
	outValue := reflect.ValueOf(out)
	outType := outValue.Type()
	if !inType.AssignableTo(outType.Elem()) {
		return NotAssignable("Value in of type '%v' can't be assigned to out  of type '%v' ", inType, outType.Elem())
	}
	newCollection := reflect.MakeSlice(inType, 0, 0)
	inLen := inValue.Len()
	for i := 0; i < inLen; i++ {
		if index, err := IndexOf(newCollection.Interface(), inValue.Index(i).Interface(), 0); err != nil {
			return err
		} else if index == -1 {
			newCollection = reflect.Append(newCollection, inValue.Index(i))
		}

	}
	outValue.Elem().Set(newCollection)
	return nil
}

func GroupBy(in Collection, function Function, out Pointer) error {
	if !IsCollection(in) {
		return NotACollection("Value %v is not a collection", in)
	}
	if !IsFunction(function) {
		return NotAFunction("Value %v is not a function", function)
	}
	if !IsPointer(out) {
		return NotPointer("Value %v is not a pointer", out)
	}
	functionValue := reflect.ValueOf(function)
	functionType := functionValue.Type()
	mapParam1 := functionType.Out(0)
	mapParam2 := functionType.Out(1)
	numOut := functionType.NumOut()
	sliceOfMapParam2Type := reflect.SliceOf(mapParam2)
	mapResultType := reflect.MapOf(mapParam1, sliceOfMapParam2Type)
	outValue := reflect.ValueOf(out)
	outType := outValue.Type()

	mapResultValue := reflect.MakeMap(mapResultType)

	collectionValue := reflect.ValueOf(in)
	l := collectionValue.Len()
	for i := 0; i < l; i++ {
		resultValues := functionValue.Call([]reflect.Value{collectionValue.Index(i)})
		if numOut == 3 {
			if err, ok := resultValues[2].Interface().(error); ok && err != nil {
				return err
			}
		}
		if !mapResultValue.MapIndex(resultValues[0]).IsValid() {
			mapResultValue.SetMapIndex(resultValues[0], reflect.MakeSlice(sliceOfMapParam2Type, 0, 0))
		}
		mapResultValue.SetMapIndex(resultValues[0], reflect.Append(mapResultValue.MapIndex(resultValues[0]), resultValues[1]))
	}
	if mapResultType.AssignableTo(outType.Elem()) {
		outValue.Elem().Set(mapResultValue)
	} else {
		return NotAssignable("Result of type '%v' is not assignable to out type '%v'", mapResultType, outType)
	}
	return nil
}

type Pipeline struct {
	Value Any
	Queue []func() error
}

func In(collection Collection) *Pipeline {
	return &Pipeline{collection, []func() error{}}
}

func (pipeline *Pipeline) Map(function Function) *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Map(pipeline.Value, function, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Reduce(function Function, initial Any) *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Reduce(pipeline.Value, function, initial, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Filter(predicate Function) *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Filter(pipeline.Value, predicate, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Union(values Collection) *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Union(pipeline.Value, values, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Xor(values Collection) *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Xor(pipeline.Value, values, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Unique() *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Unique(pipeline.Value, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Intersection(values Collection) *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Intersection(pipeline.Value, values, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Difference(collection Collection) *Pipeline {
	pipeline.Queue = append(pipeline.Queue, func() error {
		return Difference(pipeline.Value, collection, &pipeline.Value)
	})
	return pipeline
}

func (pipeline *Pipeline) Out(out Pointer) error {
	if !IsPointer(out) {
		return NotPointer("Value '%s' is not a pointer", out)
	}
	for i, operation := range pipeline.Queue {
		if err := operation(); err != nil {
			return fmt.Errorf("At step %d : %s", i, err.Error())
		}
	}
	outValue := reflect.ValueOf(out)
	outType := outValue.Type()
	v := reflect.ValueOf(pipeline.Value)
	vType := v.Type()
	if !vType.AssignableTo(outType.Elem()) {
		return NotAssignable("Value of type '%s' is not assignable to out value type '%s'", vType, outType)
	}
	outValue.Elem().Set(v)
	return nil
}

// Filter iterates over elements of collection
// returning an array of all elements predicate returns truthy for.
// The predicate is invoked with three arguments: (value, index|key, collection).

func Filter(in Collection, predicate Function, out Pointer) error {
	if !IsCollection(in) {
		return NotACollection("Value '%v' is not a collection ", in)
	}
	if !IsFunction(predicate) {
		return NotAFunction("Value '%v' is not a function", predicate)
	}
	if !IsPointer(out) {
		return NotPointer("Value '%v' is not a pointer", out)
	}
	inValue := reflect.Indirect(reflect.ValueOf(in))
	inType := inValue.Type()
	predicateValue := reflect.ValueOf(predicate)
	predicateType := predicateValue.Type()
	numIn := predicateType.NumIn()
	numOut := predicateType.NumOut()
	if numIn < 1 || numIn > 4 {
		return IncorrectInputParameterArity("In parameter arity should be greater than one and lower than four, got '%d'", numIn)
	}
	if numOut < 1 || numOut > 3 {
		return IncorrectOutputParameterArity("In parameter arity should be greater than one and lower than three,got '%d'", numOut)
	}
	outValue := reflect.ValueOf(out)
	resultValue := reflect.MakeSlice(inType, 0, 0)
	if a, b := resultValue.Type(), outValue.Elem().Type(); !a.AssignableTo(b) {
		return NotAssignable("Value of type '%v' is not assignable to out value of type '%v'", a, b)
	}
	for i := 0; i < inValue.Len(); i++ {
		var results []reflect.Value
		elementValue := inValue.Index(i)
		iValue := reflect.ValueOf(i)
		switch numIn {
		case 1:
			results = predicateValue.Call([]reflect.Value{elementValue})
		case 2:
			results = predicateValue.Call([]reflect.Value{elementValue, iValue})
		case 3:
			results = predicateValue.Call([]reflect.Value{elementValue, iValue, inValue})
		}
		if numOut == 2 {
			if err, ok := results[1].Interface().(error); ok {
				return err
			}
		}
		if results[0].Bool() {

			resultValue = reflect.Append(resultValue, elementValue)
		}
	}
	outValue.Elem().Set(resultValue)
	return nil
}

// IndexOf returns the index of a value in the array from its begining, returns -1 if not found
func IndexOf(collection Collection, element Any, start int) (int, error) {
	if !IsCollection(collection) {
		return 0, NotACollection("Value '%v' is not a collecion ", collection)
	}
	collectionValue := reflect.ValueOf(collection)
	collectionLen := collectionValue.Len()
	if start >= collectionLen {
		return -1, nil
	}
	if start < 0 {
		start = 0
	}
	for i := start; i < collectionLen; i++ {
		if reflect.DeepEqual(collectionValue.Index(i).Interface(), element) {
			return i, nil
		}
	}
	return -1, nil
}

// ForEach executes handler on each element of the collection
//func ForEach(collection []interface{}, handler func(interface{}, int, []interface{})) {
//	for index, value := range collection {
//		handler(value, index, collection)
//	}
//}

// Map maps a collection to a result by passing each element of the collection
// to a function
func Map(collection Collection, function Function, result Pointer) error {
	if !IsCollection(collection) {
		return NotACollection("Value '%v' is not a collection", collection)
	}
	if !IsFunction(function) {
		return NotAFunction("Value '%v' is not a function", function)
	}
	if !IsPointer(result) {
		return NotPointer("Value '%v' is not a pointer", result)
	}
	functionValue := reflect.ValueOf(function)
	functionType := functionValue.Type()
	numIn := functionType.NumIn()
	if numIn < 1 || numIn > 3 {
		return IncorrectInputParameterArity("Function '%s' input parameter arity should be greater than 0 or lower than 4 , got '%d' ", functionType, numIn)
	}
	numOut := functionType.NumOut()
	if numOut < 1 || numOut > 2 {
		return IncorrectInputParameterArity("Function '%s' output parameter arity should be greater than 0 or lower than 3 , got '%d' ", functionType, numOut)
	}
	if numOut == 2 {
		if a, b := functionType.Out(1), reflect.TypeOf((*error)(nil)).Elem(); !a.AssignableTo(b) {
			return IncorrectOutputParameterType("Type '%s' of output parameter 2 should be an error interface, got '%s'", a)
		}
	}
	collectionValue := reflect.Indirect(reflect.ValueOf(collection))
	resultValue := reflect.ValueOf(result)
	resultType := resultValue.Type()
	sliceOfCollectionElementType := reflect.SliceOf(functionType.Out(0))
	newCollectionValue := reflect.MakeSlice(sliceOfCollectionElementType, 0, 0)
	if !newCollectionValue.Type().AssignableTo(sliceOfCollectionElementType) {
		return NotAssignable("Type '%s' is not assignable to return type '%s'", sliceOfCollectionElementType, resultType)
	}
	for i := 0; i < collectionValue.Len(); i++ {
		var resultValues []reflect.Value
		switch numIn {
		case 1:
			resultValues = functionValue.Call([]reflect.Value{collectionValue.Index(i)})

		case 2:
			resultValues = functionValue.Call([]reflect.Value{collectionValue.Index(i), reflect.ValueOf(i)})
		case 3:
			resultValues = functionValue.Call([]reflect.Value{collectionValue.Index(i), reflect.ValueOf(i), collectionValue})
		}
		if numOut == 2 {
			if err := resultValues[1]; !err.IsNil() {
				return err.Interface().(error)
			}
		}
		newCollectionValue = reflect.Append(newCollectionValue, resultValues[0])
	}
	resultValue.Elem().Set(newCollectionValue)
	return nil
}

// Reduce returns a value from an array by applying a reducer function to each element of an array. That value can be anything
func Reduce(collection Collection, function Function, initial Any, resultPointer Pointer) error {
	collectionValue := reflect.Indirect(reflect.ValueOf(collection))
	if !IsCollection(collection) {
		return NotACollection("Collection '%v' is not a slice")
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
		return NotPointer("Result '%v' not a pointer", resultPointer)
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

// Any is purely semantic
type Any interface{}

// Pointer is purely semantic
type Pointer interface{}

func IsPointer(value Any) bool {
	if reflect.ValueOf(value).Kind() != reflect.Ptr {
		return false
	}
	return true
}

// IsFunction returns true if f is a function
func IsFunction(f Function) bool {
	if reflect.ValueOf(f).Kind() != reflect.Func {
		return false
	}
	return true
}

func IsCollection(value Any) bool {
	if k := reflect.ValueOf(value).Kind(); k == reflect.Slice || k == reflect.Array {
		return true
	}
	return false
}

// NotASlice returns a NotASliceError
func NotACollection(format string, arguments ...interface{}) NotACollectionError {
	return NotACollectionError(fmt.Sprintf(format, arguments...))
}

// NotASliceError signals that a value isn't a slice or an array
type NotACollectionError string

// Error returns a string
func (err NotACollectionError) Error() string {
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

func NotPointer(format string, arguments ...interface{}) NotAPointerError {
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
