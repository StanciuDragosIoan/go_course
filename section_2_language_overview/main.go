package main // package decoration
import (
	"fmt" // import from standard library (format)
)

func main() {
	fmt.Println("Hello, world.")
 

	var whatToSay string

	whatToSay = "Goodbye, cruel world."
	var i int

	fmt.Println(whatToSay);

	i = 7

	fmt.Println("i is set to", i);

	whatWasSaid, theOtherThingSaid := saySomething();

	fmt.Println("the function returned", whatWasSaid, theOtherThingSaid)
}


func saySomething() (string, string) {
	return "something", "else"
}

