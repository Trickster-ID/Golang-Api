package dto

import "time"

type SmartPhonePostDTO struct {
	Brand       string `json:"brand" form:"brand" binding:"required"`
	Type        string `json:"type" form:"type" binding:"required"`
	Chipset     string `json:"chipset" form:"chipset" binding:"required"`
	NFC         bool   `json:"nfc" form:"nfc" binding:"required"`
	ReleaseDate time.Time `json:"releasedate" form:"releasedate"`
}

// type SmartPhonePutDTO struct {
// 	ID          uint64 `json:"id" form:"id" binding:"required"`
// 	Brand       string `json:"brand" form:"brand" binding:"required"`
// 	Type        string `json:"type" form:"type" binding:"required"`
// 	Chipset     string `json:"chipset" form:"chipset" binding:"required"`
// 	NFC         bool   `json:"nfc" form:"nfc" binding:"required"`
// 	ReleaseDate time.Time `json:"releasedate" form:"releasedate"`
// }