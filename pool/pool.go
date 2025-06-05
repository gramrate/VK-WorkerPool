/*
Package pool предоставляет примитивный пул воркеров (WorkerPool).

Пул позволяет:

	— добавлять новых воркеров (AddWorker),
	— удалять воркеров (RemoveWorker),
	— отправлять задания (Submit),
	— ожидать завершения всех воркеров (Wait).

Каждый воркер обрабатывает задания, полученные через канал.
*/
package pool

import (
	"WorkerPool/worker"
	"errors"
	"sync"
)

// Pool управляет воркерами и отправляет им задания.
type Pool struct {
	workers []*worker.Worker // Список воркеров.
	mu      sync.Mutex       // Мьютекс для защиты доступа к срезу workers.
	jobChan chan string      // Канал для заданий.
	wg      sync.WaitGroup   // Счётчик активных воркеров.
	nextID  int              // Следующий ID для нового воркера.
}

// NewPool создаёт новый пул воркеров.
func NewPool() *Pool {
	return &Pool{
		workers: make([]*worker.Worker, 0),
		jobChan: make(chan string),
		nextID:  1,
	}
}

// AddWorker добавляет нового воркера в пул.
func (p *Pool) AddWorker() {
	p.mu.Lock()
	defer p.mu.Unlock()

	w := &worker.Worker{
		ID:      p.nextID,
		StopCh:  make(chan struct{}),
		JobChan: p.jobChan,
		WG:      &p.wg,
	}
	p.nextID++
	p.wg.Add(1)
	w.Start()
	p.workers = append(p.workers, w)
}

// RemoveWorker удаляет последнего воркера из пула.
func (p *Pool) RemoveWorker() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.workers) == 0 {
		return errors.New("no workers to remove")
	}

	w := p.workers[len(p.workers)-1]
	p.workers = p.workers[:len(p.workers)-1]
	w.Stop()
	return nil
}

// Submit отправляет задание воркерам.
func (p *Pool) Submit(job string) {
	p.jobChan <- job
}

// Wait ожидает завершения всех воркеров.
func (p *Pool) Wait() {
	p.wg.Wait()
}

// WorkerCount возвращает текущее количество воркеров в пуле.
func (p *Pool) WorkerCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.workers)
}
