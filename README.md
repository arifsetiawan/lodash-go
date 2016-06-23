#lo

[![Build Status](https://travis-ci.org/Mparaiso/lodash-go.svg?branch=master)](https://travis-ci.org/Mparaiso/lodash-go)
[![GoDoc](https://godoc.org/github.com/Mparaiso/lodash-go?status.svg)](https://godoc.org/github.com/Mparaiso/lodash-go)
## A port of lodash and underscore in golang.

lo allows Go developers to work efficiently with collections in Go by providing generic functional programming methods,
such as map, reduce and filter, that work on slices and arrays of any types. lo is coded in Go.


### Install it:

	go get github.com/mparaiso/lodash-go


### Implemented functions 

	- [x] Map
	- [x] Reduce
	- [x] Filter
	- [x] Union
	- [x] Intersection
	- [x] Xor
	- [x] Difference
	- [x] IndexOf
	
### Examples

#### Filter out odd numbers

```go
	package main
	
	import(
		"fmt"
		"github.com"/mparaiso/lodash-go"
	)
	
	func Main(){
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

```
#### Compute the mean of an array of integers
	
```go
	package main

	import "github.com/Mparaiso/lodash-go"
	import "fmt"
		
	func Main(){
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
```
		
#### Grouping adults by country using a pipeline

```go
	package main
	
	import "github.com/mparaiso/lodash-go" 
	import "fmt"
	
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
```

#### Using a pipeline to compute the sum of all people in countries

```go
	package main
	
	import "github.com/Mparaiso/lodash-go"
	import "fmt"
	
	func Main(){
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
```
