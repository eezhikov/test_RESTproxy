package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"net"
	"os"
	"testTask/pkg/test_service/config"
	pb "testTask/pkg/test_service/server"
	"testTask/pkg/test_service/service"
)






func main() {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	))

	cfg := config.New()
	testService := service.New(logger)
	srv := grpc.NewServer()

	pb.RegisterTestServiceServer(srv, testService)
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.GrpcPort))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ok")


		if err := srv.Serve(listener); err != nil{
			fmt.Println(err)
			return
		}



}
