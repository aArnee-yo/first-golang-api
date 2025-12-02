package app

import (
	"encoding/json"
	"errors"
	"time"
)

type CompleteDTO struct {
	Complete bool
}

type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) EmptyField() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Description == "" {
		return errors.New("descrption is empty")
	}
	return nil
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	str, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(str)
}
