package executor

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/bitly/go-simplejson"

	"TaskManager/internal/task"
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
		path, isField := t.Args["path"]
		if !isField {
			log.Printf(" Task.Run: поле `path` отсутствует в параметрах задачи")
			return
		}

		var jsonMock *simplejson.Json = simplejson.New()
		httpPostRequest(path, jsonMock)
		return
	}

	if t.Type == task.FileRead {
		filePath, isField := t.Args["path"]
		if !isField {
			log.Printf(" Task.Run: поле `path` отсутствует в параметрах задачи")
			return
		}

		fileText := readFile(filePath)
		log.Printf("Задача %d вернула текст файла: %s", t.Id, fileText)
		return
	}

	if t.Type == task.TextFileWrite {
		fileName, isField := t.Args["filename"]
		if !isField {
			log.Printf(" Task.Run: поле `filename` отсутствует в параметрах задачи")
			return
		}

		text, isField := t.Args["text"]
		if !isField {
			log.Printf(" Task.Run: поле `text` отсутствует в параметрах задачи")
			return
		}

		writeTextFile(text, fileName)
	}
}

// Выполняет GET-запрос HTTP по пути path. Возвращает ответ на GET-запрос.
func httpGetRequest(path string) *http.Response {
	resp, err := http.Get(path)
	if err != nil {
		log.Printf(" executor.httpGetRequest: ошибка при GET-запросе к %s", path)
		log.Printf(" executor.httpGetRequest: %s", err)
		return &http.Response{StatusCode: 400}
	}

	return resp
}

// Выполняет POST-запрос HTTP по пути path, с JSON полученным из json. Возвращает ответ на POST-запрос.
func httpPostRequest(path string, json *simplejson.Json) *http.Response {
	jsonBytes, err := json.Bytes()
	if err != nil {
		log.Printf(" executor.httpPostRequest: ошибка при POST-запросе к %s. Не удалось преобразовать входящий JSON в массив байт", path)
		log.Printf(" executor.httpPostRequest: %s", err)
		return &http.Response{StatusCode: 400}
	}

	body := bytes.NewReader(jsonBytes)
	response, err := http.Post(path, "applications/json", body)
	if err != nil {
		log.Printf(" executor.httpPostRequest: ошибка при POST-запросе к %s", path)
		log.Printf(" executor.httpPostRequest: %s", err)
		return &http.Response{StatusCode: 400}
	}
	return response
}

// Читает файл по пути filePath. Возвращает текст файла.
func readFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf(" executor.ReadFile: ошибка при чтении файла %s", filePath)
		log.Printf(" executor.ReadFile: %s", err)
		return ""
	}

	text := string(data)
	return text
}

// Записывает в текстовый файл fileName текст text
func writeTextFile(text string, fileName string) {
	textBytes := []byte(text)
	err := os.WriteFile(fileName, textBytes, os.ModeAppend)
	if err != nil {
		log.Printf(" executor.WriteTextFile: ошибка при записи текстового файла %s", fileName)
		log.Printf(" executor.WriteTextFile: %s", err)
		return
	}

	log.Printf(" executor.WriteTextFile: текстовый файл %s успешно записан", fileName)
}
