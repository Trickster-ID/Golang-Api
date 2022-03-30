package controller

import (
	"go-api/dto"
	"go-api/entity"
	"go-api/helper"
	"go-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(cx *gin.Context)
	Register(cx *gin.Context)
}

type authController struct{
	auth service.AuthService
	jwt service.JWTService
}

func NewAuthController(newauth service.AuthService, newjwt service.JWTService) AuthController{
	return &authController{
		auth: newauth,
		jwt: newjwt,
	}
}

func (c *authController) Login(cx *gin.Context){
	var loginDTO dto.LoginPostDTO
	SBerr := cx.ShouldBind(&loginDTO)
	if SBerr != nil{
		response := helper.BuildErrorResponse("failed to process request", SBerr.Error(), helper.EmptyObject{})
		cx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} 
	authresult := c.auth.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authresult.(entity.User); ok{
		v.Token = c.jwt.GenerateToken(strconv.FormatUint(v.ID, 10))
		cx.JSON(http.StatusOK, helper.BuildResponse(true, "OK!", v))
		return
	}
	respons := helper.BuildErrorResponse("Failed! check your credential", "invalid credential", helper.EmptyObject{})
	cx.AbortWithStatusJSON(http.StatusUnauthorized, respons)
}
func (c *authController) Register(cx *gin.Context){
	var registerDTO dto.RegisterPostDTO
	SBerr := cx.ShouldBind(&registerDTO)
	if SBerr != nil{
		response := helper.BuildErrorResponse("failed to process request", SBerr.Error(), helper.EmptyObject{})
		cx.JSON(http.StatusConflict, response)
	}
	if !(c.auth.IsDuplicateEmail(registerDTO.Email)){
		respons := helper.BuildErrorResponse("failed to register", "email is already exist", helper.EmptyObject{})
		cx.AbortWithStatusJSON(http.StatusUnauthorized, respons)
	}else{
		createdUser := c.auth.CreateUser(registerDTO)
		token := c.jwt.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helper.BuildResponse(true, "oke", createdUser)
		cx.JSON(http.StatusOK, response)
	}
}