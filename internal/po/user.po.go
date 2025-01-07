package po

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string `gorm:"column:uuid; type:varchar(255); not null; unique;"`
	Username string `gorm:"column:user_name; type:varchar(255); not null;"`
	Password string `gorm:"column:password; type:varchar(255); not null;"`
	IsActive string `gorm:"column:is_active; type:boolean; not null; default:true;"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
