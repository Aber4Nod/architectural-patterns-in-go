package command

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	errorTypes "github.com/aber4nod/architectural-patterns-in-go/pkg/other_tasks/shell/errors"
)

const (
	cdCommandTestCaseName            = "Change working directory"
	cdCommandNoArgsTestCaseName      = "No args"
	cdCommandTooManyArgsTestCaseName = "Too many args"
	cdCommandChdirErrorTestCaseName  = "Error in chdir func"
)

var chdirError = errors.New("error in chdir")

type cdCommandInputData struct {
	args            []string
	pathInChdirFunc string
	chdirError      error
}

type cdCommandExpectedResult struct {
	setArgsError       error
	chdirNumberOfCalls int
	errorChannelError  error
}

func Test_CdCommand(t *testing.T) {
	for _, testData := range []struct {
		testCaseName   string
		inputData      cdCommandInputData
		expectedResult cdCommandExpectedResult
	}{
		{
			testCaseName: cdCommandTestCaseName,
			inputData: cdCommandInputData{
				args:            []string{"new_working_directory"},
				pathInChdirFunc: "new_working_directory",
			},
			expectedResult: cdCommandExpectedResult{
				chdirNumberOfCalls: 1,
				errorChannelError:  nil,
			},
		},
		{
			testCaseName: cdCommandNoArgsTestCaseName,
			inputData: cdCommandInputData{
				args:            []string{},
				pathInChdirFunc: "",
			},
			expectedResult: cdCommandExpectedResult{
				chdirNumberOfCalls: 0,
			},
		},
		{
			testCaseName: cdCommandTooManyArgsTestCaseName,
			inputData: cdCommandInputData{
				args:            []string{"new_working_directory", "another_arg"},
				pathInChdirFunc: "new_working_directory",
			},
			expectedResult: cdCommandExpectedResult{
				setArgsError:       errorTypes.ErrorTooManyArguments,
				chdirNumberOfCalls: 0,
			},
		},
		{
			testCaseName: cdCommandChdirErrorTestCaseName,
			inputData: cdCommandInputData{
				args:            []string{"new_working_directory"},
				pathInChdirFunc: "new_working_directory",
				chdirError:      chdirError,
			},
			expectedResult: cdCommandExpectedResult{
				chdirNumberOfCalls: 1,
				errorChannelError:  chdirError,
			},
		},
	} {
		t.Run(testData.testCaseName, func(t *testing.T) {
			inputChannel := make(chan string)
			outputChannel := make(chan string)
			errorChannel := make(chan error)
			cdCommand := NewCdCommand(
				inputChannel,
				outputChannel,
				errorChannel,
			)
			err := cdCommand.SetArgs(testData.inputData.args)
			if testData.expectedResult.setArgsError != nil {
				assert.ErrorIs(t, err, testData.expectedResult.setArgsError)
			}

			originalChdir := chdir
			chdirCallsNumber := 0
			chdir = func(path string) error {
				chdirCallsNumber++
				assert.Equal(t, testData.inputData.pathInChdirFunc, path)
				return testData.inputData.chdirError
			}

			mainCtx := context.Background()
			ctx, _ := context.WithCancel(mainCtx)
			var wg sync.WaitGroup
			wg.Add(1)
			go cdCommand.Execute(ctx, &wg)
			var resultError error
			if testData.expectedResult.errorChannelError != nil {
				resultError = <-errorChannel
			}
			wg.Wait()
			chdir = originalChdir

			assert.Equal(t, testData.expectedResult.errorChannelError, resultError)
			assert.Equal(t, testData.expectedResult.chdirNumberOfCalls, chdirCallsNumber)
		})
	}
}
