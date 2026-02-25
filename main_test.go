package pipeline_test

import (
	"errors"
	"testing"

	"github.com/filipewelton/pipeline/v3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPipeline(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pipeline Suite")
}

var _ = Describe("Pipeline With Context", func() {
	It("should update context", func() {
		ctx := 1
		p := pipeline.NewWithContext(&ctx, false)

		updatedCtx, err := p.
			Add(func(ctx *int) (int, error) {
				return *ctx + 1, nil
			}).
			Run()

		Expect(err).To(BeNil())
		Expect(ctx).ToNot(Equal(updatedCtx))
	})

	It("should stop on first error", func() {
		ctx := 1
		p := pipeline.NewWithContext(&ctx, true)

		_, err := p.
			Add(func(ctx *int) (int, error) {
				return 0, pipeline.ErrRejectedWithoutReason
			}).
			Add(func(ctx *int) (int, error) {
				return *ctx + 1, nil
			}).
			Run()

		Expect(err).To(MatchError(pipeline.ErrRejectedWithoutReason))
	})

	It("should not stop on first error", func() {
		ctx := 1
		p := pipeline.NewWithContext(&ctx, false)
		someError := errors.New("some error")
		anotherError := errors.New("another error")

		_, err := p.
			Add(func(ctx *int) (int, error) {
				return 0, pipeline.ErrRejectedWithoutReason
			}).
			Add(func(ctx *int) (int, error) {
				return *ctx, someError
			}).
			Add(func(ctx *int) (int, error) {
				return *ctx, anotherError
			}).
			Run()

		Expect(err).To(MatchError(errors.Join(
			pipeline.ErrRejectedWithoutReason,
			someError,
			anotherError,
		)))
	})
})

var _ = Describe("Pipeline Without Context", func() {
	It("should return nil error when all executors succeed", func() {
		p := pipeline.New(false)

		err := p.
			Add(func() error {
				return nil
			}).
			Run()

		Expect(err).To(BeNil())
	})

	It("should stop on first error", func() {
		p := pipeline.New(true)

		err := p.
			Add(func() error {
				return pipeline.ErrRejectedWithoutReason
			}).
			Add(func() error {
				return nil
			}).
			Run()

		Expect(err).To(MatchError(pipeline.ErrRejectedWithoutReason))
	})

	It("should not stop on first error", func() {
		p := pipeline.New(false)
		someError := errors.New("some error")
		anotherError := errors.New("another error")

		err := p.
			Add(func() error {
				return pipeline.ErrRejectedWithoutReason
			}).
			Add(func() error {
				return someError
			}).
			Add(func() error {
				return anotherError
			}).
			Run()

		Expect(err).To(MatchError(errors.Join(
			pipeline.ErrRejectedWithoutReason,
			someError,
			anotherError,
		)))
	})
})
