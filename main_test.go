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

var _ = Describe("Promise", func() {
	It("should execute Then and not set error if no rejection", func() {
		p := NewPromise[int]()

		p.Then(func(ctx *int, reject Reject) {
			*ctx = 42
		})

		Expect(p.Catch()).To(BeNil())
		Expect(p.context).To(Equal(42))
	})

	It("should set error if rejected", func() {
		p := NewPromise[int]()
		err := errors.New("fail")

		p.Then(func(ctx *int, reject Reject) {
			reject(err)
		})

		Expect(p.Catch()).To(MatchError("fail"))
	})

	It("should return ErrRejectedWithoutReason if rejected with nil", func() {
		p := NewPromise[int]()

		p.Then(func(ctx *int, reject Reject) {
			reject(nil)
		})

		Expect(p.Catch()).To(Equal(ErrRejectedWithoutReason))
	})
})
