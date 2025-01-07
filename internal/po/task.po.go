package po

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"column:title; type:varchar(255); not null"`
	Description string `gorm:"column:description; type:text; not null"`
	Completed   bool   `gorm:"column:completed; type:boolean; not null; default:false"`
}

func (t *Task) TableName() string {
	return "go_db_task"
}
