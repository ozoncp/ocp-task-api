package internal

//go:generate mockgen -destination=./mocks/notifier_mock.go -package=mocks github.com/ozoncp/ocp-task-api/internal/notifier Notifier
//go:generate mockgen -destination=./mocks/alarm_mock.go -package=mocks github.com/ozoncp/ocp-task-api/internal/time Alarm
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-task-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozoncp/ocp-task-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/publisher_mock.go -package=mocks github.com/ozoncp/ocp-task-api/internal/metrics Publisher
