package task

type Type int

const (
	HttpGetRequst   Type = iota // GET-запрос HTTP
	HttpPostRequest Type = iota // POST-запрос HTTP
	FileRead        Type = iota // Чтение файла
	TextFileWrite   Type = iota // Запись текстового файла
)
