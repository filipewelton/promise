package promise

import "errors"

type Reject func(reason error)

var ErrRejectedWithoutReason = errors.New("promise rejected without reason")

func NewPromiseWithContext[T any](context *T) *PromiseWithContext[T] {
	var ctx T

	if context != nil {
		ctx = *context
	}

	return &PromiseWithContext[T]{
		context:  ctx,
		hasError: false,
		err:      nil,
	}
}

func NewPromise() *Promise {
	return &Promise{
		hasError: false,
		err:      nil,
	}
}
