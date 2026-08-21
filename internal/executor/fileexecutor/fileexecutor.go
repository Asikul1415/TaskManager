package fileexecutor

import (
	"TaskManager/internal/task"
	"log"
	"os"
)

// Исполнитель задач, связанных с работой с фалайми
type FileExecutor struct {
}

func (e *FileExecutor) Execute(t *task.Task) {
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
