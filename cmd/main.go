package main

import (
	"WorkerPool/pool"
	"fmt"
	"time"
)

func main() {
	p := pool.NewPool()
	p.AddWorker()
	p.AddWorker()

	for i := 1; i <= 5; i++ {
		p.Submit(fmt.Sprintf("Job-%d", i))
	}

	time.Sleep(time.Second)
	if err := p.RemoveWorker(); err != nil {
		panic(err)
	}

	time.Sleep(time.Second)
	p.AddWorker()

	for i := 6; i <= 100; i++ {
		p.Submit(fmt.Sprintf("Job-%d", i))
	}

	time.Sleep(2 * time.Second)

	for p.WorkerCount() > 0 {
		if err := p.RemoveWorker(); err != nil {
			panic(err)
		}
	}

	p.Wait()
	fmt.Println("All workers stopped. Done!")
}
