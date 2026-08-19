package main

import (
	"log"
	"net/http"
)

type Task struct {
	Id   int
	Type string
	Args map[string]string
}

func main() {
	var queue []Task = []Task{
		{Id: 0, Type: "http-request", Args: map[string]string{"path": "https://wikipedia.com"}},
		{Id: 1, Type: "http-request", Args: map[string]string{"path": "https://youtube.com"}},
		{Id: 2, Type: "http-request", Args: map[string]string{"path": "https://google.com"}},
	}

	for _, task := range queue {
		task.Run()
		log.Printf("Задача id %d была успешно выполнена", task.Id)
	}
}

// Выполняет данную задачу
func (t *Task) Run() {
	if t.Type == "http-request" {
		path, isField := t.Args["path"]
		if !isField {
			log.Printf(" Task.Run: поле `path` отсутствует в параметрах задачи")
			return
		}

		httpGetRequest(path)
	}
}

// Выполняет GET-запрос HTTP по пути path. Возвращает ответ на GET-запрос.
func httpGetRequest(path string) *http.Response {
	resp, err := http.Get(path)
	if err != nil {
		log.Printf(" main.httpGetRequest: ошибка при GET-запросе к %s", path)
		log.Printf(" main.httpGetRequest: %s", err)
		return &http.Response{StatusCode: 400}
	}

	return resp
}
