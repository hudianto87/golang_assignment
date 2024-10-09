package grpc

import (
	"context"
	"fmt"
	"golangassignment/wallet/entity"
	"golangassignment/wallet/service"
	"log"

	pb "golangassignment/wallet/proto/wallet_service/v1"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IWalletHandler interface {
	CreateWallet(c *gin.Context)
	GetWallet(c *gin.Context)
	UpdateWallet(c *gin.Context)
	DeleteWallet(c *gin.Context)
	GetAllWallets(c *gin.Context)
}

type WalletHandler struct {
	pb.UnimplementedWalletServiceServer
	walletService service.IWalletService
}

func NewWalletHandler(walletService service.IWalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

func (h *WalletHandler) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.MutationResponse, error) {
	//log.Printf("Received CreateWalletRequest: UserID: %s, Name: %s, Description: %s, Balance: %f", req.UserId, req.Name, req.Description, req.Balance)

	createdWallet, err := h.walletService.CreateWallet(ctx, &entity.Wallet{
		Name:        req.Name,
		UserID:      req.UserId,
		Description: req.Description,
		Balance:     req.Balance,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success created user with id %d", createdWallet.ID),
	}, nil
}

func (h *WalletHandler) GetWallets(ctx context.Context, _ *emptypb.Empty) (*pb.GetWalletResponse, error) {
	wallets, err := h.walletService.GetAllWallets(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var walletProto []*pb.Wallet

	for _, wallet := range wallets {
		walletProto = append(walletProto, &pb.Wallet{
			Id:          int32(wallet.ID),
			Name:        wallet.Name,
			UserId:      wallet.UserID,
			Description: wallet.Description,
			Balance:     wallet.Balance,
			CreatedAt:   timestamppb.New(wallet.CreatedAt),
			UpdatedAt:   timestamppb.New(wallet.UpdatedAt),
		})
	}

	return &pb.GetWalletResponse{
		Wallets: walletProto,
	}, nil
}

func (h *WalletHandler) GetWalletByID(ctx context.Context, req *pb.GetWalletByIDRequest) (*pb.GetWalletByIDResponse, error) {
	wallet, err := h.walletService.GetWalletByID(ctx, int(req.Id))

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(req.Id)
	res := &pb.GetWalletByIDResponse{
		Wallet: &pb.Wallet{
			Id:          int32(wallet.ID),
			Name:        wallet.Name,
			UserId:      wallet.UserID,
			Description: wallet.Description,
			Balance:     wallet.Balance,
			CreatedAt:   timestamppb.New(wallet.CreatedAt),
			UpdatedAt:   timestamppb.New(wallet.UpdatedAt),
		},
	}

	return res, nil
}

func (h *WalletHandler) UpdateWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.MutationResponse, error) {
	updatedWallet, err := h.walletService.UpdateWallet(ctx, int(req.Id), entity.Wallet{
		Name:        req.Name,
		UserID:      req.UserId,
		Description: req.Description,
		Balance:     req.Balance,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success update wallet with id %d", updatedWallet.ID),
	}, nil
}

func (h *WalletHandler) DeleteWallet(ctx context.Context, req *pb.DeleteRequest) (*pb.MutationResponse, error) {
	if err := h.walletService.DeleteWallet(ctx, int(req.Id)); err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success delete wallet with id %d", req.Id),
	}, nil
}
