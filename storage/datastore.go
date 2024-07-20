package storage

import "fairmoneytest/model"

//go:generate mockgen -source=datastore.go -destination=../mocks/datastore_mock.go -package=mocks
type DataStore interface {
	Transaction
}

type Transaction interface {
	RecordTransaction(data *model.TransactionRecords) error
	UpdateAccountBalance(ID string, data *model.TransactionRecords) error
	GetTransactionByID(ID string) (*model.TransactionRecords, error)
	GetTransactionByReference(ref string) (*model.TransactionRecords, error)
}
