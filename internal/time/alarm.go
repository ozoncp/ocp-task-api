package time

import (
	"context"
	"time"
)

type Alarm interface {
	Alarm() <-chan struct{}
	Init()
}

func NewAlarm(
	ctx context.Context,
	timeout time.Duration,
	clock Clock,
) Alarm {

	chanel := make(chan struct{})

	return &alarm{
		ctx:     ctx,
		timeout: timeout,
		clock:   clock,
		chanel:  chanel,
	}
}

type alarm struct {
	ctx     context.Context
	timeout time.Duration
	clock   Clock
	chanel  chan struct{}
}

func (a *alarm) Alarm() <-chan struct{} {
	return a.chanel
}

func (a *alarm) Init() {
	go func() {
		timer := time.After(a.timeout)
		for {

			select {
			case <-timer:
				a.chanel <- struct{}{}
			case <-a.ctx.Done():
				close(a.chanel)
				return
			}
		}
	}()
}
