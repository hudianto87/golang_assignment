package service

import (
	"context"
	"fmt"
	"golangassignment/wallet/entity"
	postgresgorm "golangassignment/wallet/repository/postgres_gorm"
)

type IWalletService interface {
	CreateWallet(ctx context.Context, user *entity.Wallet) (entity.Wallet, error)
	GetWalletByID(ctx context.Context, id int) (entity.Wallet, error)
	UpdateWallet(ctx context.Context, id int, user entity.Wallet) (entity.Wallet, error)
	DeleteWallet(ctx context.Context, id int) error
	GetAllWallets(ctx context.Context) ([]entity.Wallet, error)
}

// untuk menggunakan gorm
type walletService struct {
	walletRepo postgresgorm.IWalletRepository
}

// untuk menggunakan gorm
func NewWalletService(walletRepo postgresgorm.IWalletRepository) IWalletService {
	return &walletService{walletRepo: walletRepo}
}

func (r *walletService) CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error) {
	createdWallet, err := r.walletRepo.CreateWallet(ctx, wallet)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error created wallet: %v", err)
	}

	return createdWallet, nil
}

func (r *walletService) GetWalletByID(ctx context.Context, id int) (entity.Wallet, error) {
	wallet, err := r.walletRepo.GetWalletByID(ctx, id)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error wallet not found: %v", err)
	}

	return wallet, nil
}

func (r *walletService) UpdateWallet(ctx context.Context, id int, wallet entity.Wallet) (entity.Wallet, error) {
	updatedWallet, err := r.walletRepo.UpdateWallet(ctx, id, wallet)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error wallet not found: %v", err)
	}

	return updatedWallet, nil
}

func (r *walletService) DeleteWallet(ctx context.Context, id int) error {
	err := r.walletRepo.DeleteWallet(ctx, id)
	if err != nil {
		return fmt.Errorf("error wallet not found: %v", err)
	}

	return nil
}

func (r *walletService) GetAllWallets(ctx context.Context) ([]entity.Wallet, error) {
	wallets, err := r.walletRepo.GetAllWallets(ctx)

	if err != nil {
		return nil, fmt.Errorf("error to retrieve data wallets : %v", err)
	}

	return wallets, nil
}
