package service

import (
	"context"
	"fmt"
	"golangassignment/wallet/entity"
	"golangassignment/wallet/enum"
	postgresgorm "golangassignment/wallet/repository/postgres_gorm"
)

type ITransactionService interface {
	CreateTransaction(ctx context.Context, user *entity.Transaction) (entity.Transaction, error)
	GetTransactionByID(ctx context.Context, id int) (entity.Transaction, error)
	GetTransactionByWalletID(ctx context.Context, userid int) ([]entity.TransactionResponse, error)
	DeleteTransaction(ctx context.Context, id int) error
	UpdateWalletBalance(ctx context.Context, id int, amount float32, trxType string) error
}

// untuk menggunakan gorm
type transactionService struct {
	transactionRepo postgresgorm.ITransactionRepository
}

// untuk menggunakan gorm
func NewTransactionService(transactionRepo postgresgorm.ITransactionRepository) ITransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (r *transactionService) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error) {

	if transaction.TrxType == enum.PAYMENT || transaction.TrxType == enum.TOPUP || transaction.TrxType == enum.TRANSFER {
		var err error
		createdTransaction, err := r.transactionRepo.CreateTransaction(ctx, transaction)
		if err != nil {
			return entity.Transaction{}, fmt.Errorf("error created transaction: %v", err)
		}

		if transaction.TrxType == enum.PAYMENT || transaction.TrxType == enum.TOPUP {
			err := r.transactionRepo.UpdateWalletBalance(ctx, transaction.WalletID, transaction.Amount, transaction.TrxType)

			if err != nil {
				return entity.Transaction{}, fmt.Errorf("error update wallet balance: %v", err)
			}
		} else if transaction.TrxType == enum.TRANSFER {
			err := r.transactionRepo.UpdateWalletBalance(ctx, transaction.WalletID, transaction.Amount, enum.TRANSFERIN)
			if err != nil {
				return entity.Transaction{}, fmt.Errorf("error update wallet balance: %v", err)
			}

			err2 := r.transactionRepo.UpdateWalletBalance(ctx, transaction.WalletSourceID, transaction.Amount, enum.TRANSFEROUT)
			if err2 != nil {
				return entity.Transaction{}, fmt.Errorf("error update wallet balance: %v", err)
			}
		}

		return createdTransaction, nil
	} else {
		return entity.Transaction{}, fmt.Errorf("error unsupport transaction type")
	}
}

func (r *transactionService) GetTransactionByID(ctx context.Context, id int) (entity.Transaction, error) {
	transaction, err := r.transactionRepo.GetTransactionByID(ctx, id)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("error transaction not found: %v", err)
	}

	return transaction, nil
}

func (r *transactionService) GetTransactionByWalletID(ctx context.Context, userid int) ([]entity.TransactionResponse, error) {
	transaction, err := r.transactionRepo.GetTransactionByWalletID(ctx, userid)
	if err != nil {
		return []entity.TransactionResponse{}, fmt.Errorf("error transaction not found: %v", err)
	}

	return transaction, nil
}

func (r *transactionService) DeleteTransaction(ctx context.Context, id int) error {
	err := r.transactionRepo.DeleteTransaction(ctx, id)
	if err != nil {
		return fmt.Errorf("error transaction not found: %v", err)
	}

	return nil
}

func (r *transactionService) UpdateWalletBalance(ctx context.Context, id int, amount float32, trxType string) error {

	err := r.transactionRepo.UpdateWalletBalance(ctx, id, amount, trxType)
	if err != nil {
		return fmt.Errorf("error transaction not found: %v", err)
	}

	return nil
}
