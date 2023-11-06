package repository

import "github.com/fazarmitrais/atm-simulation-stage-2/domain/account/entity"

type AccountRepository interface {
	Store(accounts []*entity.Account) error
	GetByAccountNumber(accountNumber string) (*entity.Account, error)
	GetAll() ([]*entity.Account, error)
}
