package pipeline

import "errors"

type Reject func(reason error)

var ErrRejectedWithoutReason = errors.New("pipeline rejected without reason")

func NewWithContext[T any](
	context *T, stopOnFirstError bool,
) *PipelineWithContext[T] {
	return &PipelineWithContext[T]{
		context:          context,
		stopOnFirstError: stopOnFirstError,
	}
}

func New(stopOnFirstError bool) *Pipeline {
	return &Pipeline{
		stopOnFirstError: stopOnFirstError,
	}
}
