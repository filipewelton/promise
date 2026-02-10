package promise

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPromise(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Promise Suite")
}

var _ = Describe("Promise With Context", func() {
	It("should update context", func() {
		p := NewPromiseWithContext[int](nil)

		err := p.
			ThenWithContext(func(ctx *int, reject Reject) {
				*ctx = 42
			}).
			ThenWithContext(func(ctx *int, reject Reject) {
				*ctx = 24
			}).
			Catch()

		Expect(err).To(BeNil())
		Expect(p.context).To(Equal(24))
	})

	It("should set error if rejected", func() {
		p := NewPromiseWithContext[int](nil)
		err := errors.New("fail")

		p.ThenWithContext(func(ctx *int, reject Reject) {
			reject(err)
		})

		Expect(p.Catch()).To(MatchError("fail"))
	})

	It("should return ErrRejectedWithoutReason if rejected with nil", func() {
		p := NewPromiseWithContext[int](nil)

		p.ThenWithContext(func(ctx *int, reject Reject) {
			reject(nil)
		})

		Expect(p.Catch()).To(Equal(ErrRejectedWithoutReason))
	})
})

var _ = Describe("Promise Without Context", func() {
	It("should execute all handlers", func() {
		var x, y int

		err := NewPromise().
			Then(func(reject Reject) {
				x = 10
			}).
			Then(func(reject Reject) {
				y = 20
			}).
			Catch()

		Expect(err).Should(BeNil())
		Expect(x).Should(Equal(10))
		Expect(y).Should(Equal(20))
	})

	It("should stop on first error without reason", func() {
		err := NewPromise().
			Then(func(reject Reject) {
				reject(nil)
			}).
			Catch()

		Expect(err).Should(Equal(ErrRejectedWithoutReason))
	})

	It("should stop on first error with reason", func() {
		err := NewPromise().
			Then(func(reject Reject) {
				reject(errors.New("fail"))
			}).
			Catch()

		Expect(err).Should(MatchError("fail"))
	})
})
