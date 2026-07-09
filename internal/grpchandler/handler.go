package grpchandler

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Jurest07/user-service/internal/repository"
	pb "github.com/Jurest07/user-service/proto/user"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	repo   *repository.UserRepo
	logger *slog.Logger
}

func NewUserHandler(repo *repository.UserRepo, logger *slog.Logger) *UserHandler {
	return &UserHandler{
		repo:   repo,
		logger: logger,
	}
}

func (uh *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := uh.repo.GetUserByID(int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return &pb.GetUserResponse{
		Id:        int32(user.ID),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.repo.CreateUser(req.Username, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &pb.CreateUserResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (h *UserHandler) UserExists(ctx context.Context, req *pb.UserExistsRequest) (*pb.UserExistsResponse, error) {
	exists, err := h.repo.UserExists(int(req.Id))
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}

	return &pb.UserExistsResponse{
		Exists: exists,
	}, nil
}
