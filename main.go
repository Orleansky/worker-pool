package main

import (
	"Anastasia/worker-pool/pool"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wp := pool.NewWorkerPool()
	// Запуск воркеров в отдельных потоках.
	// Асинхронно.
	for i := 0; i < runtime.NumCPU(); i++ {
		go wp.AddWorker(&wg)
	}

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
	for i := 0; i < 1000; i++ {
		wp.AddJob(fmt.Sprintf("Job №%d", i))
	}
	close(wp.Jobs)
	wg.Wait()
}
