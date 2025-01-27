package main

import (

	"fmt"
	"time"
	"sync"
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

	var wg sync.WaitGroup

	for i :=0;i<50;i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}
	wg.Wait()


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

func printNumber (n int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 *time.Second)
	fmt.Println(n)
}