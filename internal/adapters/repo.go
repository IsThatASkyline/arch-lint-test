package adapters

import "arch-linter-test/internal/application/domain"

type Repo struct{}

func (r Repo) GetModel() domain.Model {
	return domain.Model{}
}
