package handler

import "gofr-test/http/manager"

type Http struct {
	Manager *manager.ManagerDi
	Handler *Handler
}

type Manager struct {
	UserManager *manager.UserManager
}

type Handler struct {
	UserHandler *UserHandler
}

func NewHttpDi() *Http {
	http := &Http{}
	http.Manager = manager.NewManagerDi()
	http.Handler = NewHandlerDI(http)
	return http
}

func NewHandlerDI(h *Http) *Handler {
	return &Handler{
		UserHandler: initUserHandler(h),
	}
}
