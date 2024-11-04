package main

import (
	"Anastasia/worker-pool/pool"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wp := pool.NewWorkerPool()
	// Запуск воркеров в отдельных потоках.
	// Асинхронно.
	for i := 0; i < 6; i++ {
		go wp.AddWorker(&wg)
	}

	time.Sleep(time.Second)
	wp.DeleteWorker()
	wp.DeleteWorker()
	wp.DeleteWorker()

	// Поток обработки результатов.
	// Асинхронно.
	go func(ch chan string) {
		for val := range ch {
			fmt.Println(val)
		}
	}(wp.Res)
	// Отправка заданий в поток.
	// Синхронно.
	for i := 0; i < 10; i++ {
		wp.AddJob(fmt.Sprintf("Job №%d", i))
	}

	for i := 0; i < 100; i++ {
		wp.AddJob(fmt.Sprintf("Job №%d", i))
	}
	close(wp.Jobs)
	wg.Wait()
}
