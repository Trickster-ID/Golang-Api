package service

import (
	"go-api/dto"
	"go-api/entity"
	"go-api/repository"

	"github.com/mashingan/smapping"
)

type SmartPhoneService interface {
	FindHPs() ([]entity.SmartPhone, error)
	FindHPByCon(condition string) (entity.SmartPhone, error)
	InsertHP(HPdto dto.SmartPhonePostDTO) (entity.SmartPhone, error)
	UpdateHP(id int, HPdto dto.SmartPhonePostDTO) (entity.SmartPhone, error)
	Delete(id int) (entity.SmartPhone, error)
}

type smartPhoneService struct{HpRepo repository.SmartPhoneRepo}

func NewSmartPhoneService(newHpRepo repository.SmartPhoneRepo) SmartPhoneService{
	return &smartPhoneService{
		HpRepo: newHpRepo,
	}
}

func (sps *smartPhoneService) FindHPs() ([]entity.SmartPhone, error){
	return sps.HpRepo.FindHPs()
}

func (sps *smartPhoneService) FindHPByCon(condition string) (entity.SmartPhone, error){
	return sps.HpRepo.FindHPByCon(condition)
}

func (sps *smartPhoneService) InsertHP(HPdto dto.SmartPhonePostDTO) (entity.SmartPhone, error){
	var hpToFill = entity.SmartPhone{}
	errSM := smapping.FillStruct(&hpToFill,smapping.MapFields(&HPdto))
	if errSM != nil{
		return hpToFill, errSM
	}else{
		return sps.HpRepo.InsertHP(hpToFill)
	}
}

func (sps *smartPhoneService) UpdateHP(id int, HPdto dto.SmartPhonePostDTO) (entity.SmartPhone, error){
	sp, err1 := sps.HpRepo.FindHPByID(id)
	if err1 != nil{
		return sp, err1
	}else{
		sp.Brand = HPdto.Brand
		sp.Type = HPdto.Type
		sp.Chipset = HPdto.Chipset
		sp.NFC = HPdto.NFC
		sp.ReleaseDate = HPdto.ReleaseDate
		return sps.HpRepo.UpdateHP(sp)
	}
}

func (sps *smartPhoneService) Delete(id int) (entity.SmartPhone, error){
	res1, err1 := sps.HpRepo.FindHPByID(id)
	if err1 != nil{
		return res1, err1
	}else{
		return sps.HpRepo.Delete(res1)
	}
}