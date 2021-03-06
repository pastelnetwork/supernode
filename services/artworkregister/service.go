package artworkregister

import (
	"context"
	"sync"

	"github.com/pastelnetwork/go-pastel"
	"github.com/pastelnetwork/supernode/storage"
)

// const logPrefix = "[artwork]"

// Service represent artwork service.
type Service struct {
	sync.Mutex

	config *Config
	db     storage.KeyValue
	pastel pastel.Client
	worker *Worker
	tasks  []*Task
}

// Run starts worker
func (service *Service) Run(ctx context.Context) error {
	return service.worker.Run(ctx)
}

// Task returns the task of the registration artwork by the given connID.
func (service *Service) Task(connID string) *Task {
	for _, task := range service.tasks {
		if task.ConnID == connID {
			return task
		}
	}
	return nil
}

// NewTask runs a new task of the registration artwork and returns its taskID.
func (service *Service) NewTask(ctx context.Context, connID string) *Task {
	service.Lock()
	defer service.Unlock()

	task := NewTask(service)
	service.tasks = append(service.tasks, task)
	service.worker.AddTask(ctx, task)

	return task
}

// NewService returns a new Service instance.
func NewService(config *Config, db storage.KeyValue, pastel pastel.Client) *Service {
	return &Service{
		config: config,
		db:     db,
		pastel: pastel,
		worker: NewWorker(),
	}
}
