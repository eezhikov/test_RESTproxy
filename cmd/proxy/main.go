package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"testTask/pkg/test_service/config"
	gw "testTask/pkg/test_service/server"
)

func run(cfg *config.Config) error {

	grpcServerEndpoint := fmt.Sprintf("service:%d", cfg.GrpcPort)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := gw.RegisterTestServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts); err != nil {
		return err
	}

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.RestPort), mux)
}
func main() {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	))
	cfg := config.New()
	if err := run(cfg); err != nil {
		logger.Error("run error", zap.Error(err))
	}
}
