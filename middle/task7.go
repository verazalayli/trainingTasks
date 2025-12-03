package middle

import (
	"net/http"
	"time"
)

//**Задача:**
//Сделать HTTP-сервер, который:
//
//- слушает `:8080`;
//
//- по сигналу `SIGINT/SIGTERM` вызывает `server.Shutdown(ctx)` с таймаутом;
//
//- корректно дожидается завершения запросов.

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte("done"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// TODO: слушать сигналы, запускать server.ListenAndServe в горутине,
	// по сигналу вызывать server.Shutdown с таймаутом.
}
