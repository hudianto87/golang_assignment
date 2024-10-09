package postgresgorm

import (
	"context"
	"errors"
	"golangassignment/wallet/entity"
	"log"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
}

type IWalletRepository interface {
	CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error)
	GetWalletByID(ctx context.Context, id int) (entity.Wallet, error)
	UpdateWallet(ctx context.Context, id int, wallet entity.Wallet) (entity.Wallet, error)
	DeleteWallet(ctx context.Context, id int) error
	GetAllWallets(ctx context.Context) ([]entity.Wallet, error)
}

type walletRepository struct {
	db GormDBIface
}

func NewWalletRepository(db GormDBIface) IWalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error) {
	if err := r.db.WithContext(ctx).Create(wallet).Error; err != nil {
		log.Printf("Error creating wallet : %v\n", err)
	}

	return *wallet, nil
}

func (r *walletRepository) GetWalletByID(ctx context.Context, id int) (entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.WithContext(ctx).Select("id", "name", "user_id", "description", "balance", "created_at", "updated_at").First(&wallet, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Wallet{}, nil
		}
		log.Printf("Error get wallet : %v\n", err)
		return entity.Wallet{}, nil
	}

	return wallet, nil
}

func (r *walletRepository) UpdateWallet(ctx context.Context, id int, wallet entity.Wallet) (entity.Wallet, error) {
	var walletExisting entity.Wallet
	if err := r.db.WithContext(ctx).Select("id", "name", "user_id", "description", "balance", "created_at", "updated_at").First(&walletExisting, id).Error; err != nil {
		log.Printf("Error finding wallet : %v\n", err)
		return entity.Wallet{}, nil
	}
	walletExisting.Name = wallet.Name
	walletExisting.UserID = wallet.UserID
	walletExisting.Description = wallet.Description

	if err := r.db.WithContext(ctx).Save(&walletExisting).Error; err != nil {
		return entity.Wallet{}, nil
	}

	return walletExisting, nil
}

func (r *walletRepository) DeleteWallet(ctx context.Context, id int) error {
	if err := r.db.WithContext(ctx).Delete(&entity.Wallet{}, id).Error; err != nil {
		log.Printf("Error deleting wallet : %v\n", err)
		return err
	}

	return nil
}

func (r *walletRepository) GetAllWallets(ctx context.Context) ([]entity.Wallet, error) {
	var wallets []entity.Wallet
	if err := r.db.WithContext(ctx).Select("id", "name", "user_id", "description", "balance", "created_at", "updated_at").Find(&wallets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return wallets, nil
		}
		log.Printf("Error get wallet : %v\n", err)
		return nil, err
	}

	return wallets, nil
}
