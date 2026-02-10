package promise

import "errors"

type Promise[T any] struct {
	context  T
	hasError bool
	err      error
}

type Reject func(reason error)

type Executor[T any] func(ctx *T, reject Reject)

var ErrRejectedWithoutReason = errors.New("promise rejected without reason")

func NewPromise[T any]() *Promise[T] {
	var ctx T

	return &Promise[T]{
		context:  ctx,
		hasError: false,
		err:      nil,
	}
}

func (p *Promise[T]) Then(
	executor Executor[T],
) *Promise[T] {
	if p.hasError {
		return p
	}

	executor(&p.context, func(reason error) {
		p.hasError = true
		p.err = reason
	})

	return p
}

func (p *Promise[T]) Catch() error {
	if !p.hasError {
		return nil
	} else if p.err == nil {
		return ErrRejectedWithoutReason
	}

	return p.err
}
