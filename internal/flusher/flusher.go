package flusher

import (
	"context"

	"github.com/ozoncp/ocp-task-api/internal/metrics"
	"github.com/ozoncp/ocp-task-api/internal/models"
	"github.com/ozoncp/ocp-task-api/internal/repo"
)

// Flusher - интерфейс для сброса задач в хранилище
type Flusher interface {
	Flush(ctx context.Context, tasks []models.Task) []models.Task
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	taskRepo repo.Repo,
	publisher metrics.Publisher,
) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		taskRepo:  taskRepo,
		publisher: publisher,
	}
}

type flusher struct {
	chunkSize int
	taskRepo  repo.Repo
	publisher metrics.Publisher
}

func (f *flusher) Flush(ctx context.Context, tasks []models.Task) []models.Task {

	for i := 0; i < len(tasks); i += f.chunkSize {
		j := i + f.chunkSize
		if j >= len(tasks) {
			j = len(tasks)
		}

		chunk := tasks[i:j]
		if err := f.taskRepo.AddTasks(ctx, chunk); err != nil {
			return tasks[i:]
		}
		f.publisher.PublishFlushing(ctx, len(chunk))
	}
	return nil
}
