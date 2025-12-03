package senior

import (
	"context"
	"net/http"
	"time"
)

//Batch-логер с backpressure
//
//**Задача:**
//
//- Хендлер `POST /log` принимает JSON строку и кладёт в канал;
//
//- отдельная горутина раз в N миллисекунд собирает batch до K сообщений и пишет в stdout;
//
//- queue bounded; при переполнении — `503` или блокировка.

type LogItem struct {
	Message string `json:"message"`
}

type Logger struct {
	ch      chan LogItem
	batchSz int
	flushMs time.Duration
}

func NewLogger(batchSz int, flushMs time.Duration) *Logger {
	// TODO
	return nil
}

func (l *Logger) Run(ctx context.Context) {
	// TODO: тикер + сбор батча
}

func (l *Logger) Handler(w http.ResponseWriter, r *http.Request) {
	// TODO: читать JSON и отправлять в канал
}
