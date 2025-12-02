package todo

import "sync"

type List struct {
	Tasks map[string]Task
	mtx   sync.RWMutex
}

func NewList() *List {
	return &List{
		Tasks: make(map[string]Task),
	}
}

func (l *List) AddTask(task Task) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	if _, ok := l.Tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}

	l.Tasks[task.Title] = task
	return nil
}

func (l *List) GetTask(title string) (Task, error) {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	task, ok := l.Tasks[title]

	if !ok {
		return Task{}, ErrTaskAlreadyExists
	}

	return task, nil
}

func (l *List) GetTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	tmp := make(map[string]Task, len(l.Tasks))

	for k, v := range l.Tasks {
		tmp[k] = v
	}

	return tmp
}

func (l *List) GetUncompleteTasks() map[string]Task {
	l.mtx.RLock()
	defer l.mtx.RUnlock()

	uncompletedTasks := make(map[string]Task)

	for title, task := range l.Tasks {
		if !task.Complet {
			uncompletedTasks[title] = task
		}
	}

	return uncompletedTasks
}

func (l *List) CompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.Tasks[title]

	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Complete()

	l.Tasks[title] = task

	return task, nil
}

func (l *List) UncompleteTask(title string) (Task, error) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	task, ok := l.Tasks[title]

	if !ok {
		return Task{}, ErrTaskNotFound
	}

	task.Uncomplete()

	l.Tasks[title] = task

	return task, nil
}

func (l *List) DeleteTask(title string) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	_, ok := l.Tasks[title]
	if !ok {
		return ErrTaskNotFound
	}

	delete(l.Tasks, title)

	return nil
}
