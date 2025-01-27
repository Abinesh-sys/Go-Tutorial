package main

import (

	"fmt"
) 

func main() {
	var name string = "Abinesh"
	var age int = 22
	fmt.Printf("Name: %s\nAge: %d\n", name, age)

	game := "RDR"
	rating := 10
	fmt.Printf("Game: %s\nRating: %d\n", game, rating)
    
	if rating >=10 {
		fmt.Println("Best game")
	} else {
		fmt.Println("Not best game")
	}

	for i:=0;i<5;i++ {
		fmt.Println(i)
	}

	fmt.Println(add(5, 6))

	p := Person{"Jhon Wick", 52}
	fmt.Printf("Name: %s\nAge: %d\n", p.Name, p.Age)
}

func add(a int, b int) int {
	return a+b
}

type Person struct {
	Name string
	Age int
}
