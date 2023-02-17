package plugins

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	//seed random input for random number
	rand.Seed(time.Now().Unix())
	value := rand.Intn(n)
	
	return value
}