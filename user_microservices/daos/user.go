package daos

import (
	"errors"
	"fmt"
	"user_microservices/databases"
	"user_microservices/helper"
	"user_microservices/middleware"
	"user_microservices/models"
)

type User struct {
	helper helper.Helper
}

func (m *User) UserCreate(params models.CreateUser) (models.UserCreate, error) {

	user := models.UserCreate{}

	user.IdUser = m.helper.StringWithCharset()
	user.Nama = params.Nama
	user.NoTelp = params.NoTelp
	user.Email = params.Email
	user.Password, _ = EncryptPassword(params.Password)
	user.CreatedAt = m.helper.GetTimeNow()
	//user.Token = m.helper.StringWithCharset()

	err := databases.DatabaseWaizly.DB.Table("users").Create(&user).Error

	if err != nil {
		return models.UserCreate{}, err
	}

	return user, nil
}

func (m *User) UserGet(params models.GetUser) ([]models.UserGet, error) {

	user := []models.UserGet{}

	err := databases.DatabaseWaizly.DB.Table("users")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}
	if params.Email != "" {
		err = err.Where("email = ?", params.Email)
	}
	if params.Search != "" {
		err = err.Where("nama ilike '%" + params.Search + "%' OR email ilike '%" + params.Search + "%'")
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
		return []models.UserGet{}, errx
	}

	return user, nil
}

func (m *User) LoginCheck(params models.UserToken) error {

	checkakun := models.UserGet{}
	var check bool

	check = databases.DatabaseWaizly.DB.Table("users").
		Where("email = ?", params.Email).Find(&checkakun).RecordNotFound()

	if check == true {
		err := errors.New("Email Tidak Ditemukan")
		return err
	}

	return error(nil)

}

func (m *User) UserGetLogin(params models.UserLogin) ([]models.UserGet, error) {

	user := []models.UserGet{}

	err := databases.DatabaseWaizly.DB.Table("user")
	if params.IdUser != "" {
		err = err.Where("id_user = ?", params.IdUser)
	}

	err = err.Find(&user)

	errx := err.Error

	if errx != nil {
		return []models.UserGet{}, errx
	}

	return user, nil
}

func (m *User) Signin(params models.UserToken) ([]models.UserGet, string, error) {

	userGet := models.GetUser{}
	userRead := []models.UserGet{}
	updateToken := models.UserUpdate{}
	var token string
	var er error

	err := m.LoginCheck(params)

	if err != nil {
		return userRead, "", err
	}

	if params.Email != "" {
		userGet.Email = params.Email
	}
	if params.Password != "" {
		userGet.Password = params.Password
	}

	userRead, err = m.UserGet(userGet)

	if err != nil {
		return userRead, "", err
	}

	//token, er := m.helper.GetToken(userRead[0].IdUser)
	password, _ := DecryptPassword(userRead[0].Password)

	if userRead[0].Email == params.Email && params.Password == password {
		fmt.Println("cocok")
		token, er = middleware.CreateAuth(userRead[0].IdUser, "user", "none", "none")

		if er != nil {
			return userRead, "", er
		}

		updateToken.Token = token

		err = databases.DatabaseWaizly.DB.Table("users").Where("id_user = ?", userRead[0].IdUser).Update(&updateToken).Error

		if err != nil {
			return userRead, "", err
		}
	} else {
		fmt.Println("tidak cocok")
	}

	return userRead, token, nil

}
