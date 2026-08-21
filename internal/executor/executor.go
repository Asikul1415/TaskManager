package executor

import (
	"TaskManager/internal/task"
)

// Интерфейс исполнителя задач
type Executor interface {
	Execute(t *task.Task)
}
