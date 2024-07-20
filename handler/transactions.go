package handler

import (
	"fairmoneytest/model"
	"fairmoneytest/storage"
	"fmt"
)

type TransactionHandler struct {
	store storage.DataStore
}

func NewTransactionHandler(store storage.DataStore) TransactionHandler {
	return TransactionHandler{store: store}
}

func (trx *TransactionHandler) RecordCreditTransaction(data *model.Transactions) (*model.Transactions, error) {
	//Handle invalid amount
	if data.Amount <= 0 {
		return nil, fmt.Errorf("invalid amount:%v", data.Amount)
	}

	//Check for duplicate transaction
	_, err := trx.store.GetTransactionByReference(data.Reference)
	if err == nil {
		return nil, fmt.Errorf("duplicate transaction reference:%v", data.Reference)
	}

	//Record the transaction
	err = trx.store.RecordTransaction(&model.TransactionRecords{
		ID:        "test-1234", // For simplicity
		AccountID: data.AccountID,
		Reference: data.Reference,
		Amount:    data.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("error recording the transaction:%v", err)
	}

	//Get previous account balance and sum it with the new amount
	rec, err := trx.store.GetTransactionByID("test-1234")
	if err != nil {
		return nil, err
	}

	//Update the account balance
	rec.Balance += data.Amount

	err = trx.store.UpdateAccountBalance("test-1234", &model.TransactionRecords{
		Balance: rec.Balance,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to update account:%v", err)
	}

	return &model.Transactions{
		AccountID: data.AccountID,
		Reference: data.Reference,
		Amount:    data.Amount,
	}, nil
}

func (trx *TransactionHandler) RecordDebitTransaction(data *model.Transactions) (*model.Transactions, error) {
	//Handle the case invalid amount
	if data.Amount <= 0 {
		return nil, fmt.Errorf("error invalid amount:%v", data.Amount)
	}

	//Check for duplicate transaction
	_, err := trx.store.GetTransactionByReference(data.Reference)
	if err == nil {
		return nil, fmt.Errorf("duplicate transaction:%v", data.Reference)
	}

	//Insufficient account balance
	rec, _ := trx.store.GetTransactionByID(data.AccountID)
	if rec.Balance < data.Amount {
		return nil, fmt.Errorf("insuficient account balance:%v", rec.Balance)
	}

	//Record the transactions
	err = trx.store.RecordTransaction(&model.TransactionRecords{
		ID:        "test-1234", // For simplicity
		AccountID: data.AccountID,
		Reference: data.Reference,
		Amount:    data.Amount,
	})
	if err != nil {
		return nil, fmt.Errorf("error recording the transaction:%v", err)
	}
	//Update the account
	rec.Balance -= data.Amount

	err = trx.store.UpdateAccountBalance("test-1234", &model.TransactionRecords{
		Balance: rec.Balance,
	})

	if err != nil {
		return nil, fmt.Errorf("error updating the account:%v", err)
	}
	return &model.Transactions{
		AccountID: data.AccountID,
		Reference: data.Reference,
		Amount:    data.Amount,
	}, nil
}

/*func (trx *TransactionHandler) GetTransaction(ref string) (*model.Transactions, error) {
	rec, err := trx.store.GetTransactionByReference(ref)
	if err != nil {
		return nil, err
	}
	return &model.Transactions{
		AccountID: rec.AccountID,
		Reference: rec.Reference,
		Amount:    rec.Amount,
	}, nil
}

*/
