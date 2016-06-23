package lo_test

import (
	"fmt"

	//"github.com/interactiv/expect"
	"reflect"
	"testing"

	"strings"

	"github.com/mparaiso/lodash-go"
)

func ExampleGroupBy() {
	// Group people by country

	type Person struct {
		Name    string
		Country string
	}
	var group map[string][]Person
	err := lo.GroupBy([]Person{
		{"John", "France"},
		{"Jack", "Ireland"},
		{"Igor", "Germany"},
		{"Anna", "Germany"},
		{"Louise", "France"},
	}, func(person Person) (string, Person) {
		return person.Country, person
	}, &group)
	fmt.Println(err)
	fmt.Println(group["Germany"])
	fmt.Println(group["Ireland"])

	// Output:
	// <nil>
	// [{Igor Germany} {Anna Germany}]
	// [{Jack Ireland}]
}

//func TestWithout(t *testing.T) {
//	tests := []struct {
//		// Test description.
//		name string
//		// Parameters.
//		array  []interface{}
//		values []interface{}
//		// Expected results.
//		want []interface{}
//	}{
//		{"Array 1 without Array 2", []interface{}{1, 2, 3, 4}, []interface{}{2, 4}, []interface{}{1, 3}},
//	}
//	for _, tt := range tests {
//		if got := lo.Without(tt.array, tt.values...); !reflect.DeepEqual(got, tt.want) {
//			t.Errorf("%q. Without() = %v, want %v", tt.name, got, tt.want)
//		}
//	}
//}

//func TestLastIndexOf(t *testing.T) {
//	tests := []struct {
//		// Test description.
//		name string
//		// Parameters.
//		array     []interface{}
//		value     interface{}
//		fromIndex int
//		// Expected results.
//		want int
//	}{
//		{"Last index", []interface{}{1, 2, 3, 4, 5, 2}, 2, 0, 5},
//	}
//	for _, tt := range tests {
//		if got := lo.LastIndexOf(tt.array, tt.value, tt.fromIndex); got != tt.want {
//			t.Errorf("%q. LastIndexOf() = %v, want %v", tt.name, got, tt.want)
//		}
//	}
//}

func Example() {
	// Counting words
	const words = `Lorem ipsum nascetur,
            nascetur adipiscing. Aenean commodo nascetur.
            Aenean nascetur commodo ridiculus nascetur,
            commodo ,nascetur consequat.`

	var result map[string]int
	err := lo.In(strings.Split(words, " ")).
		Map(func(word string) string {
		return strings.Trim(strings.Trim(word, "\n\t "), ".,!")
	}).
		Filter(func(word string) bool {
		return word != ""
	}).
		Reduce(func(Map map[string]int, word string) map[string]int {
		Map[word] = Map[word] + 1
		return Map
	}, map[string]int{}).
		Out(&result)
	fmt.Println(err)
	fmt.Println(result["nascetur"])
	fmt.Println(result["commodo"])

	// Output:
	// <nil>
	// 6
	// 3
}

func ExamplePipeline() {
	// Error while queued operations are performed

	var result interface{}
	err := lo.In([]string{"a"}).
		Filter(func(element string) (bool, error) {
		return false, fmt.Errorf("Something went wrong")
	}).Out(&result)
	fmt.Println(err)

	// Output:
	// At step 0 : Something went wrong

}

func ExamplePipeline_Map() {
	// Compute the sum of all countries people
	type Country struct {
		Name       string
		Population int
	}
	var total int
	err := lo.In([]Country{{"France", 1000}, {"Spain", 5000}}).
		Map(func(country Country) int { return country.Population }).
		Reduce(func(total int, count int) int { return total + count }, 0).
		Out(&total)
	fmt.Println(err)
	fmt.Println(total)
	// Output:
	// <nil>
	// 6000
}

func ExamplePipeline_Reduce() {
	// Group people by country
	// Demonstrates the use of Reduce to transform an collection
	// into a map
	type Person struct {
		Name    string
		Country string
		Age     int
	}

	var adultPeopleByCountry map[string]int
	err := lo.In([]Person{
		{"John", "France", 18},
		{"Jane", "England", 16},
		{"Jack", "France", 20},
		{"Anna", "Spain", 19},
		{"Eduardo", "Spain", 30},
		{"Michel", "France", 12}}).
		Filter(func(person Person) bool { return person.Age >= 18 }).
		Reduce(func(list map[string]int, person Person) map[string]int {
		list[person.Country] = list[person.Country] + 1
		return list
	}, map[string]int{}).
		Out(&adultPeopleByCountry)

	fmt.Println(err)
	fmt.Println(adultPeopleByCountry["France"])
	fmt.Println(adultPeopleByCountry["Spain"])
	fmt.Println(adultPeopleByCountry["England"])
	// Output:
	// <nil>
	// 2
	// 2
	// 0

}

//func TestLast(t *testing.T) {
//	tests := []struct {
//		// Test description.
//		name string
//		// Parameters.
//		array []interface{}
//		// Expected results.
//		want interface{}
//	}{
//		{"Last element of collection", []interface{}{3, 4}, 4},
//	}
//	for _, tt := range tests {
//		if got := lo.Last(tt.array); !reflect.DeepEqual(got, tt.want) {
//			t.Errorf("%q. Last() = %v, want %v", tt.name, got, tt.want)
//		}
//	}
//}

//func TestZip(t *testing.T) {
//	tests := []struct {
//		// Test description.
//		name string
//		// Parameters.
//		arrays [][]interface{}
//		// Expected results.
//		want []interface{}
//	}{
//	// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		if got := lo.Zip(tt.arrays...); !reflect.DeepEqual(got, tt.want) {
//			t.Errorf("%q. Zip() = %v, want %v", tt.name, got, tt.want)
//		}
//	}
//}

func TestFilter(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		collection interface{}
		predicate  interface{}
		// Expected results.
		want interface{}
	}{
		{"Filter even numbers", []int{1, 2, 3, 4}, func(element int, index int, collection []int) bool {
			return element%2 == 0
		}, []int{2, 4}},
	}
	var got interface{}
	for _, tt := range tests {

		if err := lo.Filter(tt.collection, tt.predicate, &got); err != nil {
			t.Fatalf("Error Should be nil, got", err)
		}
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("Got '%v', want '%v'", got, tt.want)
		}
	}
}

func TestIndexOf(t *testing.T) {
	tests := []struct {
		// Test description.
		name string
		// Parameters.
		collection interface{}
		element    interface{}
		index      int
		// Expected results.
		want int
	}{
		{"", []int{1, 2, 3}, 3, 0, 2},
	}
	for _, tt := range tests {
		if got, err := lo.IndexOf(tt.collection, tt.element, tt.index); err != nil {
			t.Fatal(err)
		} else if got != tt.want {
			t.Errorf("lo.IndexOf : got '%d', want '%d'", got, tt.want)
		}
	}
}

func ExampleXor() {
	// Compute the symetrical difference between 2 collections of the same type
	var xor []string
	err := lo.Xor([]string{"a", "b", "c", "d"}, []string{"b", "c", "d", "e", "f"}, &xor)
	fmt.Println(err)
	fmt.Println(xor)
	// Output:
	// <nil>
	// [a e f]
}

func ExampleUnion() {
	// Compute the union of 2 slices
	var union []string
	err := lo.Union([]string{"a", "b", "c", "d", "e"},
		[]string{"a", "c", "g", "x", "e"},
		&union)
	fmt.Println(err)
	fmt.Println(union)
	// Output:
	// <nil>
	// [a b c d e g x]

}

func ExampleDifference() {
	var difference []string
	err := lo.Difference([]string{"a", "b", "c", "d"}, []string{"a", "c", "x"}, &difference)
	fmt.Println(err)
	fmt.Println(difference)
	// Output:
	// <nil>
	// [b d]
}

func ExampleIntersection() {
	var intersection []int
	err := lo.Intersection([]int{1, 2, 3, 4}, []int{0, 2, 4, 5}, &intersection)
	fmt.Println(err)
	fmt.Println(intersection)
	// Ouput:
	// <nil>
	// [2 4]
}

func ExampleUnique() {
	var uniqueValues []string
	err := lo.Unique([]string{"a", "e", "a", "c", "b", "d"}, &uniqueValues)
	fmt.Println(err)
	fmt.Println(uniqueValues)
	// Output:
	// <nil>
	// [a e c b d]
}

func ExampleFilter() {
	var evenNumbers []int
	err := lo.Filter([]int{0, 1, 2, 3, 4}, func(element int) bool {
		return element%2 == 0
	}, &evenNumbers)
	fmt.Println(err)
	fmt.Println(evenNumbers)
	// Output:
	// <nil>
	// [0 2 4]
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
		{"", func() {}, "", reflect.TypeOf(lo.NotACollection("")), nil, "NotASliceError"},
		{[8]byte{}, 1, "", reflect.TypeOf(lo.NotAFunction("")), nil, "NotAFunctionError"},
		{[]int{}, func(result int, element int) int { return result }, 0, reflect.TypeOf(lo.NotPointer("")), 0, "NotAPointer"},
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
		//t.Log(test.Result, err)
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

func ExampleMap() {
	var result []int
	err := lo.Map([]int{2, 3, 4}, func(element int) int { return element * 2 }, &result)
	fmt.Println(err)
	fmt.Println(result)

	// Output:
	// <nil>
	// [4 6 8]

}

func ExampleMap_second() {
	type Data struct {
		Name  string
		Index int
	}
	var result2 []Data
	err := lo.Map([]string{"John", "Jane", "Jack"}, func(element string, index int) Data {
		return Data{element, index}
	}, &result2)
	fmt.Println(err)
	fmt.Println(result2)
	// Output:
	// <nil>
	// [{John 0} {Jane 1} {Jack 2}]
}

func ExampleMap_third() {
	// The function mapper can return an error as a second argument,
	// which will be returned by lo.Map
	var result []int
	err := lo.Map([]bool{true, false, true}, func(element bool, index int, collection []bool) (int, error) {
		if !element {
			return 0, fmt.Errorf("Elements shouldn't be false, got '%v' at index '%d' for collection '%v'", element, index, collection)
		}
		return index, nil
	}, &result)
	fmt.Println(err)
	// Output:
	// Elements shouldn't be false, got 'false' at index '1' for collection '[true false true]'
}
