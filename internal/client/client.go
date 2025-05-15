package client

import (
	"context"
	pb "github.com/Suhach/test_protoc-cont/proto/user"
	"google.golang.org/grpc"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &UserClient{client: pb.NewUserServiceClient(conn)}, nil
}

func (c *UserClient) GetUser(ctx context.Context, id uint32) (*pb.GetUserResponse, error) {
	return c.client.GetUser(ctx, &pb.GetUserRequest{Id: id})
}
