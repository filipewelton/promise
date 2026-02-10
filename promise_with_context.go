package promise

type PromiseWithContext[T any] struct {
	context  T
	hasError bool
	err      error
}

type ExecutorWithContext[T any] func(ctx *T, reject Reject)

func (p *PromiseWithContext[T]) Then(
	executor ExecutorWithContext[T],
) *PromiseWithContext[T] {
	if p.hasError {
		return p
	}

	executor(&p.context, func(reason error) {
		p.hasError = true
		p.err = reason
	})

	return p
}

func (p *PromiseWithContext[T]) Catch() error {
	if !p.hasError {
		return nil
	} else if p.err == nil {
		return ErrRejectedWithoutReason
	}

	return p.err
}
