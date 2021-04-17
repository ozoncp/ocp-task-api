package notifier

import (
	"context"
)

// Notifier
type Notifier interface {
	Notify(ctx context.Context, taskId uint64, event tasks.Event)
}
