package config

import (
	"fmt"
	"go-api/entity"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB{
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("fail to load .env")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("PASS"),
		os.Getenv("DBNAME"),
		os.Getenv("PORT"))
	
	fmt.Println(dsn)
	db, errDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errDB != nil {
		panic(errDB)
	}
	db.AutoMigrate(&entity.User{}, &entity.SmartPhone{})
	return db
}

func CloseDatabaseConnection(db *gorm.DB){
	dbcon, err := db.DB()
	if err != nil {
		panic("fail to close con database")
	}
	dbcon.Close()
}