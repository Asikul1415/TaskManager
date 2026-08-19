package main

import (
	"TaskManager/internal/executor"
	"TaskManager/internal/task"
	"log"
)

func main() {
	var queue []task.Task = []task.Task{
		{Id: 0, Type: task.HttpGetRequst, Args: map[string]string{"path": "https://wikipedia.com"}},
		{Id: 1, Type: task.HttpGetRequst, Args: map[string]string{"path": "https://youtube.com"}},
		{Id: 2, Type: task.HttpGetRequst, Args: map[string]string{"path": "https://google.com"}},
	}

	for _, task := range queue {
		executor.Execute(&task)
		log.Printf("Задача id %d была успешно выполнена", task.Id)
	}
}
