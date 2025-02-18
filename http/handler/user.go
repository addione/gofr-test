package handler

import (
	"gofr-test/http/manager"

	"gofr.dev/pkg/gofr"
)

type UserHandler struct {
	userManager *manager.UserManager
}

func initUserHandler(http *Http) *UserHandler {
	uh := &UserHandler{}
	uh.userManager = http.Manager.GetUserManager()
	return uh
}

func (uh *UserHandler) GetUserHandler(ctx *gofr.Context) (any, error) {

	return uh.userManager.GetUser(ctx)
}

func (uh *UserHandler) CreateUserHandler(ctx *gofr.Context) (any, error) {
	return "Hello world!!", nil
}

func (uh *UserHandler) UpdateUserHandler(ctx *gofr.Context) (any, error) {
	return "Hello world!!", nil
}

func (uh *UserHandler) DeleteUserHandler(ctx *gofr.Context) (any, error) {
	return "Hello world!!", nil
}
