package promise

type Promise struct {
	hasError bool
	err      error
}

type Executor func(reject Reject)

func (p *Promise) Then(executor Executor) *Promise {
	if p.hasError {
		return p
	}

	executor(func(reason error) {
		p.hasError = true
		p.err = reason
	})

	return p
}

func (p *Promise) Catch() error {
	if !p.hasError {
		return nil
	} else if p.err == nil {
		return ErrRejectedWithoutReason
	}

	return p.err
}
