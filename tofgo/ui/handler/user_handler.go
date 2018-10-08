package handler

type userHandler struct {
}

type UserHandler interface {
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}
