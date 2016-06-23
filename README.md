#lo

[![Build Status](https://travis-ci.org/Mparaiso/lodash-go.svg?branch=master)](https://travis-ci.org/Mparaiso/lodash-go)
[![GoDoc](https://godoc.org/Mparaiso/lodash-go?status.svg)](https://godoc.org/Mparaiso/lodash-go)
## A port of lodash and underscore in golang.

lo allows Go developers to work efficiently with collections in Go by providing generic functional programming methods,
such as map, reduce and filter, that work on slices and arrays of any types. lo is coded in Go.


###Install it:

	go get github.com/mparaiso/lodash-go
	
###Examples:

	import "github.com/mparaiso/lodash-go" 
	import "fmt
	
	func Main(){
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
