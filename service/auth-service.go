package service

import (
	"go-api/dto"
	"go-api/entity"
	"go-api/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterPostDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct{userRepo repository.UserRepo}

func NewAuthService(userRep repository.UserRepo) AuthService{
	return &authService{
		userRepo: userRep,
	}
}

func (serv *authService) VerifyCredential(email string, password string) interface{}{
	res := serv.userRepo.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPass := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPass{
			return res
		}
		return false
	}
	return false
}

func (serv *authService) CreateUser(user dto.RegisterPostDTO) entity.User{
	var userToFill = entity.User{}
	err := smapping.FillStruct(&userToFill,smapping.MapFields(&user))
	if err != nil{
		log.Fatalf("failed to create user : %v", err)
	}
	return serv.userRepo.InsertUser(userToFill)
}

func (serv *authService) FindByEmail(email string) entity.User{
	return serv.userRepo.FindByEmail(email)
}

func (serv *authService) IsDuplicateEmail(email string) bool{
	res := serv.userRepo.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPass []byte) bool{
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPass)
	if err != nil{
		log.Println(err)
		return false
	}
	return true
}