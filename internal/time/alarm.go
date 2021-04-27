package time

import (
	"context"
	"time"
)

type Alarm interface {
	Alarm() <-chan struct{}
	Init()
	Close()
}

func NewAlarm(
	ctx context.Context,
	timeout time.Duration,
	clock Clock,
) Alarm {

	alarms := make(chan struct{})
	done := make(chan struct{})

	return &alarm{
		ctx:     ctx,
		timeout: timeout,
		clock:   clock,
		alarms:  alarms,
		done:    done,
	}
}

type alarm struct {
	ctx     context.Context
	timeout time.Duration
	clock   Clock
	alarms  chan struct{}
	done    chan struct{}
}

func (a *alarm) Alarm() <-chan struct{} {
	return a.alarms
}

func (a *alarm) Init() {
	go func() {
		timer := time.After(a.timeout)
		for {

			select {
			case <-timer:
				a.alarms <- struct{}{}
				timer = time.After(a.timeout)
			case <-a.ctx.Done():
				close(a.alarms)
				a.done <- struct{}{}
				return
			}
		}
	}()
}

func (a *alarm) Close() {
	<-a.done
}
