package user

import (
	"context"

	"github.com/santoshanand/sample-grpc-go/gen"
)

type NewUserService struct{}

func (s *NewUserService) GetUsers(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	res := &gen.Response{
		UserId:    req.UserId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		FullName:  req.FirstName + " " + req.LastName,
	}
	return res, nil
}
