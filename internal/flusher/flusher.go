package flusher

import (
	"context"

	"github.com/ozoncp/ocp-task-api/internal/models"
	"github.com/ozoncp/ocp-task-api/internal/repo"
)

type Flusher interface {
	Flush(ctx context.Context, tasks []models.Task) []models.Task
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	taskRepo repo.Repo,
) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		taskRepo:  taskRepo,
	}
}

type flusher struct {
	chunkSize int
	taskRepo  repo.Repo
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
	}
	return nil
}
