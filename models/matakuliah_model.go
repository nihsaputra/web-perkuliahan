package models

import (
	"log"
	"web-perkuliahan/config"
	"web-perkuliahan/entity"
)

func GetAll() []entity.Matakuliah {
	var matakuliahs []entity.Matakuliah
	err := config.DB.Find(&matakuliahs).Error
	if err != nil {
		log.Println(err)
	}

	return matakuliahs
}
func GetById(id string) entity.Matakuliah {
	var matakuliahs entity.Matakuliah
	err := config.DB.Find(&matakuliahs, "kode_mk=?", id).Error
	if err != nil {
		log.Println(err)
	}

	return matakuliahs
}
func Create(matakuliah entity.Matakuliah) bool {
	errCreate := config.DB.Create(&matakuliah)
	if errCreate.Error != nil {
		log.Println(errCreate)
		panic(errCreate)
		return false
	}

	return true
}
func Update(matakuliah entity.Matakuliah) bool {
	errCreate := config.DB.Where("kode_mk=?", matakuliah.KodeMk).Save(&matakuliah)
	if errCreate.Error != nil {
		log.Println(errCreate)
		panic(errCreate)
		return false
	}
	return true
}
func Delete(id string) error {
	byId := GetById(id)
	errDelete := config.DB.Where("kode_mk", id).Delete(&byId)
	return errDelete.Error
}
