package command

import (
	"context"
	"sync"

	errorTypes "github.com/aber4nod/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

type forkCommand struct {
	args          []string
	inputChannel  <-chan string
	outputChannel chan<- string
	errorChannel  chan<- error
}

func (f *forkCommand) Execute(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	//if _, err := fork(f.args[0], f.args[1:]); err != nil {
	//	f.errorChannel <- err
	//}
}

func (f *forkCommand) SetArgs(args []string) error {
	if len(args) == 0 {
		return errorTypes.ErrorNotEnoughArguments
	}
	f.args = args
	return nil
}

func NewForkCommand(
	inputChannel <-chan string,
	outputChannel chan<- string,
	errorChannel chan<- error,
) Command {
	return &forkCommand{
		inputChannel:  inputChannel,
		outputChannel: outputChannel,
		errorChannel:  errorChannel,
	}
}
