# Worker Pool
VK test task

___by_MishMish___

## Архитектура проекта (tree)

```text
.
├── cmd/
│   └── main.go       — пример использования пула воркеров
├── pool/
│   └── pool.go       — реализация пула воркеров
├── worker/
│   └── worker.go     — реализация воркера
├── go.mod            — зависимости Go
└── README.md         — описание проекта (этот файл)
```

## Установка и запуск

```shell
git clone https://github.com/gramrate/VK-WorkerPool.git
cd WorkerPool
go run cmd/main.go
```

## Примеры использования

### 1) Создание пула:

```go
p := pool.NewPool()
```

### 2) Добавление воркеров:

```go
p.AddWorker()
p.AddWorker()
```

### 3) Отправка заданий:

```go
p.Submit("Hello world!")
```

### 4) Удаление воркеров:

```go
if err := p.RemoveWorker(); err != nil {
panic(err)
}
```

### 5) Завершение работы:

```go
p.Wait()
```