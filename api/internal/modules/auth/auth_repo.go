package auth

import "db"

type AuthRepo interface {
	Create()
	GetByUsername()
}

var _ AuthRepo = (*AuthRepoImpl)(nil)

type AuthRepoImpl struct {
	db db.Database
}

func (r *AuthRepoImpl) Create() {
	println("Create")
}

func (r *AuthRepoImpl) GetByUsername() {
	println("GetByUsername")
}
