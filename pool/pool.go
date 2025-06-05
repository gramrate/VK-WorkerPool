package pool

import (
	"WorkerPool/worker"
	"errors"
	"sync"
)

type Pool struct {
	workers []*worker.Worker
	mu      sync.Mutex
	jobChan chan string
	wg      sync.WaitGroup
	nextID  int
}

func NewPool() *Pool {
	return &Pool{
		workers: make([]*worker.Worker, 0),
		jobChan: make(chan string),
		nextID:  1,
	}
}

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

func (p *Pool) Submit(job string) {
	p.jobChan <- job
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) WorkerCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.workers)
}
