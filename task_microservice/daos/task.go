package daos

import (
	"task_microservices/databases"
	"task_microservices/helper"
	"task_microservices/models"
)

type Merk struct {
	helper helper.Helper
}

func (m *Merk) TaskCreate(params models.CreateTask) (models.TaskCreate, error) {

	inserttask := models.TaskCreate{}

	inserttask.Nama = params.Nama
	inserttask.Deskripsi = params.Deskripsi
	inserttask.CreatedAt = m.helper.GetTimeNow()
	inserttask.IdUser = params.IdUser
	inserttask.IdTask = m.helper.StringWithCharset()
	inserttask.Progress = 0
	inserttask.Status = "progress"
	inserttask.Pin = false

	err := databases.DatabaseWaizly.DB.Table("tasks").Create(&inserttask).Error

	if err != nil {
		return models.TaskCreate{}, err
	}

	return inserttask, nil
}

func (m *Merk) TaskUpdate(params models.UpdateTask) (models.TaskUpdate, error) {

	newtask := models.TaskUpdate{}

	newtask.UpdatedAt = m.helper.GetTimeNow()

	newtask.Nama = params.Nama
	newtask.Deskripsi = params.Deskripsi
	newtask.Status = params.Status
	newtask.Pin = params.Pin
	newtask.Progress = params.Progress

	err := databases.DatabaseWaizly.DB.Table("tasks").Where("id_task = ?", params.IdTask).Update(&newtask).Error

	if err != nil {
		return models.TaskUpdate{}, err
	}

	return newtask, nil

}

func (m *Merk) TaskGet(params models.GetTask) ([]models.TaskGet, error) {

	user := []models.TaskGet{}

	err := databases.DatabaseWaizly.DB.Table("tasks")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}
	if params.IdTask != "" {
		err = err.Where("id_task = ?", params.IdTask)
	}
	if params.Status != "" {
		err = err.Where("status = ?", params.Status)
	}
	if params.Pin != "" {
		err = err.Where("pin = ?", params.Pin)
	}
	if params.Search != "" {
		err = err.Where("nama ilike '%" + params.Search + "%' OR deskripsi ilike '%" + params.Search + "%'")
	}
	if params.Limit != nil {
		err = err.Limit(*params.Limit)
	}
	if params.Offset != nil {
		err = err.Offset(*params.Offset)
	}

	err = err.Find(&user)

	errx := err.Error

	if errx != nil {
		return []models.TaskGet{}, errx
	}

	return user, nil
}

//func (m *Merk) MerkDelete(params models.DeleteMerk) (models.DeleteMerk, error) {
//
//	merk := models.DeleteMerk{}
//
//	merk.DeletedAt = m.helper.GetTimeNow()
//
//	err := databases.DatabaseSellPump.DB.Table("merk").Where("id_merk = ?", params.IdMerk).Update(&merk).Error
//
//	if err != nil {
//		return models.DeleteMerk{}, err
//	}
//
//	return merk, nil
//
//}