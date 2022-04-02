package repository

import (
	"go-api/entity"

	"gorm.io/gorm"
)

type SmartPhoneRepo interface {
	FindHPs() ([]entity.SmartPhone, error)
	FindHPByID(id int) (entity.SmartPhone, error)
	FindHPByCon(condition string) (entity.SmartPhone, error)
	InsertHP(HP entity.SmartPhone) (entity.SmartPhone, error)
	UpdateHP(entity.SmartPhone) (entity.SmartPhone, error)
	Delete(HP entity.SmartPhone) (entity.SmartPhone, error)
}

type smartPhoneConn struct{
	connection *gorm.DB
}

func NewSmartPhoneRepo(db *gorm.DB) SmartPhoneRepo{
	return &smartPhoneConn{
		connection: db,
	}
}

func (db *smartPhoneConn) FindHPs() ([]entity.SmartPhone, error){
	var entHPs []entity.SmartPhone
	err := db.connection.Find(&entHPs).Error
	return entHPs, err
}

func (db *smartPhoneConn) FindHPByID(id int) (entity.SmartPhone, error){
	var entHP entity.SmartPhone
	err := db.connection.Find(&entHP, id).Error
	return entHP, err
}

func (db *smartPhoneConn) FindHPByCon(condition string) (entity.SmartPhone, error){
	var entHP entity.SmartPhone
	err := db.connection.Where(condition).Take(&entHP).Error
	return entHP, err
}

func (db *smartPhoneConn) InsertHP(HP entity.SmartPhone) (entity.SmartPhone, error){
	err := db.connection.Create(&HP).Error
	return HP, err
}

func (db *smartPhoneConn) UpdateHP(HP entity.SmartPhone) (entity.SmartPhone, error){
	err := db.connection.Save(&HP).Error
	return HP, err
}

func (db *smartPhoneConn) Delete(HP entity.SmartPhone) (entity.SmartPhone, error){
	err := db.connection.Delete(&HP).Error
	return HP, err
}