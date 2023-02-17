package usecase

import "git.01.alem.school/Sultanye/problems-database/internal/repository"

type UseCase struct {
	ToDoProblem
}

func NewUseCase(repos *repository.Repository) *UseCase {
	return &UseCase{
		ToDoProblem: NewTodoProblemUseCase(repos),
	}
}
