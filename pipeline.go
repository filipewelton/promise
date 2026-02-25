package pipeline

import "errors"

type Pipeline struct {
	executors        []Executor
	stopOnFirstError bool
	err              []error
}

type Executor func() error

func (p *Pipeline) Add(executor Executor) *Pipeline {
	p.executors = append(p.executors, executor)
	return p
}

func (p *Pipeline) Run() error {
	for _, executor := range p.executors {
		err := executor()

		if err == nil {
			continue
		} else if p.stopOnFirstError {
			return err
		}

		p.err = append(p.err, err)
	}

	return errors.Join(p.err...)
}
