package entity

import (
	"time"
)

type SmartPhone struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Brand       string `gorm:"type:varchar(255)" json:"brand"`
	Type        string `gorm:"type:varchar(255)" json:"type"`
	Chipset     string `gorm:"type:varchar(255)" json:"chipset"`
	NFC         bool `gorm:"type:bool" json:"nfc"`
	ReleaseDate time.Time
	CreateAt	time.Time
	UpdateAt	time.Time
}