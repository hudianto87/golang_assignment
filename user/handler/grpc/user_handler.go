package grpc

import (
	"context"
	"fmt"
	"golangassignment/user/entity"
	"golangassignment/user/service"
	"log"

	pb "golangassignment/user/proto/user_service/v1"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IUserHandler interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.MutationResponse, error) {
	createdUser, err := h.userService.CreateUser(ctx, &entity.User{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success created user with id %d", createdUser.ID),
	}, nil
}

func (h *UserHandler) GetUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetUserResponse, error) {
	users, err := h.userService.GetAllUsers(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var userProto []*pb.User

	for _, user := range users {
		userProto = append(userProto, &pb.User{
			Id:        int32(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Address:   user.Address,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}

	return &pb.GetUserResponse{
		Users: userProto,
	}, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := h.userService.GetUserByID(ctx, int(req.Id))

	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(req.Id)
	res := &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:        int32(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Address:   user.Address,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}

	return res, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.MutationResponse, error) {
	updatedUser, err := h.userService.UpdateUser(ctx, int(req.Id), entity.User{
		Name:    req.Name,
		Email:   req.Email,
		Address: req.Address,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success update user with id %d", updatedUser.ID),
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.MutationResponse, error) {
	if err := h.userService.DeleteUser(ctx, int(req.Id)); err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("Success delete user with id %d", req.Id),
	}, nil
}
