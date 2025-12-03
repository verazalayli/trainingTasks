package senior

import (
	"net/http"
	"sync"
	"time"
)

//Reverse-proxy с простым circuit-breaker
//
//**Задача:**
//
//- `/proxy/*` проксирует запросы на upstream (`UPSTREAM_URL`);
//
//- если подряд случилось ≥N ошибок (`5xx` или network), breaker «открывается» и в течение `M` секунд сразу отдаёт `503` без запроса к upstream;
//
//- потом переходит в half-open и пробует один запрос.

type CircuitBreaker struct {
	mu          sync.Mutex
	failures    int
	state       string // "closed", "open", "half-open"
	openedUntil time.Time
}

func (cb *CircuitBreaker) Allow() bool {
	// TODO: логика состояний
	return true
}

func (cb *CircuitBreaker) OnSuccess() {
	// TODO
}

func (cb *CircuitBreaker) OnFailure() {
	// TODO
}

func proxyHandler(cb *CircuitBreaker, upstream string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: использовать cb.Allow, http.Client, копировать статус/тело
	}
}
