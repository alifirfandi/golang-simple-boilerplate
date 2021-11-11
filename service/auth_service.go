package service

import "golang-simple-boilerplate/model"

type AuthService interface {
	Login(Request model.AuthRequest) (Response model.AuthResponse, UserExists bool, Error error)
}
