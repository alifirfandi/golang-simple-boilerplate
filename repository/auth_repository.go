package repository

import "golang-simple-boilerplate/model"

type AuthRepository interface {
	Login(Request model.AuthRequest) (Response model.AuthResponse, UserExists bool, Error error)
}
