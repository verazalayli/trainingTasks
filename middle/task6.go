package middle

import (
	"context"
	"net/http"
	"time"
)

//Хендлер `/proxy`:
//
//- принимает `GET ?url=...`;
//
//- создаёт `context.WithTimeout` (например, 2s);
//
//- делает запрос к переданному URL;
//
//- возвращает статус и первые 100 байт тела;
//
//- если timeout — отдаёт `504 Gateway Timeout`.

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		http.Error(w, "bad url", http.StatusBadRequest)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	// TODO: обработать ошибки/timeout через ctx.Err(), прочитать первые байты
}
