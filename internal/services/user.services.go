package services

import (
	"go-services/internal/po"
	"go-services/internal/repo"
)

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetInfoUser() string {
	return us.UserRepo.GetInfoUser()
}

func (us *UserService) CreateUser(user *po.User) error {
	return us.UserRepo.Create(user)
}

func (us *UserService) GetUserByID(id uint) (*po.User, error) {
	return us.UserRepo.GetByID(id)
}

func (us *UserService) UpdateUser(user *po.User) error {
	return us.UserRepo.Update(user)
}

func (us *UserService) DeleteUser(id uint) error {
	return us.UserRepo.Delete(id)
}
