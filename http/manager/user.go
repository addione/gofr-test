package manager

import (
	"gofr-test/repository"

	"gofr.dev/pkg/gofr"
)

type UserManager struct {
	ur *repository.UserRepository
}

func newUserManager(mdi *ManagerDi) *UserManager {

	return &UserManager{
		ur: mdi.RepositoryDI.UserRepository,
	}
}

func (um *UserManager) GetUser(ctx *gofr.Context) (any, error) {
	return um.ur.GetUsers(ctx)
}
