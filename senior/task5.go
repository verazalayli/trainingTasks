package senior

import "sync"

//Idempotent POST
//
//**Задача:**
//
//- Хендлер `POST /payments` принимает JSON с полями `{ "idempotency_key": "...", "amount": ... }`.
//
//- При повторном запросе с тем же ключом:
//
//- не создаёт новую «платёжку»;
//
//- возвращает тот же результат, что и в первый раз.
//
//
//Нужно хранить результаты в памяти (`map[idempotencyKey]Response`).

type PaymentRequest struct {
	Key    string  `json:"idempotency_key"`
	Amount float64 `json:"amount"`
}

type PaymentResponse struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Amount float64 `json:"amount"`
}

type IdemStore struct {
	mu   sync.Mutex
	data map[string]PaymentResponse
}
