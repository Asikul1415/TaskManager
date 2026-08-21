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
		{Id: 3, Type: task.HttpPostRequest, Args: map[string]string{"path": "https://google.com"}},
		{Id: 4, Type: task.FileRead, Args: map[string]string{"path": "./test_file.json"}},
		{Id: 5, Type: task.TextFileWrite, Args: map[string]string{"filename": "test_write.txt", "text": "something about Lorem ipsum"}},
	}

	for _, task := range queue {
		executor.Execute(&task)
		log.Printf("Задача id %d была успешно выполнена", task.Id)
	}
}
