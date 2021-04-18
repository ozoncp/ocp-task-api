package repo

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-task-api/internal/models"

	"github.com/jmoiron/sqlx"
)

var (
	errUnimplemented = errors.New("unimplemented")
)

type Repo interface {
	AddTask(ctx context.Context, task models.Task) (uint64, error)
	AddTasks(ctx context.Context, tasks []models.Task) error
	RemoveTask(taskId uint64) error
	DescribeTask(taskId uint64) (*models.Task, error)
	ListTasks(limit, offset uint64) ([]models.Task, error)
}

func NewRepo(
	db sqlx.DB,
) Repo {
	return &repo{db: db}
}

type repo struct {
	db sqlx.DB
}

func (r *repo) AddTask(ctx context.Context, task models.Task) (uint64, error) {

	return 0, errUnimplemented
}

func (r *repo) AddTasks(ctx context.Context, task []models.Task) error {
	return errUnimplemented
}

func (r *repo) RemoveTask(taskId uint64) error {

	return errUnimplemented
}

func (r *repo) DescribeTask(taskId uint64) (*models.Task, error) {

	return nil, errUnimplemented
}

func (r *repo) ListTasks(limit, offset uint64) ([]models.Task, error) {
	return nil, errUnimplemented
}
