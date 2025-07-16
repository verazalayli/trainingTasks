package task6_generatorAndFilter

import (
	"fmt"
	"math/rand/v2"
)

func Task6() {
	NumberGeneratorAndFilter()
}

func NumberGeneratorAndFilter() {
	chUnfiltered := make(chan int)
	chFiltered := make(chan int)

	go func() {
		for {
			chUnfiltered <- rand.IntN(100) + 1
		}
	}()

	go func() {
		for num := range chUnfiltered {
			if num%2 == 0 {
				chFiltered <- num
			}
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Printf("Even #%d: %d\n", i+1, <-chFiltered)
	}

	unfilteredNumbers := <-chUnfiltered
	filteredNumbers := <-chFiltered
	fmt.Println("Unfiltered", unfilteredNumbers)
	fmt.Println("Filtered", filteredNumbers)
}
