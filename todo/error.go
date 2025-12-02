package todo

import "errors"

var ErrTaskNotFound = errors.New("Task not found")
var ErrTaskAlreadyExists = errors.New("Task already exists")
