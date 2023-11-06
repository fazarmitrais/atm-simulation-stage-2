package service

import (
	"github.com/fazarmitrais/atm-simulation-stage-2/domain/account/entity"
	"github.com/fazarmitrais/atm-simulation-stage-2/domain/account/repository"
)

type Service struct {
	AccountRepository    repository.AccountRepository
	AccountCsvRepository repository.AccountCsvRepository
}

func NewService(accountRepository repository.AccountRepository, accountCsvRepository repository.AccountCsvRepository) *Service {
	return &Service{accountRepository, accountCsvRepository}
}

type ServiceInterface interface {
	Import() error
	GetAll() ([]*entity.Account, error)
}
