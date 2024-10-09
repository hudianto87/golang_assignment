package main

import (
	"context"
	grpcHanlder "golangassignment/user/handler/grpc"
	pb "golangassignment/user/proto/user_service/v1"
	postgresgorm "golangassignment/user/repository/postgres_gorm"
	"golangassignment/user/service"
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
	userRepo := postgresgorm.NewUserRepository(gormDB)
	userService := service.NewUserService(userRepo)
	userHanlder := grpcHanlder.NewUserHandler(userService)

	//run grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHanlder)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("running server on port : 50051")
	grpcServer.Serve(lis)
}
