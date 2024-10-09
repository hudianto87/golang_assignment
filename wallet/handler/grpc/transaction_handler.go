package grpc

import (
	"context"
	"fmt"
	"golangassignment/wallet/entity"
	"golangassignment/wallet/service"
	"log"

	pb "golangassignment/wallet/proto/wallet_service/v1"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ITransactionHandler interface {
	CreateTransaction(c *gin.Context)
	GetTransactionByID(c *gin.Context)
	GetTransactionByWalletID(c *gin.Context)
	DeleteTransaction(c *gin.Context)
}

type TransactionHandler struct {
	pb.UnimplementedTransactionServiceServer
	transactionService service.ITransactionService
}

func NewTransactionHandler(tranasctioinService service.ITransactionService) *TransactionHandler {
	return &TransactionHandler{
		transactionService: tranasctioinService,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.MutationTransactionResponse, error) {
	//log.Printf("Received CreateTransactionRequest: WalletID: %d, TrxType: %s, Description: %s", req.Walletid, req.Trxtype, req.Description)
	createdTransaction, err := h.transactionService.CreateTransaction(ctx, &entity.Transaction{
		WalletID:       int(req.Walletid),
		TrxType:        req.Trxtype,
		WalletSourceID: int(req.Walletsourceid),
		Description:    req.Description,
		Amount:         req.Amount,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationTransactionResponse{
		Message: fmt.Sprintf("Success created transaction with id %d", createdTransaction.ID),
	}, nil
}

func (h *TransactionHandler) GetTransactionByID(ctx context.Context, req *pb.GetTransactionByIDRequest) (*pb.GetTransactoinByIDResponse, error) {
	transaction, err := h.transactionService.GetTransactionByID(ctx, int(req.Id))

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(req.Id)
	res := &pb.GetTransactoinByIDResponse{
		Transaction: &pb.Transaction{
			Id:             int32(transaction.ID),
			Walletid:       int32(transaction.WalletID),
			Trxtype:        transaction.TrxType,
			Walletsourceid: int32(transaction.WalletSourceID),
			Description:    transaction.Description,
			Amount:         transaction.Amount,
			CreatedAt:      timestamppb.New(transaction.CreatedAt),
			UpdatedAt:      timestamppb.New(transaction.UpdatedAt),
		},
	}

	return res, nil
}

func (h *TransactionHandler) GetTransactionByWalletID(ctx context.Context, req *pb.GetTransactionByWalletIDRequest) (*pb.GetTransactoinByWalletIDResponse, error) {
	wallettransactions, err := h.transactionService.GetTransactionByWalletID(ctx, int(req.Walletid))

	if err != nil {
		log.Println(err)
		return nil, err
	}
	var transactionProtos []*pb.TransactionByWalletID

	for _, transaction := range wallettransactions {
		transactionProtos = append(transactionProtos, &pb.TransactionByWalletID{
			Id:             int32(transaction.ID),
			Walletid:       int32(transaction.WalletID),
			Walletname:     transaction.WalletName,
			Trxtype:        transaction.TrxType,
			Walletsourceid: int32(transaction.WalletSourceID),
			Amount:         float32(transaction.Amount),
			Name:           transaction.UserName,
			CreatedAt:      timestamppb.New(transaction.CreatedAt),
			UpdatedAt:      timestamppb.New(transaction.UpdatedAt),
		})
	}

	return &pb.GetTransactoinByWalletIDResponse{
		Transactions: transactionProtos,
	}, nil
}

func (h *TransactionHandler) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.MutationTransactionResponse, error) {
	if err := h.transactionService.DeleteTransaction(ctx, int(req.Idtrx)); err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationTransactionResponse{
		Message: fmt.Sprintf("Success delete transaction with id %d", req.Idtrx),
	}, nil
}
