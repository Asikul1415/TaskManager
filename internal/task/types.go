package task

type Type int

const (
	HttpGetRequst   Type = iota // GET-запрос HTTP
	HttpPostRequest Type = iota // POST-запрос HTTP
)
