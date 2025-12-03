package senior

import (
	"sync"
	"time"
)

//In-memory KV-store c HTTP API
//
//**Задача:**
//Реализовать сервер:
//
//- `PUT /kv/{key}` с JSON `{"value": "..."}` — сохраняет;
//
//- `GET /kv/{key}` — отдаёт;
//
//- `DELETE /kv/{key}` — удаляет.
//
//
//Требования:
//
//- потокобезопасный `map[string]string` (`sync.RWMutex`);
//
//- опциональный `ttl` в секундах в запросе, фоновой cleaner истекающих ключей.

type KV struct {
	mu   sync.RWMutex
	data map[string]item
}

type item struct {
	value string
	ttl   time.Time // zero если без TTL
}

// TODO: методы Set/Get/Delete + HTTP handlers
