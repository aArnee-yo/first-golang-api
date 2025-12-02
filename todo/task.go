package todo

import "time"

type Task struct {
	Title       string
	Description string
	Complet     bool

	CreatAt   time.Time
	CompletAt *time.Time
}

func NewTask(title string, description string) Task {
	return Task{
		Title:       title,
		Description: description,
		Complet:     false,

		CreatAt:   time.Now(),
		CompletAt: nil,
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()

	t.Complet = true
	t.CompletAt = &completeTime
}

func (t *Task) Uncomplete() {
	t.Complet = false
	t.CompletAt = nil
}
