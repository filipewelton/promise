package pipeline

import (
	"errors"
)

type PipelineWithContext[T any] struct {
	context          *T
	executors        []ExecutorWithContext[T]
	stopOnFirstError bool
	errors           []error
}

type ExecutorWithContext[T any] func(ctx *T) (T, error)

func (p *PipelineWithContext[T]) Add(
	executor ExecutorWithContext[T],
) *PipelineWithContext[T] {
	p.executors = append(p.executors, executor)
	return p
}

func (p *PipelineWithContext[T]) Run() (T, error) {
	for _, executor := range p.executors {
		ctx, err := executor(p.context)

		if err == nil {
			p.context = &ctx
			continue
		} else if p.stopOnFirstError {
			return *p.context, err
		}

		p.errors = append(p.errors, err)
	}

	return *p.context, errors.Join(p.errors...)
}
