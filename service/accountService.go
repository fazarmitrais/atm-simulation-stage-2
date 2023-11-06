package service

import (
	"errors"

	"github.com/fazarmitrais/atm-simulation-stage-2/domain/account/entity"
)

func (s *Service) Import() error {
	accounts, err := s.AccountCsvRepository.Import()
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		return errors.New("no data imported")
	}
	accMap := make(map[string]*entity.Account)
	for _, a := range accounts {
		if accMap[a.AccountNumber] == nil {
			accMap[a.AccountNumber] = a
		} else {
			return errors.New("duplicate account number")
		}
	}
	err = s.AccountRepository.Store(accounts)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAll() ([]*entity.Account, error) {
	accounts, err := s.AccountRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
