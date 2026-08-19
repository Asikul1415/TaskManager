package main

import (
	"log"
)

func main() {
	type Task struct {
		Id          int
		Type        string
		Description string
	}

	var queue []Task = []Task{
		{Id: 0, Type: "http-request", Description: "path: https://google.com"},
		{Id: 1, Type: "http-request", Description: "path: https://youtube.com"},
		{Id: 2, Type: "http-request", Description: "path: https://wikipedia.com"},
	}

	for _, task := range queue {
		log.Printf("Задача id %d была успешно выполнена", task.Id)
	}
}
