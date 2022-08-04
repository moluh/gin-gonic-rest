package user

import (
	// "github.com/moluh/ginrest/model"
	model "github.com/moluh/ginrest/model"
	repositorys "github.com/moluh/ginrest/repository/users"
	schemas "github.com/moluh/ginrest/schema"
)

type ServiceGetAll interface {
	GetAllUsersService() (*[]model.UserModel, schemas.SchemaDatabaseError)
}

type serviceGetAll struct {
	repository repositorys.RepositoryGetAll
}

func NewServiceGetAll(repository repositorys.RepositoryGetAll) *serviceGetAll {
	return &serviceGetAll{repository: repository}
}

func (s *serviceGetAll) GetAllUsersService() (*[]model.UserModel, schemas.SchemaDatabaseError) {

	res, err := s.repository.GetAllUsersRepository()
	return res, err
}
