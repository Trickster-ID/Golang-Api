package controller

import (
	"fmt"
	"go-api/dto"
	"go-api/helper"
	"go-api/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetProfile(c *gin.Context)
	PutProfile(c *gin.Context)
}

type userController struct{
	userService service.UserService
	jwtService service.JWTService
}

func NewUserController(newUserService service.UserService, newJwtService service.JWTService) UserController{
	return &userController{
		userService: newUserService,
		jwtService: newJwtService,
	}
}

func (uc *userController) GetProfile(c *gin.Context){
	authHeader := c.GetHeader("Authorization")
	token, errToken := uc.jwtService.ValidateToken(authHeader)
	if errToken != nil{
		panic(errToken.Error())
	}else{
		claims := token.Claims.(jwt.MapClaims)
		userid := fmt.Sprintf("%v", claims["user_id"])
		res := uc.userService.FindProfile(userid)
		response := helper.BuildResponse(true, "OK", res)
		c.JSON(http.StatusOK, response)
	}
}

func (uc *userController) PutProfile(c *gin.Context){
	var userUpdateDTO dto.UserPutDTO
	errDTO := c.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}else{
		authHeader := c.GetHeader("Authorization")
		token, errToken := uc.jwtService.ValidateToken(authHeader)
		if errToken != nil {
			panic(errToken.Error())
		}else{
			claims := token.Claims.(jwt.MapClaims)
			id, errPrs := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
			if errPrs != nil {
				panic(errPrs.Error())
			}else{
				userUpdateDTO.ID = id
				u, errServ := uc.userService.UpdateProfile(userUpdateDTO)
				if errServ != nil{
					panic(errServ.Error())
				}else{
					res := helper.BuildResponse(true, "OK!", u)
					c.JSON(http.StatusOK, res)
				}
			}
		}
	}
}