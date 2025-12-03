package middle

import (
	"net/http"
	"sync"
)

//**Задача:**
//Хендлер `/stats`:
//
//- хранит количество обращений к каждому пути (`/echo`, `/sum` и т.п.);
//
//- возвращает JSON вида `{"/echo": 10, "/sum": 3}`.
//
//
//Нужно использовать `sync.Mutex` или `sync.RWMutex`.

type Stats struct {
	mu   sync.RWMutex
	hits map[string]int
}

func newStats() *Stats {
	return &Stats{hits: make(map[string]int)}
}

func (s *Stats) Middleware(next http.Handler) http.Handler {
	// TODO: инкрементировать счётчик по r.URL.Path
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func (s *Stats) Handler(w http.ResponseWriter, r *http.Request) {
	// TODO: отдать копию карты в JSON
}
