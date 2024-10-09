package main

import (
	handler "golangassignment/gateway/handler/grpc"
	pb "golangassignment/gateway/proto/gateway_service/v1"
	"golangassignment/gateway/router"
	"golangassignment/gateway/service"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Use grpc.NewClient to create a new gRPC connection
	connU, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to create gRPC client:", err)
	}
	defer connU.Close()

	connWT, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to create gRPC client:", err)
	}
	defer connWT.Close()

	// Create gRPC service clients
	userServiceClient := pb.NewUserServiceClient(connU)
	walletServiceClient := pb.NewWalletServiceClient(connWT)
	transactionServiceClient := pb.NewTransactionServiceClient(connWT)

	// Initialize the Gateway Service with gRPC clients
	gatewayService := service.NewGatewayService(userServiceClient, walletServiceClient, transactionServiceClient)

	// Initialize HTTP Handler for Gin
	gatewayHandler := handler.NewGatewayHandler(gatewayService)

	// Set up Gin routes
	r := gin.Default()

	// Setup router by passing the handler
	router.SetupRouter(r, gatewayHandler)

	// Run the Gin server
	log.Fatal(r.Run(":8080"))
}
