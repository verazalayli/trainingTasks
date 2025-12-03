package senior

//Собственный retry-HTTP-клиент с backoff
//
//**Задача:**
//Функция:
//- ретраит при 5xx и network-ошибках;
//- использует экспоненциальный backoff + jitter;
//- останавливается по `ctx`.

//func DoWithRetry(ctx context.Context, client *http.Client, req *http.Request, maxRetries int) (*http.Response, error) {
//
//}
