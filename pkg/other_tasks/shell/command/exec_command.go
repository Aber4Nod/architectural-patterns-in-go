package command

import (
	"context"
	"sync"

	errorTypes "github.com/aber4nod/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

type execCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (e *execCommand) Execute(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := exec(e.args[0], e.args[1:]); err != nil {
		e.errorChannel <- err
	}
}

func (e *execCommand) SetArgs(args []string) error {
	if len(args) == 0 {
		return errorTypes.ErrorNotEnoughArguments
	}
	e.args = args
	return nil
}

func NewExecCommand(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &execCommand{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
