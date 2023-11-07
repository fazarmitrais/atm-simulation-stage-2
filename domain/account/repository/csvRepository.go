package repository

import "github.com/fazarmitrais/atm-simulation-stage-2/domain/account/entity"

type AccountCsvRepository interface {
	Import() ([]*entity.Account, error)
}
