package repository

type RepositoryDI struct {
	UserRepository *UserRepository
}

func NewRepositoryDI() *RepositoryDI {
	return &RepositoryDI{UserRepository: newUserRepository()}
}
