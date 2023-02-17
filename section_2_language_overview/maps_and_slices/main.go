package main

import (
	"log"
	"sort"
)


type User struct {
	FirstName string
	LastName string
}

func main() {
	//simple variables
	var myString string
	var myInt int

	myString = "Hi"
	myInt = 11

	mySecondString := "another string"

	log.Println(myString, mySecondString, myInt)


	//map
	// var myOtherMap map[string]string //bad way to create maps

	// make(map[key]value)
	myMap := make(map[string]string) //good way to create a map
	myMap2 := make(map[string]int)
	myMap["dog"] = "Samson"
	myMap["otherDog"] = "Cassy"
	

	//overwrite key
	myMap["dog"] = "Fido"

	myMap2["first"] = 1

	log.Println((myMap["dog"]))
	log.Println((myMap["otherDog"]))
	log.Println((myMap2["first"]))



	myCustomMap := make(map[string]User)
	me := User {
		FirstName: "Trevor",
		LastName: "Solar",
	}

	myCustomMap["me"] = me

	log.Println(myCustomMap["me"].FirstName)


	var myNewVar float32

	myNewVar = 11.1
	log.Println(myNewVar)



	// Slices

	var mySlice []string
	var mySlice2 []int
	//add stuff to slice
	mySlice = append(mySlice, "crevor")
	mySlice = append(mySlice, "srevor2")
	mySlice = append(mySlice, "arevor23")

	mySlice2 = append(mySlice2, 3);
	mySlice2 = append(mySlice2, 2);
	mySlice2 = append(mySlice2, 1);


	//sort slices
	sort.Ints(mySlice2)
	sort.Strings(mySlice)
	log.Println((mySlice2))
	log.Println((mySlice))


	numbers := []int{1,2,3,4,5,6,7,8,9, 10}

	log.Println((numbers))
	//print first 3
	log.Println((numbers[4:6]))

	//slice of diff types

	Names := []string{"one", "seven", "fish", "cat"}
	log.Println(Names)
}