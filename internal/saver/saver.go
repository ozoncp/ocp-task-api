package saver

import (
	"context"

	"github.com/ozoncp/ocp-task-api/internal/flusher"
	"github.com/ozoncp/ocp-task-api/internal/models"
	"github.com/ozoncp/ocp-task-api/internal/time"
)

type Saver interface {
	Init(ctx context.Context)
	Save(ctx context.Context, task models.Task) error
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(
	capacity int,
	alarm time.Alarm,
	flusher flusher.Flusher,
) Saver {

	tasks := make(chan models.Task, capacity)

	return &saver{
		tasks:   tasks,
		alarm:   alarm,
		flusher: flusher,
	}
}

type saver struct {
	tasks   chan models.Task
	alarm   time.Alarm
	flusher flusher.Flusher
}

func (s *saver) Init(ctx context.Context) {
	go s.flushing(ctx)
}

func (s *saver) Save(ctx context.Context, task models.Task) error {
	s.tasks <- task
	return nil
}

func (s *saver) flushing(ctx context.Context) {

	var tasks []models.Task

	for {
		select {
		case task := <-s.tasks:
			tasks = append(tasks, task)

		case <-ctx.Done():
			tasks = s.flusher.Flush(ctx, tasks)
			return

		case <-s.alarm.Alarm():
			tasks = s.flusher.Flush(ctx, tasks)
		}
	}
}
