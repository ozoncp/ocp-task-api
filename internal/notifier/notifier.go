package notifier

import (
	"context"

	"github.com/ozoncp/ocp-lib/pkg/tasks"
)

// Notifier
type Notifier interface {
	Notify(ctx context.Context, taskId uint64, event tasks.Event)
}
