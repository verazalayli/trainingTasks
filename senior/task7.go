package senior

import (
	"fmt"
	"net/http"
	"time"
)

//SSE / streaming responses
//
//**Задача:**
//
//Сделать хендлер `/stream`, который:
//
//- держит соединение открытым (Content-Type: `text/event-stream`);
//
//- раз в секунду шлёт `"data: <timestamp>\n\n"`;
//
//- корректно завершает поток при закрытии соединения/контекста.

func streamHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "stream unsupported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.Context().Done():
			return
		case t := <-ticker.C:
			fmt.Fprintf(w, "data: %s\n\n", t.Format(time.RFC3339))
			flusher.Flush()
		}
	}
}
