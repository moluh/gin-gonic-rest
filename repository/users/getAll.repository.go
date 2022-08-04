package repository

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/moluh/ginrest/model"
	"github.com/moluh/ginrest/schema"
)

type RepositoryGetAll interface {
	GetAllUsersRepository() (*[]model.UserModel, schema.SchemaDatabaseError)
}

type repositoryGetAll struct {
	db *gorm.DB
}

func NewRepositoryGetAll(db *gorm.DB) *repositoryGetAll {
	return &repositoryGetAll{db: db}
}

func (r *repositoryGetAll) GetAllUsersRepository() (*[]model.UserModel, schema.SchemaDatabaseError) {

	var users []model.UserModel
	db := r.db.Model(&users)
	errorCode := make(chan schema.SchemaDatabaseError, 1)

	// usersResult := db.Debug().First(&users)
	usersResult := db.First(&users)

	if usersResult.RowsAffected < 1 {
		errorCode <- schema.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &users, <-errorCode
	}

	return &users, <-errorCode
}
