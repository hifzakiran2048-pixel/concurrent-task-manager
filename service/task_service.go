package service

import (
	"errors"
	"mymodule/model"
	"mymodule/store"
)

var idCounter = 1

func AddTask(name string) model.Task {
	task := model.Task{

		ID:     idCounter,
		Name:   name,
		Status: "Pending",
	}

	store.Task = append(store.Task, task)
	idCounter++

	return task
}
func GetTask() []model.Task {
	return store.Task
}

func GetTaskByID(id int) (*model.Task, error) {
	for i := range store.Task {
		if store.Task[i].ID == id {
			return &store.Task[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func MarkDone(id int) {
	for i := range store.Task {
		if store.Task[i].ID == id {
			store.Task[i].Status = "Done"
			return
		}
	}
}

func UpdateTask(id int, newName string) error {
	for i := range store.Task {
		if store.Task[i].ID == id {

			if store.Task[i].Status == "Done" {
				return errors.New("cannot update completed task")
			}

			store.Task[i].Name = newName
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id int) error {
	for i := range store.Task {
		if store.Task[i].ID == id {

			store.Task = append(store.Task[:i], store.Task[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
