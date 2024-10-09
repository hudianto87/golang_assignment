package postgresgorm

import (
	"context"
	"errors"
	"fmt"
	"golangassignment/wallet/entity"
	"golangassignment/wallet/enum"
	"log"
	"time"

	"gorm.io/gorm"
)

type GormDBInterface interface {
	WithContext(ctx context.Context) *gorm.DB
}

type ITransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error)
	GetTransactionByID(ctx context.Context, id int) (entity.Transaction, error)
	GetTransactionByWalletID(ctx context.Context, walletid int) ([]entity.TransactionResponse, error)
	DeleteTransaction(ctx context.Context, id int) error
	UpdateWalletBalance(ctx context.Context, id int, amount float32, trxType string) error
}

type transactionRepository struct {
	db GormDBIface
}

func NewTransactionRepository(db GormDBInterface) ITransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error) {

	if err := r.db.WithContext(ctx).Create(transaction).Error; err != nil {
		log.Printf("Error creating transaction : %v\n", err)
	}

	return *transaction, nil
}

func (r *transactionRepository) GetTransactionByID(ctx context.Context, id int) (entity.Transaction, error) {
	var transaction entity.Transaction
	if err := r.db.WithContext(ctx).Select("id", "walletid", "trxtype", "walletsourceid", "description", "amount", "created_at", "updated_at").First(&transaction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Transaction{}, nil
		}
		log.Printf("Error get wallet : %v\n", err)
		return entity.Transaction{}, nil
	}

	return transaction, nil
}

func (r *transactionRepository) GetTransactionByWalletID(ctx context.Context, walletid int) ([]entity.TransactionResponse, error) {
	var transaction []entity.TransactionResponse
	if err := r.db.WithContext(ctx).
		Table("transactions").
		Select("transactions.id, transactions.walletid, wallets.name as walletname, transactions.trxtype, transactions.walletsourceid, transactions.amount, users.name as username, transactions.created_at, transactions.updated_at").
		Joins("JOIN wallets on wallets.id = transactions.walletid").
		Joins("JOIN users on users.id = CAST(wallets.user_id AS INTEGER)").
		Where("transactions.walletid = ?", walletid).
		Find(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entity.TransactionResponse{}, nil
		}
		log.Printf("Error retrieve wallet transactions : %v\n", err)
		return []entity.TransactionResponse{}, nil
	}

	return transaction, nil
}

func (r *transactionRepository) DeleteTransaction(ctx context.Context, id int) error {

	var transaction entity.Transaction
	if err := r.db.WithContext(ctx).Select("id", "walletid", "trxtype", "walletsourceid", "description", "amount", "created_at", "updated_at").First(&transaction, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("transaction not found")
		}
	}

	if err := r.db.WithContext(ctx).Delete(&entity.Transaction{}, id).Error; err != nil {
		log.Printf("Error deleting transaction : %v\n", err)
		return err
	}

	var walletExisting entity.Wallet
	if err := r.db.WithContext(ctx).Where("id = ?", transaction.WalletID).Select("id, name, user_id, description, balance, created_at, updated_at").First(&walletExisting).Error; err != nil {
		log.Printf("Error finding wallet : %v\n", err)
		return nil
	}

	if transaction.TrxType == enum.TOPUP {
		walletExisting.Balance -= transaction.Amount
		walletExisting.UpdatedAt = time.Now()
	} else if transaction.TrxType == enum.PAYMENT {
		walletExisting.Balance += transaction.Amount
		walletExisting.UpdatedAt = time.Now()
	} else if transaction.TrxType == enum.TRANSFERIN {
		walletExisting.Balance -= transaction.Amount
		walletExisting.UpdatedAt = time.Now()
	} else if transaction.TrxType == enum.TRANSFEROUT {
		walletExisting.Balance += transaction.Amount
		walletExisting.UpdatedAt = time.Now()
	}

	if err := r.db.WithContext(ctx).Save(&walletExisting).Error; err != nil {
		log.Printf("Error update wallet after delete transaction : %v\n", err)
		return nil
	}

	return nil
}

func (r *transactionRepository) UpdateWalletBalance(ctx context.Context, id int, amount float32, trxType string) error {
	var walletExisting entity.Wallet
	if err := r.db.WithContext(ctx).Where("id = ?", id).Select("id, name, user_id, description, balance, created_at, updated_at").First(&walletExisting).Error; err != nil {
		log.Printf("Error finding wallet : %v\n", err)
		return nil
	}

	if (trxType == enum.PAYMENT) && walletExisting.Balance < amount {

		log.Printf("Error insufficient balance")
		return fmt.Errorf("error insufficient balance")
	}

	if trxType == enum.TRANSFEROUT {
		var walletTrsOut entity.Wallet
		if err := r.db.WithContext(ctx).Where("id = ?", id).Select("id, name, user_id, description, balance, created_at, updated_at").First(&walletTrsOut).Error; err != nil {
			log.Printf("Error finding wallet : %v\n", err)
			return nil
		}

		if walletTrsOut.Balance < amount {
			log.Printf("Error insufficient balance for transfer")
			return fmt.Errorf("error insufficient balance transfer")
		}
	}

	if trxType == enum.TOPUP {
		walletExisting.Balance += amount
		walletExisting.UpdatedAt = time.Now()
	} else if trxType == enum.PAYMENT {
		walletExisting.Balance -= amount
		walletExisting.UpdatedAt = time.Now()
	} else if trxType == enum.TRANSFERIN {
		walletExisting.Balance += amount
		walletExisting.UpdatedAt = time.Now()
	} else if trxType == enum.TRANSFEROUT {
		walletExisting.Balance -= amount
		walletExisting.UpdatedAt = time.Now()
	}
	if err := r.db.WithContext(ctx).Save(&walletExisting).Error; err != nil {
		return nil
	}

	return nil
}
