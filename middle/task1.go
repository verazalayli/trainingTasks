package middle

//**Задача:**
//Написать HTTP-хендлер `/echo`, который:
//
//- принимает POST с JSON `{ "message": "..." }`;
//
//- валидирует, что `message` не пустой;
//
//- отвечает `200 OK` с `{ "echo": "<message>" }` или `400` с ошибкой в JSON.

import (
	"log"
	"net/http"
)

type EchoRequest struct {
	Message string `json:"message"`
}

type EchoResponse struct {
	Echo string `json:"echo"`
	Err  string `json:"error,omitempty"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", echoHandler)

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// TODO: реализовать
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Проверить метод
	// 2. Распарсить JSON в EchoRequest
	// 3. Провалидировать Message
	// 4. Вернуть JSON EchoResponse
}
