package senior

import "sync"

//Worker-pool для обработки HTTP задач
//
//**Задача:**
//
//- Хендлер `POST /tasks` принимает JSON `{ "payload": "..." }` и кладёт задачу в очередь.
//
//- Worker-pool из N воркеров обрабатывает задачи (например, пишет в лог с задержкой).
//
//- `GET /stats` отдаёт количество обработанных задач.
//
//
//Требования:
//
//- ограниченная очередь (channel);
//
//- при переполнении — либо `429`, либо блокировка (по конфигу);
//
//- корректное закрытие воркеров при shutdown.

type Task struct {
	Payload string `json:"payload"`
}

type WorkerPool struct {
	jobs chan Task
	wg   sync.WaitGroup
}

func NewWorkerPool(n, queue int) *WorkerPool {
	// TODO
	return nil
}

func (wp *WorkerPool) Start() {}
func (wp *WorkerPool) Stop()  {}

func (wp *WorkerPool) Enqueue(t Task) error {
	// TODO: отправить в канал или вернуть ошибку переполнения
	return nil
}
