package repo

import (
	"go-services/global"
	"go-services/internal/po"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{db: global.PDB}
}

func (ur *UserRepo) GetInfoUser() string {

	return "Vu Duc Tien"
}

func (ur *UserRepo) Create(user *po.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepo) GetByID(id uint) (*po.User, error) {
	var user po.User
	err := ur.db.First(&user, id).Error
	return &user, err
}

func (ur *UserRepo) Update(user *po.User) error {
	return ur.db.Save(user).Error
}

func (ur *UserRepo) Delete(id uint) error {
	return ur.db.Delete(&po.User{}, id).Error
}
