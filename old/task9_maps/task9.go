package task9_maps

import "fmt"

func Task9() {
	wordsSlice := []string{"go", "java", "go", "python", "java", "go"}
	dictionaryMap := make(map[string]int)

	for _, word := range wordsSlice {
		dictionaryMap[word]++
	}
	for word, count := range dictionaryMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}
