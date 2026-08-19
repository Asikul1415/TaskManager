package executor

import (
	"TaskManager/internal/task"
	"log"
	"net/http"
)

// Выполняет данную задачу task
func Execute(t *task.Task) {
	if t.Type == task.HttpGetRequst {
		path, isField := t.Args["path"]
		if !isField {
			log.Printf(" Task.Run: поле `path` отсутствует в параметрах задачи")
			return
		}

		httpGetRequest(path)
		return
	}

	if t.Type == task.HttpPostRequest {
		return
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
