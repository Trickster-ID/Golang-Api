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

type SmartPhoneController interface {
	SpGetAll(g *gin.Context)
	SpGetByCond(g *gin.Context)
	SpPost(g *gin.Context)
	SpPut(g *gin.Context)
	SpDelete(g *gin.Context)
}

type smartPhoneController struct{spService service.SmartPhoneService}

func NewSmartPhoneController(newSpService service.SmartPhoneService) SmartPhoneController{
	return &smartPhoneController{newSpService}
}

func (spc *smartPhoneController) SpGetAll(c *gin.Context){
	res, err := spc.spService.FindHPs()
	if err != nil {
		response := helper.BuildErrorResponse("Failed to find all smartphone!", err.Error(), helper.EmptyObject{})
		c.JSON(http.StatusBadRequest, response)
	}else{
		response := helper.BuildResponse(true, "OK", res)
		c.JSON(http.StatusOK, response)
	}
}

func (spc *smartPhoneController) SpGetByCond(c *gin.Context){
	res, err := spc.spService.FindHPByCon(c.Param("condition"))
	resultHandling("Failed to find smartphone with condition!", res, err, c)
}

func (spc *smartPhoneController) SpPost(c *gin.Context){
	var spDTO dto.SmartPhonePostDTO
	errBind := c.ShouldBind(&spDTO)
	if errBind != nil {
		errorHandling("error when binding json raw param", errBind, c)
	}else{
		res, err := spc.spService.InsertHP(spDTO)
		resultHandling("Failed to find smartphone with condition!", res, err, c)
	}
}

func (spc *smartPhoneController) SpPut(c *gin.Context){
	var spDTO dto.SmartPhonePostDTO
	id,errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		errorHandling("error when converting ID to int", errConv, c)
	}else{
		errBind := c.ShouldBind(&spDTO)
		if errBind != nil {
			errorHandling("error when binding json raw param", errBind, c)
		}else{
			res, err := spc.spService.UpdateHP(id, spDTO)
			resultHandling("Failed to find smartphone with condition!", res, err, c)
		}
	}
}

func (spc *smartPhoneController) SpDelete(c *gin.Context){
	id,errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		errorHandling("error when converting ID to int", errConv, c)
	}else{
		res, err := spc.spService.Delete(id)
		resultHandling("Failed to delete data!", res, err, c)
	}
}

func resultHandling(errMessage string, result entity.SmartPhone, err error,c *gin.Context) {
	if err != nil {
		response := helper.BuildErrorResponse(errMessage, err.Error(), helper.EmptyObject{})
		c.JSON(http.StatusBadRequest, response)
	}else{
		response := helper.BuildResponse(true, "OK", result)
		c.JSON(http.StatusOK, response)
	}
}

func errorHandling(errMessage string, err error, c *gin.Context){
	response := helper.BuildErrorResponse(errMessage, err.Error(), helper.EmptyObject{})
	c.JSON(http.StatusBadRequest, response)
}
