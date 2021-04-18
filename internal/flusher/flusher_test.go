package flusher_test

import (
	"context"
	"errors"

	"github.com/ozoncp/ocp-task-api/internal/flusher"
	"github.com/ozoncp/ocp-task-api/internal/mocks"
	"github.com/ozoncp/ocp-task-api/internal/models"

	"github.com/golang/mock/gomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	errDeadlineExceeded = errors.New("unimplemented")
)

var _ = Describe("Flusher", func() {

	var (
		err error

		ctrl *gomock.Controller
		ctx  context.Context

		mockRepo      *mocks.MockRepo
		mockPublisher *mocks.MockPublisher

		tasks []models.Task
		rest  []models.Task

		f flusher.Flusher

		chunkSize int
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockRepo = mocks.NewMockRepo(ctrl)
		mockPublisher = mocks.NewMockPublisher(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepo, mockPublisher)
		rest = f.Flush(ctx, tasks)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("repo save all tasks", func() {

		BeforeEach(func() {
			chunkSize = 2
			tasks = []models.Task{{}}

			mockRepo.EXPECT().AddTasks(gomock.Any(), gomock.Any()).Return(nil).MinTimes(1)
			mockPublisher.EXPECT().PublishFlushing(gomock.Any(), len(tasks))
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeNil())
		})
	})

	Context("repo don't save tasks", func() {

		BeforeEach(func() {
			chunkSize = 2
			tasks = []models.Task{{}, {}}

			mockRepo.EXPECT().AddTasks(gomock.Any(), gomock.Any()).Return(errDeadlineExceeded)
			mockPublisher.EXPECT().PublishFlushing(gomock.Any(), gomock.Any()).Times(0)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeEquivalentTo(tasks))
		})
	})

	Context("repo save half tasks", func() {

		var (
			halfSize int
		)

		BeforeEach(func() {
			tasks = []models.Task{{}, {}}
			halfSize = int(len(tasks) / 2)
			chunkSize = halfSize

			mockRepo.EXPECT().AddTasks(gomock.Any(), gomock.Any()).Return(nil)
			mockRepo.EXPECT().AddTasks(gomock.Any(), gomock.Any()).Return(errDeadlineExceeded)
			mockPublisher.EXPECT().PublishFlushing(gomock.Any(), halfSize)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(rest).Should(BeEquivalentTo(tasks[halfSize:]))
		})
	})
})
