package service

import (
	"go-api/dto"
	"go-api/entity"
	"go-api/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	UpdateProfile(userDto dto.UserPutDTO) (entity.User, error) 
	FindProfile(UserID string) entity.User
}

type userService struct{UserRepo repository.UserRepo}

func NewUserService(newUserRepo repository.UserRepo) UserService{
	return &userService{
		UserRepo: newUserRepo,
	}
}

func (us *userService) UpdateProfile(userDto dto.UserPutDTO) (entity.User, error) {
	entUser := entity.User{}
	err := smapping.FillStruct(&entUser, smapping.MapFields(&userDto))
	if err != nil{
		return entUser, err
	}else{
		res := us.UserRepo.UpdateUser(entUser)
		return res, err
	}
}

func (us *userService) FindProfile(UserID string) entity.User {
	res := us.UserRepo.ProfileUser(UserID)
	return res
}