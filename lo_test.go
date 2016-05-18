package lo_test

import (
	//"github.com/interactiv/expect"
	"github.com/mparaiso/lodash-go"
	"reflect"
	"testing"
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