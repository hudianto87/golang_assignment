package router

import (
	"golangassignment/gateway/handler/grpc"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, gatewayHandler *grpc.GatewayHandler) {
	// User routes
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", gatewayHandler.CreateUser)
		userRoutes.GET("/:id", gatewayHandler.GetUserByID)
		userRoutes.PUT("/:id", gatewayHandler.UpdateUser)
		userRoutes.DELETE("/:id", gatewayHandler.DeleteUser)
		userRoutes.GET("", gatewayHandler.GetAllUsers)
	}

	// Wallet routes
	walletRoutes := r.Group("/wallets")
	{
		walletRoutes.POST("", gatewayHandler.CreateWallet)
		walletRoutes.GET("/:id", gatewayHandler.GetWalletByID)
		walletRoutes.PUT("/:id", gatewayHandler.UpdateWallet)
		walletRoutes.DELETE("/:id", gatewayHandler.DeleteWallet)
		walletRoutes.GET("", gatewayHandler.GetAllWallets)
	}

	// Transaction routes
	transactionRoutes := r.Group("/transactions")
	{
		transactionRoutes.POST("", gatewayHandler.CreateTransaction)
		transactionRoutes.GET("/:id", gatewayHandler.GetTransactionByID)
		transactionRoutes.GET("/wallet/:walletid", gatewayHandler.GetTransactionByWalletID)
		transactionRoutes.DELETE("/:id", gatewayHandler.DeleteTransaction)
	}
}
