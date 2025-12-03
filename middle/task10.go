package middle

import (
	"bufio"
	"log"
	"os"
)

//**Задача:**
//Приложение читает из `os.Stdin` поток строк, каждая — JSON `{ "id": 1, "value": "..." }`, и:
//
//- для каждой строки парсит JSON;
//
//- печатает только `id` и длину `value`.
//
//
//Фокус на `bufio.Scanner` + `encoding/json`.

type Item struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		// TODO: распарсить JSON в Item и вывести id и len(value)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
