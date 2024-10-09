package service

import (
	"context"
	"fmt"
	"golangassignment/gateway/entity"
	transaction_Service "golangassignment/gateway/proto/gateway_service/v1"
	user_Service "golangassignment/gateway/proto/gateway_service/v1"
	wallet_Service "golangassignment/gateway/proto/gateway_service/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

type IGatewayService interface {
	CreateUser(ctx context.Context, user *entity.User) (entity.User, error)
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]entity.User, error)

	CreateWallet(ctx context.Context, user *entity.Wallet) (entity.Wallet, error)
	GetWalletByID(ctx context.Context, id int) (entity.Wallet, error)
	UpdateWallet(ctx context.Context, id int, user entity.Wallet) (entity.Wallet, error)
	DeleteWallet(ctx context.Context, id int) error
	GetAllWallets(ctx context.Context) ([]entity.Wallet, error)

	CreateTransaction(ctx context.Context, user *entity.Transaction) (entity.Transaction, error)
	GetTransactionByID(ctx context.Context, id int) (entity.Transaction, error)
	GetTransactionByWalletID(ctx context.Context, userid int) ([]entity.TransactionResponse, error)
	DeleteTransaction(ctx context.Context, id int) error
}

type gatewayService struct {
	userService        user_Service.UserServiceClient
	walletService      wallet_Service.WalletServiceClient
	transactionService transaction_Service.TransactionServiceClient
}

func NewGatewayService(user_Service user_Service.UserServiceClient, wallet_Service wallet_Service.WalletServiceClient, transaction_Service transaction_Service.TransactionServiceClient) IGatewayService {
	return &gatewayService{
		userService:        user_Service,
		walletService:      wallet_Service,
		transactionService: transaction_Service,
	}
}

func (g *gatewayService) CreateUser(ctx context.Context, user *entity.User) (entity.User, error) {
	req := &user_Service.CreateUserRequest{
		Name:  user.Name,
		Email: user.Email,
	}
	res, err := g.userService.CreateUser(ctx, req)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to call CreateUser on user service: %v", err)
	}
	return entity.User{
		ID:    int(res.Id),
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func (g *gatewayService) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	req := &user_Service.GetUserByIDRequest{Id: int32(id)}
	res, err := g.userService.GetUserByID(ctx, req)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to get user by ID: %v", err)
	}
	return entity.User{
		ID:    int(res.User.Id),
		Name:  res.User.Name,
		Email: res.User.Email,
	}, nil
}

func (g *gatewayService) UpdateUser(ctx context.Context, id int, user entity.User) (entity.User, error) {
	req := &user_Service.UpdateUserRequest{
		Id:      int32(id),
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}
	updatedUser, err := g.userService.UpdateUser(ctx, req)
	if err != nil {
		return entity.User{}, fmt.Errorf("error user not found: %v", err)
	}

	return entity.User{
		ID:    int(updatedUser.Id),
		Name:  updatedUser.Name,
		Email: updatedUser.Email,
	}, nil
}

func (g *gatewayService) DeleteUser(ctx context.Context, id int) error {
	req := &user_Service.DeleteUserRequest{
		Id: int32(id),
	}

	_, err := g.userService.DeleteUser(ctx, req)
	if err != nil {
		return fmt.Errorf("error user not found: %v", err)
	}

	return nil
}

func (g *gatewayService) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	res, err := g.userService.GetUsers(ctx, &emptypb.Empty{})

	if err != nil {
		return nil, fmt.Errorf("error to retrieve data users: %v", err)
	}

	var users []entity.User
	for _, u := range res.Users {
		users = append(users, entity.User{
			ID:        int(u.Id),
			Name:      u.Name,
			Email:     u.Email,
			Address:   u.Address,
			CreatedAt: u.CreatedAt.AsTime(),
			UpdatedAt: u.UpdatedAt.AsTime(),
		})
	}

	return users, nil
}

func (g *gatewayService) CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error) {
	req := &wallet_Service.CreateWalletRequest{
		Name:        wallet.Name,
		UserId:      wallet.UserID,
		Balance:     wallet.Balance,
		Description: wallet.Description,
	}
	res, err := g.walletService.CreateWallet(ctx, req)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("failed to create wallet: %v", err)
	}
	return entity.Wallet{
		ID:          int(res.Id),
		Name:        res.Name,
		UserID:      res.UserId,
		Balance:     res.Balance,
		Description: res.Description,
	}, nil
}

func (g *gatewayService) GetWalletByID(ctx context.Context, id int) (entity.Wallet, error) {
	req := &wallet_Service.GetWalletByIDRequest{Id: int32(id)}
	res, err := g.walletService.GetWalletByID(ctx, req)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("failed to get wallet by ID: %v", err)
	}
	return entity.Wallet{
		ID:          int(res.Wallet.Id),
		Name:        res.Wallet.Name,
		UserID:      res.Wallet.UserId,
		Description: res.Wallet.Description,
		Balance:     res.Wallet.Balance,
		CreatedAt:   res.Wallet.CreatedAt.AsTime(),
		UpdatedAt:   res.Wallet.UpdatedAt.AsTime(),
	}, nil
}

func (g *gatewayService) UpdateWallet(ctx context.Context, id int, wallet entity.Wallet) (entity.Wallet, error) {
	req := &wallet_Service.UpdateWalletRequest{
		Id:          int32(id),
		Name:        wallet.Name,
		UserId:      wallet.UserID,
		Description: wallet.Description,
		Balance:     wallet.Balance,
	}
	res, err := g.walletService.UpdateWallet(ctx, req)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error wallet not found: %v", err)
	}

	return entity.Wallet{
		ID:          int(res.Id),
		Name:        res.Name,
		UserID:      res.UserId,
		Description: res.Description,
		Balance:     res.Balance,
		CreatedAt:   res.CreatedAt.AsTime(),
		UpdatedAt:   res.UpdatedAt.AsTime(),
	}, nil
}

func (g *gatewayService) DeleteWallet(ctx context.Context, id int) error {
	req := &wallet_Service.DeleteWalletRequest{
		Id: int32(id),
	}

	_, err := g.walletService.DeleteWallet(ctx, req)
	if err != nil {
		return fmt.Errorf("error wallet not found: %v", err)
	}

	return nil
}

func (g *gatewayService) GetAllWallets(ctx context.Context) ([]entity.Wallet, error) {
	res, err := g.walletService.GetWallets(ctx, &emptypb.Empty{})

	if err != nil {
		return nil, fmt.Errorf("error to retrieve data wallets: %v", err)
	}

	var wallets []entity.Wallet
	for _, u := range res.Wallets {
		wallets = append(wallets, entity.Wallet{
			ID:          int(u.Id),
			Name:        u.Name,
			UserID:      u.UserId,
			Description: u.Description,
			Balance:     u.Balance,
			CreatedAt:   u.CreatedAt.AsTime(),
			UpdatedAt:   u.UpdatedAt.AsTime(),
		})
	}

	return wallets, nil
}

func (g *gatewayService) GetTransactionByWalletID(ctx context.Context, walletID int) ([]entity.TransactionResponse, error) {
	req := &transaction_Service.GetTransactionByWalletIDRequest{Walletid: int32(walletID)}
	res, err := g.transactionService.GetTransactionByWalletID(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions by wallet ID: %v", err)
	}

	var transactions []entity.TransactionResponse
	for _, t := range res.Transactions {
		transactions = append(transactions, entity.TransactionResponse{
			ID:             int(t.Id),
			WalletID:       int(t.Walletid),
			WalletName:     t.Walletname,
			TrxType:        t.Trxtype,
			WalletSourceID: int(t.Walletsourceid),
			Amount:         t.Amount,
			UserName:       t.Name,
			CreatedAt:      t.CreatedAt.AsTime(),
			UpdatedAt:      t.UpdatedAt.AsTime(),
		})
	}
	return transactions, nil
}

func (g *gatewayService) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error) {
	req := &transaction_Service.CreateTransactionRequest{
		Walletid:       int32(transaction.WalletID),
		Trxtype:        transaction.TrxType,
		Walletsourceid: int32(transaction.WalletSourceID),
		Amount:         transaction.Amount,
		Description:    transaction.Description,
	}
	res, err := g.transactionService.CreateTransaction(ctx, req)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("failed to create transaction: %v", err)
	}

	return entity.Transaction{
		ID:             int(res.Id),
		WalletID:       int(res.Walletid),
		TrxType:        res.Trxtype,
		WalletSourceID: int(res.Walletsourceid),
		Amount:         res.Amount,
		Description:    res.Description,
		CreatedAt:      res.CreatedAt.AsTime(),
		UpdatedAt:      res.UpdatedAt.AsTime(),
	}, nil
}

func (g *gatewayService) GetTransactionByID(ctx context.Context, id int) (entity.Transaction, error) {
	req := &transaction_Service.GetTransactionByIDRequest{Id: int32(id)}
	res, err := g.transactionService.GetTransactionByID(ctx, req)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("failed to get transaction by ID: %v", err)
	}

	return entity.Transaction{
		ID:             int(res.Transaction.Id),
		WalletID:       int(res.Transaction.Walletid),
		TrxType:        res.Transaction.Trxtype,
		WalletSourceID: int(res.Transaction.Walletsourceid),
		Amount:         res.Transaction.Amount,
		Description:    res.Transaction.Description,
		CreatedAt:      res.Transaction.CreatedAt.AsTime(),
		UpdatedAt:      res.Transaction.UpdatedAt.AsTime(),
	}, nil
}

func (g *gatewayService) DeleteTransaction(ctx context.Context, id int) error {
	req := &transaction_Service.DeleteTransactionRequest{Idtrx: int32(id)}
	_, err := g.transactionService.DeleteTransaction(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to delete transaction: %v", err)
	}

	return nil
}
