package middle

import "net/http"

//**Задача:**
//Хендлер `/sum`:
//
//- принимает `POST` с JSON `{ "numbers": [1,2,3] }`;
//
//- считает сумму и среднее;
//
//- возвращает `{ "sum": ..., "avg": ... }`.

type SumRequest struct {
	Numbers []float64 `json:"numbers"`
}

type SumResponse struct {
	Sum float64 `json:"sum"`
	Avg float64 `json:"avg"`
	Err string  `json:"error,omitempty"`
}

func sumHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать аналогично M1, добавить сумму и среднее
}
