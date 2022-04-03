package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretkey(),
		issuer: "pikapi",
	}
}

func getSecretkey() string{
	sk := os.Getenv("JWT_SECRET")
	if sk != "" {
		sk = "pikapi"
	}
	return sk
}

func (js *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1,0,0).Unix(),
			Issuer: js.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(js.secretKey))
	if err != nil{
		panic(err)
	}
	return t
}

func (js *jwtService) ValidateToken(token string) (*jwt.Token, error){
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method %v", t_.Header["alg"])
		}
		return []byte(js.secretKey), nil
	})
}