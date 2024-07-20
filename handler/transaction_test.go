package handler

import (
	"errors"
	"fairmoneytest/mocks"
	"fairmoneytest/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestTransactionHandler_RecordCreditTransaction_EdgeCases(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mocks.NewMockDataStore(ctrl)

	// Test invalid amount
	t.Run("InvalidAmount", func(t *testing.T) {
		handler := NewTransactionHandler(mockStore)
		_, err := handler.RecordCreditTransaction(&model.Transactions{
			AccountID: "test-1234",
			Reference: "ref-01",
			Amount:    -100,
		})
		assert.Error(t, err)
		assert.Equal(t, "invalid amount:-100", err.Error())
	})

	// Test duplicate transaction
	t.Run("DuplicateTransaction", func(t *testing.T) {
		handler := NewTransactionHandler(mockStore)
		mockStore.EXPECT().GetTransactionByReference("ref-01").Return(&model.TransactionRecords{}, nil).Times(1)
		_, err := handler.RecordCreditTransaction(&model.Transactions{
			AccountID: "test-1234",
			Reference: "ref-01",
			Amount:    100,
		})
		assert.Error(t, err)
		assert.Equal(t, "duplicate transaction reference:ref-01", err.Error())
	})

	// Test database error on RecordTransaction
	t.Run("RecordTransactionError", func(t *testing.T) {
		handler := NewTransactionHandler(mockStore)
		mockStore.EXPECT().GetTransactionByReference("ref-01").Return(nil, mongo.ErrNoDocuments).Times(1)
		mockStore.EXPECT().RecordTransaction(gomock.Any()).Return(errors.New("db error")).Times(1)
		_, err := handler.RecordCreditTransaction(&model.Transactions{
			AccountID: "test-1234",
			Reference: "ref-01",
			Amount:    100,
		})
		assert.Error(t, err)
		assert.Equal(t, "error recording the transaction:db error", err.Error())
	})
}

func TestTransactionHandler_RecordDebitTransaction_EdgeCases(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctrl.Finish()

	mockStore := mocks.NewMockDataStore(ctrl)

	//Handle case invalid amount
	t.Run("Invalid amount", func(t *testing.T) {
		handler := NewTransactionHandler(mockStore)
		_, err := handler.RecordDebitTransaction(&model.Transactions{
			AccountID: "test-1234",
			Reference: "ref-01",
			Amount:    -1000,
		})
		assert.Error(t, err)
		assert.Equal(t, "error invalid amount:-1000", err.Error())
	})

	//Handle check for duplicate transaction
	t.Run("Check for duplicate transaction", func(t *testing.T) {
		handler := NewTransactionHandler(mockStore)
		mockStore.EXPECT().GetTransactionByReference(gomock.Any()).Return(&model.TransactionRecords{
			Balance: 50000,
		}, nil).Times(1)
		_, err := handler.RecordDebitTransaction(&model.Transactions{
			AccountID: "test-1234",
			Reference: "ref-01",
			Amount:    50000,
		})
		assert.Error(t, err)
		assert.Equal(t, "duplicate transaction:ref-01", err.Error())
	})

	//Insufficient account balance
	t.Run("Insufficient account balance", func(t *testing.T) {
		handler := NewTransactionHandler(mockStore)
		mockStore.EXPECT().GetTransactionByReference(gomock.Any()).Return(&model.TransactionRecords{}, nil).Times(1)
		mockStore.EXPECT().GetTransactionByID("test-1234").Return(&model.TransactionRecords{
			Balance: 100,
		}, nil).Times(1)
		_, err := handler.RecordDebitTransaction(&model.Transactions{
			AccountID: "test-1234",
			Reference: "ref-01",
			Amount:    200,
		})
		assert.Error(t, err)
		assert.Equal(t, "duplicate transaction:ref-01", err.Error())
	})

	//Test database error on RecordTransaction
	t.Run("RecordTransaction error", func(t *testing.T) {
		handler := NewTransactionHandler(mockStore)
		mockStore.EXPECT().GetTransactionByReference("ref-01").Return(&model.TransactionRecords{Balance: 50000}, nil).Times(1)
		mockStore.EXPECT().RecordTransaction(gomock.Any()).Return(nil).Times(1)
		_, err := handler.RecordDebitTransaction(&model.Transactions{
			AccountID: "test-1234",
			Reference: "ref-01",
			Amount:    -200,
		})
		assert.Error(t, err)
		assert.Equal(t, "error invalid amount:-200", err.Error())
	})
}
