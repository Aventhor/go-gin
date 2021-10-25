package users

import (
	"rest/database"
)

func FindAll(condition interface{}) ([]UserModel, error) {
	db := database.GetDB()
	var models []UserModel
	err := db.Find(&models).Error
	return models, err
}

func FindOne(id string) (UserModel, error) {
	db := database.GetDB()
	var model UserModel
	err := db.First(&model, id).Error
	return model, err
}

func CreateOne(data interface{}) error {
	db := database.GetDB()
	var model UserModel
	err := db.Model(&model).Save(data).Error
	return err
}

func (model *UserModel) Update(data interface{}) error {
	db := database.GetDB()
	err := db.Model(model).Updates(data).Error
	return err
}

func (model *UserModel) Delete() error {
	db := database.GetDB()
	err := db.Delete(model).Error
	return err
}
