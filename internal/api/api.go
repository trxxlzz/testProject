package api

import (
	"context"
	"testProject/internal/service"
	pb "testProject/protos/gen/go"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Service *service.UserService
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	id, err := s.Service.CreateUser(ctx, req.Name, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Id: id}, nil
}
