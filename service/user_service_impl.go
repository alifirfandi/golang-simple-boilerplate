package service

import (
	"golang-simple-boilerplate/model"
	"golang-simple-boilerplate/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(UserRepository *repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: *UserRepository,
	}
}

func (Service UserServiceImpl) Profile(Request model.ProfileRequest) (Response model.ProfileResponse, Error error) {
	Response, Error = Service.UserRepository.Profile(Request)
	return Response, Error
}
