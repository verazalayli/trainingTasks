package task6_generatorAndFilter

import (
	"fmt"
	"math/rand/v2"
)

func Task6() {
	NumberGeneratorAndFilter()
}

func NumberGeneratorAndFilter() {
	number := rand.IntN(100)
	chUnfiltered := make(chan int)
	chFiltered := make(chan int)
	go func() {
		chUnfiltered <- number
	}()

	go func() {
		if number%2 == 0 {
			chFiltered <- number
		} else {
			chFiltered <- 0
		}
	}()

	unfilteredNumbers := <-chUnfiltered
	filteredNumbers := <-chFiltered
	fmt.Println("Unfiltered", unfilteredNumbers)
	fmt.Println("Filtered", filteredNumbers)
}
