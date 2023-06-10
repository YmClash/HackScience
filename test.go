package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("holla word ")
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(101)
	fmt.Println(randomNumber)
}
