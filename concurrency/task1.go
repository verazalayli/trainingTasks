package concurrency

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"sync"
)

//Реализовать worker-pool с ограничением по горутинам**
//
//**Условие:**
//Дан список URL. Нужно реализовать worker-pool, который делает не более **5** параллельных запросов, собирает результаты и возвращает слайс.
//
//Условия:
//
//- использовать `context.Context` для отмены;
//
//- если одна горутина упала с ошибкой — отменить все остальные;
//
//- использовать `errgroup`.

func task1() {
	urls := []string{"http://...", "..."}
	ctx := context.Background()

	r, err := workerPool(ctx, urls)

	if err != nil {
		panic(err)
	}

	fmt.Println(r)
}

func workerPool(ctx context.Context, urls []string) ([]string, error) {
	_, ctx = errgroup.WithContext(context.Background())
	sem := make(chan struct{}, 5)
	results := make([]string, len(urls))
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, url := range urls {
		sem <- struct{}{}
		wg.Add(1)
		go func() error {
			defer func() { <-sem }()

			body, err := fetch(ctx, url)
			if err != nil {
				return err
			}

			mu.Lock()
			results = append(results, body)
			mu.Unlock()

			wg.Done()

			return nil
		}()
	}

	wg.Wait()

	return results, nil
}

func fetch(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
