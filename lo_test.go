package lo_test

import (
	"fmt"

	//"github.com/interactiv/expect"
	"reflect"
	"testing"

	"github.com/mparaiso/lodash-go"
)

func TestChunk(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		collection []interface{}
		length     int
		// Expected results.
		want []interface{}
	}{
		{"first test", []interface{}{1, 2, 3, 4}, 2, []interface{}{[]interface{}{1, 2}, []interface{}{3, 4}}},
	}
	for _, tt := range tests {
		if got := lo.Chunk(tt.collection, tt.length); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Chunk() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		array  []interface{}
		values []interface{}
		// Expected results.
		want []interface{}
	}{
		{"Difference between 2 arrays", []interface{}{1, 2, 3, 4}, []interface{}{2, 4}, []interface{}{1, 3}},
	}
	for _, tt := range tests {
		if got := lo.Difference(tt.array, tt.values); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Difference() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestWithout(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		array  []interface{}
		values []interface{}
		// Expected results.
		want []interface{}
	}{
		{"Array 1 without Array 2", []interface{}{1, 2, 3, 4}, []interface{}{2, 4}, []interface{}{1, 3}},
	}
	for _, tt := range tests {
		if got := lo.Without(tt.array, tt.values...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Without() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		arrays [][]interface{}
		// Expected results.
		want []interface{}
	}{
		{
			"Intersection of arrays",
			[][]interface{}{{1, 2, 3, 4}, {2, 3, 5}, {0, 2, 3, 6}},
			[]interface{}{2, 3},
		},
	}
	for _, tt := range tests {
		if got := lo.Intersection(tt.arrays...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Intersection() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		arrays [][]interface{}
		// Expected results.
		want []interface{}
	}{
		{
			"Xor of arrays",
			[][]interface{}{{1, 2, 3, 4}, {2, 3, 5}, {0, 2, 3, 6}},
			[]interface{}{1, 4, 5, 0, 2, 3, 6},
		},
	}
	for _, tt := range tests {
		if got := lo.Xor(tt.arrays...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Xor() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestLastIndexOf(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		array     []interface{}
		value     interface{}
		fromIndex int
		// Expected results.
		want int
	}{
		{"Last index", []interface{}{1, 2, 3, 4, 5, 2}, 2, 0, 5},
	}
	for _, tt := range tests {
		if got := lo.LastIndexOf(tt.array, tt.value, tt.fromIndex); got != tt.want {
			t.Errorf("%q. LastIndexOf() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		arrays [][]interface{}
		// Expected results.
		want []interface{}
	}{}
	for _, tt := range tests {
		if got := lo.Union(tt.arrays...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Union() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		array []interface{}
		// Expected results.
		want []interface{}
	}{
		{"Unique values", []interface{}{6, 4, 5, 6, 2, 4}, []interface{}{6, 4, 5, 2}},
	}
	for _, tt := range tests {
		if got := lo.Unique(tt.array); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Unique() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFirst(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		array []interface{}
		// Expected results.
		want interface{}
	}{
		{"First element of collection", []interface{}{3, 4}, 3},
	}
	for _, tt := range tests {
		if got := lo.First(tt.array); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. First() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		array []interface{}
		// Expected results.
		want interface{}
	}{
		{"Last element of collection", []interface{}{3, 4}, 4},
	}
	for _, tt := range tests {
		if got := lo.Last(tt.array); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Last() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestZip(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		arrays [][]interface{}
		// Expected results.
		want []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := lo.Zip(tt.arrays...); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Zip() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		arrays [][]interface{}
		// Expected results.
		want bool
	}{
		{"Array are equals", [][]interface{}{{1, 2}, {1, 2}, {1, 2}}, true},
	}
	for _, tt := range tests {
		if got := lo.Equal(tt.arrays...); got != tt.want {
			t.Errorf("%q. Equal() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestCompact(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		array []interface{}
		// Expected results.
		want []interface{}
	}{
		{"Remove nil values", []interface{}{0, nil, 2, nil}, []interface{}{0, 2}},
	}
	for _, tt := range tests {
		if got := lo.Compact(tt.array); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Compact() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		collection []interface{}
		predicate  func(interface{}, int, []interface{}) bool
		// Expected results.
		want []interface{}
	}{
		{"Filter even numbers", []interface{}{1, 2, 3, 4}, func(element interface{}, index int, collection []interface{}) bool {
			return element.(int)%2 == 0
		}, []interface{}{2, 4}},
	}
	for _, tt := range tests {
		if got := lo.Filter(tt.collection, tt.predicate); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Filter() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		collection []interface{}
		element    interface{}
		index      int
		// Expected results.
		want int
	}{
		{"", []interface{}{1, 2, 3}, 3, 0, 2},
	}
	for _, tt := range tests {
		if got := lo.IndexOf(tt.collection, tt.element, tt.index); got != tt.want {
			t.Errorf("%q. IndexOf() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestForEach(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		collection []interface{}
		handler    func(interface{}, int, []interface{})
	}{
		{"", []interface{}{1, 2, 3, 4}, func(element interface{}, i int, collection []interface{}) { _ = element.(int) }},
	}
	for _, tt := range tests {
		lo.ForEach(tt.collection, tt.handler)
	}
}

func TestReduce(t *testing.T) {
	type Args struct {
		Collection interface{}
		Function   interface{}
		Initial    interface{}
		Expected   interface{}
		Result     interface{}
		Label      string
	}
	t.Log("TestReduce Valid arguments")
	for _, test := range []Args{
		{[]int{1, 2, 3, 4}, func(result int, element int) int { return result + element }, 0, 10, 0, "Sum"},
		{[]string{"a", "b", "c", "d", "e"}, func(result string, element string, index int) string { return element + result }, "", "edcba", "", "Concat"},
		{[]int{10, 20, 30}, func(result int, element int, index int, collection []int) int {
			if l := len(collection); l-1 == index {
				return (result + element) / l
			} else {
				return result + element
			}
		}, 0, 20, 0, "average"},
	} {
		t.Logf("\t%s", test.Label)

		err := lo.Reduce(test.Collection, test.Function, test.Initial, &test.Result)
		if err != nil {
			t.Fatalf("\t\tError should be nil,got '%s'", err)
		}
		if test.Result != test.Expected {
			t.Fatalf("\t\tResult should be '%#v', got '%#v'", test.Expected, test.Result)
		}
	}
	t.Log("Test Reduce Invalid arguments")
	var illegalParameterType = reflect.TypeOf(lo.IncorrectInputParameterType(""))
	var incorrectOutputParameterType = reflect.TypeOf(lo.IncorrectOutputParameterType(""))
	for _, test := range []Args{
		{"", func() {}, "", reflect.TypeOf(lo.NotASlice("")), nil, "NotASliceError"},
		{[8]byte{}, 1, "", reflect.TypeOf(lo.NotAFunction("")), nil, "NotAFunctionError"},
		{[]int{}, func(result int, element int) int { return result }, 0, reflect.TypeOf(lo.NotAPointer("")), 0, "NotAPointer"},
		{[]int{1, 2, 3}, func(result int, element int) int { return result }, 0, reflect.TypeOf(lo.NotAssignable("")), &struct{}{}, "NotAssignableError"},
		{[]int{1, 2, 3}, func(result string, element string) string { return result }, 0, illegalParameterType, nil, "IllegalParameterType first"},
		{[]int{1, 2, 3}, func(result int, element string) int { return result }, 0, illegalParameterType, nil, "IllegalParameterType second"},
		{[]int{}, func(result int, element int, index string) int { return result }, 0, illegalParameterType, nil, "IllegalParameterType third"},
		{[]int{}, func(result int, element int, index int, collection []interface{}) int { return result }, 0, illegalParameterType, nil, "IllegalParameterType fourth"},
		{[]int{}, func(result int, element int, index int, collection []int) string { return string(result) }, 0, incorrectOutputParameterType, nil, "IncorrectOutputParameterType first"},
		{[]int{}, func(result int, element int, index int, collection []int) (int, int) { return result, 0 }, 0, incorrectOutputParameterType, nil, "IncorrectOutputParamaterType second"},
	} {
		t.Logf("\t%s", test.Label)

		err := lo.Reduce(test.Collection, test.Function, test.Initial, test.Result)
		t.Log(test.Result, err)
		if err == nil {
			t.Fatalf("\t\tError %#v should not be nil", err)
		}
		if actual := reflect.TypeOf(err); actual != test.Expected {
			t.Fatalf("\t\tError should be of type '%v' , got %v", test.Expected, actual)
		}
	}

}

// ExampleReduce shows how to use lo.Reduce to compute the mean of an array of number
func ExampleReduce() {
	// Compute the mean of an array of integers
	var mean int
	err := lo.Reduce([]int{10, 20, 30}, func(result int, element int, index int, collection []int) int {
		if l := len(collection); l-1 == index {
			return (result + element) / l
		} else {
			return result + element
		}
	}, 0, &mean)
	fmt.Println(mean)
	fmt.Println(err)
	// Output:
	// 20
	// <nil>
}

func ExampleReduce_second() {
	// Comput the sum of an array of integers
	var sum int
	lo.Reduce([]int{1, 2, 3, 4}, func(result int, element int) int { return result + element }, 0, &sum)
	fmt.Print(sum)
	// Output: 10
}
