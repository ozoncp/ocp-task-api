package metrics

import "context"

// Publisher интефейс публикации метрик
type Publisher interface {
	PublishFlushing(ctx context.Context, count int)
}

// NewPublisher возвращает Publisher для prometheus
func NewPublisher() Publisher {
	return &publisher{}
}

type publisher struct {
}

func (p *publisher) PublishFlushing(ctx context.Context, count int) {
}
