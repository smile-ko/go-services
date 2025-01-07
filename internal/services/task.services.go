package services

import (
	"go-services/internal/po"
	"go-services/internal/repo"
)

type TaskService struct {
	taskRepo *repo.TaskRepo
}

func NewTaskService() *TaskService {
	return &TaskService{taskRepo: repo.NewTaskRepo()}
}

func (ts *TaskService) Create(task *po.Task) error {

	return ts.taskRepo.Create(task)
}

func (ts *TaskService) GetAll() ([]po.Task, error) {
	return ts.taskRepo.GetAll()
}

func (ts *TaskService) GetByID(id uint) (*po.Task, error) {
	return ts.taskRepo.GetByID(id)
}

func (ts *TaskService) Update(task *po.Task) error {
	return ts.taskRepo.Update(task)
}

func (ts *TaskService) Delete(id uint) error {
	return ts.taskRepo.Delete(id)
}
