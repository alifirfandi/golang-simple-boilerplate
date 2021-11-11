package service

import "golang-simple-boilerplate/model"

type UserService interface {
	Profile(Request model.ProfileRequest) (Response model.ProfileResponse, Error error)
}
