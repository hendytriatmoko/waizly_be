package models

import "mime/multipart"

type CreateMerk struct {
	NamaMerk   string               `json:"nama_merk" form:"nama_merk"`
	GambarMerk multipart.FileHeader `json:"gambar_merk" form:"gambar_merk"`
	CreatedAt  string               `json:"created_at" form:"created_at"`
}
type CreateTask struct {
	IdUser    string `json:"id_user" form:"id_user"`
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
}

type TaskCreate struct {
	IdUser    string `json:"id_user" form:"id_user"`
	IdTask    string `json:"id_task" form:"id_task"`
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Progress  int8   `json:"progress" form:"progress"`
	Status    string `json:"status" form:"status"`
	Pin       bool   `json:"pin" form:"pin"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

type UpdateTask struct {
	IdTask    string `json:"id_task" form:"id_task"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Progress  int8   `json:"progress" form:"progress"`
	Status    string `json:"status" form:"status"`
	Pin       string `json:"pin" form:"pin"`
}

type TaskUpdate struct {
	UpdatedAt string `json:"updated_at" form:"updated_at"`
	IdUser    string `json:"id_user" form:"id_user"`
	IdTask    string `json:"id_task" form:"id_task"`
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Progress  int8   `json:"progress" form:"progress"`
	Status    string `json:"status" form:"status"`
	Pin       string `json:"pin" form:"pin"`
}

type GetTask struct {
	IdTask string `json:"id_task" form:"id_task"`
	IdUser string `json:"id_user" form:"id_user"`
	Pin    string `json:"pin" form:"pin"`
	Search string `json:"search" form:"search"`
	Status string `json:"status" form:"status"`
	Limit  *int64 `json:"limit" form:"limit"`
	Offset *int64 `json:"offset" form:"offset"`
}

type TaskGet struct {
	IdTask    string `json:"id_task" form:"id_task"`
	IdUser    string `json:"id_user" form:"id_user"`
	Nama      string `json:"nama" form:"nama"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Progress  string `json:"progress" form:"progress"`
	Pin       string `json:"pin" form:"pin"`
	Status    string `json:"status" form:"status"`
	CreatedAt string `json:"created_at" form:"created_at"`
}

type DeleteTask struct {
	IdTask    string `json:"id_task" form:"id_task"`
	DeletedAt string `json:"deleted_at" form:"deleted_at"`
}
