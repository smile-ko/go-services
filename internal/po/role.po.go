package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	UUID     string `gorm:"column:uuid; type:varchar(255); not null; unique;"`
	Name     string `gorm:"column:name; type:varchar(255); not null;"`
	IsActive string `gorm:"column:is_active; type:boolean; not null; default:true;"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
