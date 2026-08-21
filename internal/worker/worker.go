package worker

import (
	"TaskManager/internal/executor"
	"TaskManager/internal/executor/fileexecutor"
	"TaskManager/internal/executor/httpexecutor"
	"TaskManager/internal/task"
)

func Run(t *task.Task) {
	var executor executor.Executor

	if t.Type == task.HttpGetRequst || t.Type == task.HttpPostRequest {
		executor = &httpexecutor.HttpExecutor{}
	}

	if t.Type == task.FileRead || t.Type == task.TextFileWrite {
		executor = &fileexecutor.FileExecutor{}
	}

	executor.Execute(t)
}
