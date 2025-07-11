package application

import (
	"arch-linter-test/internal/adapters"
	"arch-linter-test/internal/application/domain"
)

type IRepo interface {
	GetModel() domain.Model
}

type UseCase struct {
	repo adapters.Repo // - ЛИНТЕР ПАДАЕТ
	//repo IRepo // - ЛИНТЕР НЕ ПАДАЕТ
}

func (u *UseCase) DoSomething() {
	u.repo.GetModel()
}
