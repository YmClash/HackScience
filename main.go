package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	var nombre_chiffre int

	fmt.Println("Entre nombre de chiffre :  ")
	fmt.Scanf("%d", &nombre_chiffre)
	rand.Seed(time.Now().UnixNano())

	var numbers []int
	//var liste_arrange []int
	for i := 0; i < nombre_chiffre; i++ {
		randomNumber := rand.Intn(100)
		numbers = append(numbers, randomNumber)
	}

	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)

}
