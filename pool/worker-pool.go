package pool

import (
	"fmt"
	"sync"
)

type Worker struct {
	id   int
	stop chan struct{}
	done chan struct{}
}

// Создает новый экземпляр воркера

func NewWorker(id int) *Worker {
	return &Worker{
		id:   id,
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}
}

type WorkerPool struct {
	workers []*Worker
	Res     chan string
	Jobs    chan string
	mu      sync.Mutex
}

// Создает новый экземпляр пула

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		Res:  make(chan string),
		Jobs: make(chan string),
	}
}

// Добавляет новый воркер в пул

func (wp *WorkerPool) AddWorker(wg *sync.WaitGroup) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	worker := NewWorker(len(wp.workers))
	wp.workers = append(wp.workers, worker)

	wg.Add(1)
	go wp.startWorker(worker, wg)
}

// Удаляет один воркер из пула

func (wp *WorkerPool) DeleteWorker() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if len(wp.workers) > 1 {
		worker := wp.workers[len(wp.workers)-1]
		close(worker.stop)
		<-worker.done
		wp.workers = wp.workers[:len(wp.workers)-1]
	}
}

// Добавляет задачу

func (wp *WorkerPool) AddJob(job string) {
	wp.Jobs <- job
}

// Рабочий поток.
func (wp *WorkerPool) startWorker(worker *Worker, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(worker.done)

	for {
		select {
		case job, ok := <-wp.Jobs:
			if !ok {
				return
			}
			wp.Res <- fmt.Sprintf("Процесс №%d: %s", worker.id, job)
		case <-worker.stop:
			return
		}
	}
}
