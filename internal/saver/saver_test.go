package saver_test

import (
	"context"

	"github.com/ozoncp/ocp-task-api/internal/mocks"
	"github.com/ozoncp/ocp-task-api/internal/models"
	"github.com/ozoncp/ocp-task-api/internal/saver"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	capacity = 100
)

var _ = Describe("Saver", func() {

	var (
		err error

		ctrl *gomock.Controller
		ctx  context.Context

		mockFlusher *mocks.MockFlusher
		mockAlarm   *mocks.MockAlarm

		task models.Task
		s    saver.Saver

		alarms chan struct{}
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockAlarm = mocks.NewMockAlarm(ctrl)
		mockFlusher = mocks.NewMockFlusher(ctrl)

		alarms = make(chan struct{})
		mockAlarm.EXPECT().Alarm().Return(alarms).AnyTimes()

		s = saver.NewSaver(capacity, mockAlarm, mockFlusher)
	})

	JustBeforeEach(func() {
		s.Init(ctx)
		err = s.Save(ctx, task)
	})

	AfterEach(func() {
		s.Close()
		ctrl.Finish()
	})

	Context("ctx canceled", func() {

		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil)
		})

		JustBeforeEach(func() {
			cancelFunc()
		})

		It("", func() {
			Expect(err).Should(BeNil())
		})
	})

	Context("alarm is occurring", func() {

		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1).MaxTimes(2)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
			cancelFunc()
		})

		It("", func() {
			Expect(err).Should(BeNil())
		})
	})
})
