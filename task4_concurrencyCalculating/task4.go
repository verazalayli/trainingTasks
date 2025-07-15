package task4_concurrencyCalculating

import "fmt"

func Task4(workers int) { //Название такое потому что это моя четвертая задачка
	slice := createSlice()
	chunkSize := len(slice) / workers
	remainder := len(slice) % workers

	resultChan := make(chan int, workers)

	start := 0
	for i := 0; i < workers; i++ {
		end := start + chunkSize
		if i == workers-1 {
			end += remainder
		}
		go func(sub []int) {
			sum := 0
			for _, v := range sub {
				sum += v
			}
			resultChan <- sum
		}(slice[start:end])
		start = end
	}

	total := 0
	for i := 0; i < workers; i++ {
		total += <-resultChan
	}

	fmt.Println("sum:", total)
}

func createSlice() []int {
	slice := []int{}
	for i := 1; i < 10001; i++ {
		slice = append(slice, i)
	}
	fmt.Println("slice length:", len(slice))
	return slice
}
