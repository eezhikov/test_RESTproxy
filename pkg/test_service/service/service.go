package service

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	pb "testTask/pkg/test_service/server"
)

type TestService struct {
	pb.TestServiceServer
	logger *zap.Logger
}

func New(logger *zap.Logger) *TestService {
	return &TestService{
		logger: logger,
	}
}
func (t *TestService) ConcStrings(ctx context.Context, request *pb.ConcStringsRequest) (*pb.ConcStringsResponse, error) {

	if request.FirstStr == ""{
		return nil, errors.New("first string is empty")
	}
	fmt.Println("Validate first string")

	if request.SecondStr == ""{
		return nil, errors.New("second string is empty")
	}
	fmt.Println("Validate second string")
	var resp pb.ConcStringsResponse
	resp.Result = request.FirstStr + request.SecondStr
	fmt.Println("Concatenate strings and return result")
	return &resp, nil
}

