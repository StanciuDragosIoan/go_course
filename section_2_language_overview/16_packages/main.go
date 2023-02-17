package main

import (
	"log"

	greetings "example.com/greetings/helpers"
)

func main() {
 log.Println("Hello");

 message := greetings.Hello("Gladys")
 log.Println(message) 

 
}