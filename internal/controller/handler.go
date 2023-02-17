package controller

import "git.01.alem.school/Sultanye/problems-database/internal/usecase"

type Handler struct {
	usecases *usecase.UseCase
}

func NewHandler(usecases *usecase.UseCase) *Handler {
	return &Handler{usecases: usecases}
}
