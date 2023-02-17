package main

import (
	"log"
)

func main() {
	// loop over custom type
	type User struct {
		FirstName string
		LastName string
		Email string 
		Age int 
	}

	var users []User

	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 30})
	users = append(users, User{"ALex", "Smith", "alex@smith.com", 30})
	users = append(users, User{"Timmy", "Cook", "timmy@cook.com", 30})


	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}

	// for i :=0; i <= 10; i++ {
	// 	log.Println(i)
	// }


	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	//itertation, item
	for i, animal := range animals {
		log.Println(i, animal)
	}

	//no itertation
	for _, animal2 := range animals {
		log.Println(animal2)
	}


	//no iteration but with key specified!
	animalsMap := make(map[string]string)

	animalsMap["dog"] = "Fido"
	animalsMap["cat"] = "Timmy"
	for animalType, animal2 := range animalsMap {
		log.Println(animalType, animal2)
	}


	//loop over string
	var string = "Once upon a midnight dreary"

	for i, l := range string {
		log.Println(i, ":", l)
	}




}