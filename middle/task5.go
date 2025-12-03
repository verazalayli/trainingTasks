package middle

import (
	"log"
	"net/http"
)

//**Задача:**
//
//- `POST /enqueue` — кладёт строку в очередь;
//
//- `GET /dequeue` — достаёт следующий элемент или `204` если пусто.
//
//Использовать **канал** как очередь.

type Queue struct {
	ch chan string
}

func newQueue(size int) *Queue {
	return &Queue{ch: make(chan string, size)}
}

func (q *Queue) EnqueueHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: прочитать body как строку и отправить в q.ch (с обработкой блокировки)
}

func (q *Queue) DequeueHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: если элемент есть - вернуть, если нет - 204
}

func main() {
	q := newQueue(100)
	mux := http.NewServeMux()
	mux.HandleFunc("/enqueue", q.EnqueueHandler)
	mux.HandleFunc("/dequeue", q.DequeueHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
