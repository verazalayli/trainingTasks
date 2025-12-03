package middle

import (
	"net"
	"net/http"
	"sync"
	"time"
)

//**Задача:**
//Сделать middleware, который ограничивает количество запросов с одного IP:
//
//- максимум N запросов в минуту;
//
//- при превышении отдаёт `429 Too Many Requests`.

type limiter struct {
	mu    sync.Mutex
	hits  map[string]int
	reset time.Time
	limit int
}

func newLimiter(limit int) *limiter {
	return &limiter{
		hits:  make(map[string]int),
		reset: time.Now().Add(time.Minute),
		limit: limit,
	}
}

func (l *limiter) allow(ip string) bool {
	// TODO: сбрасывать счётчик раз в минуту, считать запросы по IP
	return true
}

func rateLimit(l *limiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		if !l.allow(ip) {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
