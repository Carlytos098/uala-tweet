package factory

import (
	"uala-tweet/internal/repository"
	"uala-tweet/internal/repository/interfaces"
)

type RepositoryFactory struct{}

func NewFactory() *RepositoryFactory {
	return &RepositoryFactory{}
}

func (rf *RepositoryFactory) NewMemoryRepository() interfaces.UalaRepository {
	return repository.NewMemoryRepo()
}
