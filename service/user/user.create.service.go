package user

import (
	models "github.com/moluh/ginrest/model"
	repositorys "github.com/moluh/ginrest/repository/users"
	schemas "github.com/moluh/ginrest/schema"
)

type ServiceCreate interface {
	CreateUserService(input *schemas.SchemaUser) (*models.UserModel, schemas.SchemaDatabaseError)
}

type serviceCreate struct {
	repository repositorys.RepositoryCreate
}

func NewServiceCreate(repository repositorys.RepositoryCreate) *serviceCreate {
	return &serviceCreate{repository: repository}
}

func (s *serviceCreate) CreateUserService(input *schemas.SchemaUser) (*models.UserModel, schemas.SchemaDatabaseError) {

	var user schemas.SchemaUser
	user.Name = input.Name

	res, err := s.repository.CreateUserRepository(&user)
	return res, err
}
