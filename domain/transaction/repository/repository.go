package repository

import (
	"context"

	"github.com/fazarmitrais/atm-simulation-stage-2/domain/transaction/entity"
)

type TransactionRepository interface {
	Add(ctx context.Context, transaction *entity.Transaction) error
	GetLastTransaction(ctx context.Context, accountNumber string, numberOfTransaction int) ([]*entity.Transaction, error)
}
