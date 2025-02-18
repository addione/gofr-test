package manager

import "gofr-test/repository"

type ManagerDi struct {
	UserManager  *UserManager
	RepositoryDI *repository.RepositoryDI
}

func NewManagerDi() *ManagerDi {
	mdi := &ManagerDi{RepositoryDI: repository.NewRepositoryDI()}
	mdi.UserManager = newUserManager(mdi)
	return mdi
}

func (mdi *ManagerDi) GetUserManager() *UserManager {
	return mdi.UserManager
}
