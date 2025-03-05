package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"epsilon",
		"eta",
		"gamma",
		"pi",
		"theta",
		"zeta",
	}

	wg.Add(len(words))

	// wait groups execute in to particular order - it is up to the go scheduler
	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("This is the 2nd to be printed", &wg)
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // decrements wait group by 1
	fmt.Println(s)
}
