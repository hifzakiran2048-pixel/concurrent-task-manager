package service

import (
	"mymodule/model"
	"mymodule/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}
func (s *TaskService) AddTask(name string) error {
	return s.repo.Create(model.Task{Name: name})
}

func (s *TaskService) ListTasks() ([]model.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) MarkDone(id int) error {
	return s.repo.Update(id)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
