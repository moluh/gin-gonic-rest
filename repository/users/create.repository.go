package repository

import (
	"net/http"

	"gorm.io/gorm"

	"github.com/moluh/ginrest/model"
	"github.com/moluh/ginrest/schema"
)

type RepositoryCreate interface {
	CreateUserRepository(input *schema.SchemaUser) (*model.UserModel, schema.SchemaDatabaseError)
}

type repositoryCreate struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repositoryCreate {
	return &repositoryCreate{db: db}
}

func (r *repositoryCreate) CreateUserRepository(input *schema.SchemaUser) (*model.UserModel, schema.SchemaDatabaseError) {

	var users model.UserModel
	db := r.db.Model(&users)
	errorCode := make(chan schema.SchemaDatabaseError, 1)

	checkUserExist := db.Debug().First(&users, "name = ?", input.Name)

	if checkUserExist.RowsAffected > 0 {
		errorCode <- schema.SchemaDatabaseError{
			Code: http.StatusConflict,
			Type: "error_01",
		}
		return &users, <-errorCode
	}

	users.Name = input.Name

	addNewUser := db.Debug().Create(&users).Commit()

	if addNewUser.RowsAffected < 1 {
		errorCode <- schema.SchemaDatabaseError{
			Code: http.StatusForbidden,
			Type: "error_02",
		}
		return &users, <-errorCode
	}

	return &users, <-errorCode
}
