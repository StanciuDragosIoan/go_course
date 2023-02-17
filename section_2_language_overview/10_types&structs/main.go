package main
import "log"
import "time"

// struct (custom type)
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {
	
	user := User {
		FirstName: "Trevor",
		LastName: "Sawler",
		PhoneNumber: "1 222 333 444",
	}

	log.Println(user.FirstName, user.LastName, user.BirthDate)

	log.Println(user.LastName)
}
//private
func whatever() {

}


//public (capital name makes property or function public)
func Whatever() {
	
}