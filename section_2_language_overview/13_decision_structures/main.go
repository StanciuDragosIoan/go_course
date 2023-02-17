package main

import "log"

func main() {


	var isTrue bool
	isTrue = true

	if(isTrue) {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	cat := "cat"

	if (cat == "cat") {
		log.Println("cat is", cat)
	} else {
		log.Println("cat is not a cat")
	}


	myNum := 100
	isTrue2 := false

	if(myNum > 99 && !isTrue2) {
		log.Println("myNum is greater than 99 and isTrue2 is", isTrue2)
	} else if(myNum < 100 && isTrue) {
		log.Println("1")
	} else if(myNum == 101 || isTrue) {
		log.Println("2")
	} else if (myNum > 1000 && isTrue == false) {
		log.Println("3")
	}






	//switch

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("default cat is something else")
	}
}