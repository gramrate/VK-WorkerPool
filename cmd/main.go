/*
Пример использования пула воркеров (WorkerPool).

1. Создаётся пул воркеров.
2. Добавляются два воркера.
3. Отправляются задания (строки "Job-1" .. "Job-5").
4. Удаляется один воркер.
5. Добавляется новый воркер.
6. Отправляются новые задания ("Job-6" .. "Job-100").
7. Завершаются все воркеры, пул ждёт их завершения.
*/
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
