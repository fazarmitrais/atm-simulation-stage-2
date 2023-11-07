package service

import (
	"context"

	"github.com/fazarmitrais/atm-simulation-stage-2/domain/account/entity"
	"github.com/fazarmitrais/atm-simulation-stage-2/domain/account/repository"
	"github.com/fazarmitrais/atm-simulation-stage-2/lib/responseFormatter"
)

type Service struct {
	AccountRepository    repository.AccountRepository
	AccountCsvRepository repository.AccountCsvRepository
}

func NewService(accountRepository repository.AccountRepository, accountCsvRepository repository.AccountCsvRepository) *Service {
	return &Service{accountRepository, accountCsvRepository}
}

type ServiceInterface interface {
	PINValidation(c context.Context, account entity.Account) *responseFormatter.ResponseFormatter
	Transfer(ctx, transfer entity.Transfer) (*entity.Account, *responseFormatter.ResponseFormatter)
	BalanceCheck(ctx context.Context, acctNbr string) (*entity.Account, *responseFormatter.ResponseFormatter)
	Import() error
	GetAll() ([]*entity.Account, error)
}
