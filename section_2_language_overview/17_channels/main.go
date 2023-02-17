package main

import (
	"log"

	plugins "test.com/plugins/test"
)

const numPool = 1000;

func CalculateValue(intChannel chan int) {
	randomNumber := plugins.RandomNumber(numPool)
	// pass random number to intChannel
	intChannel <- randomNumber
} 

func main() {
	 intChannel := make(chan int)
 
	 // close channel after function execution
	 defer close(intChannel)// defer (execute what s after defer after the current function is done)
	 //defer closes connection to file or to DB

	 //call function in a go routine (like a thread)
	 go CalculateValue(intChannel)

	 //get value from channel
	 num := <-intChannel
	 log.Println((num))
}

 