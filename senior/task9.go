package senior

//**Задача:**
//Без сторонних библиотек, только stdlib.
//- В приложении считают:
//
//- общее число запросов;
//
//- время обработки каждого запроса (histogram buckets вручную или просто `avg`/`max`).
//
//- `/metrics` отдаёт текст в стиле Prometheus, например:
//    http_requests_total 123
//    http_request_duration_ms_avg 12.3
