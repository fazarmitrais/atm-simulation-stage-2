package mapdatastore

import "github.com/fazarmitrais/atm-simulation-stage-2/domain/account/entity"

type MapDatastore struct {
	Account map[string]*entity.Account
}

func NewMapDatastore() *MapDatastore {
	m := &MapDatastore{}
	m.Account = make(map[string]*entity.Account)
	return m
}

func (m *MapDatastore) Store(accounts []*entity.Account) error {
	for _, account := range accounts {
		m.Account[account.AccountNumber] = account
	}
	return nil
}

func (m *MapDatastore) GetByAccountNumber(accountNumber string) (*entity.Account, error) {
	return m.Account[accountNumber], nil
}

func (m *MapDatastore) GetAll() ([]*entity.Account, error) {
	var accounts []*entity.Account
	for _, account := range m.Account {
		accounts = append(accounts, account)
	}
	return accounts, nil
}
