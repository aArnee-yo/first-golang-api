package app

import (
	"encoding/json"
	"errors"
	"firstapi/todo"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Handler struct {
	list *todo.List
}

func NewHandler(list *todo.List) *Handler {
	return &Handler{
		list: list,
	}
}

func (h *Handler) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var task TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		err_msg := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, err_msg.ToString(), http.StatusBadRequest)
		return
	}

	if err := task.EmptyField(); err != nil {
		err_msg := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, err_msg.ToString(), http.StatusBadRequest)
		return
	}

	t := todo.NewTask(task.Title, task.Description)
	if err := h.list.AddTask(t); err != nil {
		err_msg := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, err_msg.ToString(), http.StatusConflict)
		} else {
			http.Error(w, err_msg.ToString(), http.StatusInternalServerError)
		}
		return
	}
	resp, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(resp); err != nil {
		fmt.Println(err)
		return
	}
}

func (h *Handler) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	task, err := h.list.GetTask(title)
	if err != nil {
		err_msg := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, err_msg.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, err_msg.ToString(), http.StatusInternalServerError)
		}

		return
	}

	resp, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		fmt.Println(err)
		return
	}
}

func (h *Handler) HandleGetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.list.GetTasks()
	resp, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		fmt.Println(err)
		return
	}
}

func (h *Handler) HandleGetUncomleteTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.list.GetUncompleteTasks()
	resp, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		fmt.Println(err)
		return
	}
}

func (h *Handler) HandleComleteTask(w http.ResponseWriter, r *http.Request) {
	var compl CompleteDTO
	if err := json.NewDecoder(r.Body).Decode(&compl); err != nil {
		err_msg := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, err_msg.ToString(), http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]

	var (
		changedTask todo.Task
		err         error
	)

	if compl.Complete {
		changedTask, err = h.list.CompleteTask(title)
	} else {
		changedTask, err = h.list.UncompleteTask(title)
	}

	if err != nil {
		err_msg := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, err_msg.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, err_msg.ToString(), http.StatusInternalServerError)
		}

		return
	}

	resp, err := json.MarshalIndent(changedTask, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		fmt.Println(err)
		return
	}

}

func (h *Handler) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if err := h.list.DeleteTask(title); err != nil {
		err_msg := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, err_msg.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, err_msg.ToString(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
