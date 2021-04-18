package repo

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-task-api/internal/models"

	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "tasks"
)

var (
	errUnimplemented = errors.New("unimplemented")
)

type Repo interface {
	AddTask(ctx context.Context, task models.Task) (uint64, error)
	AddTasks(ctx context.Context, tasks []models.Task) error
	RemoveTask(ctx context.Context, taskId uint64) error
	DescribeTask(ctx context.Context, taskId uint64) (*models.Task, error)
	ListTasks(ctx context.Context, limit, offset uint64) ([]models.Task, error)
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

	query := sq.Insert(tableName).
		Columns("description").
		Values(task.Description).
		Suffix("RETURNING \"id\"").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	query.QueryRow().Scan(&task.Id)

	return task.Id, nil
}

func (r *repo) AddTasks(ctx context.Context, tasks []models.Task) error {

	query := sq.Insert(tableName).
		Columns("description").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, task := range tasks {

		query = query.Values(task.Description)
	}

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) RemoveTask(ctx context.Context, taskId uint64) error {

	query := sq.Delete(tableName).
		Where(sq.Eq{"id": taskId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) DescribeTask(ctx context.Context, taskId uint64) (*models.Task, error) {

	query := sq.Select("id", "description").
		From(tableName).
		Where(sq.Eq{"id": taskId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var task models.Task
	if err := query.QueryRowContext(ctx).Scan(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *repo) ListTasks(ctx context.Context, limit, offset uint64) ([]models.Task, error) {

	query := sq.Select("id", "description").
		From(tableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.Id,
			&task.Description,
		); err != nil {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
