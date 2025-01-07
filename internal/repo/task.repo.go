package repo

import (
	"go-services/global"
	"go-services/internal/po"

	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{db: global.PDB}
}

func (tr *TaskRepo) Create(task *po.Task) error {
	return tr.db.Create(task).Error
}

func (tr *TaskRepo) GetAll() ([]po.Task, error) {
	var tasks []po.Task
	err := tr.db.Find(&tasks).Error
	return tasks, err
}

func (tr *TaskRepo) GetByID(id uint) (*po.Task, error) {
	var task po.Task
	err := tr.db.First(&task, id).Error
	return &task, err
}

func (tr *TaskRepo) Update(task *po.Task) error {
	return tr.db.Save(task).Error
}

func (tr *TaskRepo) Delete(id uint) error {
	return tr.db.Delete(&po.Task{}, id).Error
}
