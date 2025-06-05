package worker

import (
	"fmt"
	"sync"
)

type Worker struct {
	ID      int
	StopCh  chan struct{}
	WG      *sync.WaitGroup
	JobChan <-chan string
}

func (w *Worker) Start() {
	go func() {
		defer w.WG.Done()
		fmt.Printf("Worker %d started\n", w.ID)
		for {
			select {
			case job := <-w.JobChan:
				fmt.Printf("Worker %d processing job: %s\n", w.ID, job)
			case <-w.StopCh:
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.StopCh)
}
