package middle

import (
	"log"
	"net/http"
)

//**Задача:**
//Написать middleware `logging(next http.Handler) http.Handler`, который:
//
//- замеряет время обработки запроса;
//
//- логирует метод, путь, статус и длительность.

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func logging(next http.Handler) http.Handler {
	// TODO: обернуть next, логировать начало/конец
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", logging(mux)))
}
