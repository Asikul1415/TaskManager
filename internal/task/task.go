package task

type Task struct {
	Id   int
	Type Type
	Args map[string]string
}
