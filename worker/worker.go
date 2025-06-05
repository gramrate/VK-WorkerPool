/*
Package worker предоставляет реализацию отдельного воркера.

Воркеры получают задания через канал и обрабатывают их.
*/
package worker

import (
	"fmt"
	"sync"
)

// Worker представляет воркера.
type Worker struct {
	ID      int             // Уникальный идентификатор воркера.
	StopCh  chan struct{}   // Канал для остановки воркера.
	WG      *sync.WaitGroup // Счётчик для ожидания завершения воркера.
	JobChan <-chan string   // Канал для получения заданий.
}

// Start запускает обработку заданий в отдельной горутине.
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

// Stop останавливает воркера.
func (w *Worker) Stop() {
	close(w.StopCh)
}
