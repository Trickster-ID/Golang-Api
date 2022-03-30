package middleware

import (
	"go-api/helper"
	"go-api/service"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthJWT(jwtService service.JWTService) gin.HandlerFunc{
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			response := helper.BuildErrorResponse("failed to process requeest", "no token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claim[user_id] : ", claims["user_id"])
			log.Println("claim[issuer] : ", claims["issuer"])
		}else{
			log.Println(err)
			response := helper.BuildErrorResponse("token is not valid", err.Error(),nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}