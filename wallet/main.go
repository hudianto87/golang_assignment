package main

import (
	"context"
	grpcHanlder "golangassignment/wallet/handler/grpc"
	pb "golangassignment/wallet/proto/wallet_service/v1"
	postgresgorm "golangassignment/wallet/repository/postgres_gorm"
	"golangassignment/wallet/service"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	dsn := "postgresql://postgres:P4ssw0rd@192.168.26.50:5432/traininggolang"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalln(err)
	}
	pgxpool.New(context.Background(), dsn)

	//untuk menggunakan gorm
	walletRepo := postgresgorm.NewWalletRepository(gormDB)
	walletService := service.NewWalletService(walletRepo)
	walletHanlder := grpcHanlder.NewWalletHandler(walletService)

	transactionRepo := postgresgorm.NewTransactionRepository(gormDB)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHanlder := grpcHanlder.NewTransactionHandler(transactionService)

	//run grpc server
	grpcServer := grpc.NewServer()
	// Register WalletService
	pb.RegisterWalletServiceServer(grpcServer, walletHanlder)

	// Register TransactionService (this was missing before)
	pb.RegisterTransactionServiceServer(grpcServer, transactionHanlder)

	lis, err := net.Listen("tcp", "localhost:50052")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("running server on port : 50052")
	grpcServer.Serve(lis)
}
