package service_test

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	pb "testTask/pkg/test_service/server"
	"testTask/pkg/test_service/service"
	"testing"
)



func TestConcStrings(t *testing.T) {
	var StringsRequest = []struct {
		in *pb.ConcStringsRequest
		answer string
		err error
	}{
		{
			&pb.ConcStringsRequest{
				FirstStr: "qwerqwe",
				SecondStr: "qwerqwe",
			},
			"qwerqweqwerqwe",
			nil,
		},
		{
			&pb.ConcStringsRequest{
				FirstStr: "",
				SecondStr: "qwerqwe",
			},
			"",
			errors.New("first string is empty"),


		},
		{
			&pb.ConcStringsRequest{
				FirstStr: "qwerqwe",
				SecondStr: "",
			},
			"",
			errors.New("second string is empty"),
		},

	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		os.Stdout,
		zap.DebugLevel,
	))

	testService := service.New(logger)


	ctx := context.Background()
	for _, testValue := range StringsRequest {
                res, err := testService.ConcStrings(ctx, testValue.in)
				if err != nil && !errors.As(err, &testValue.err){
					t.Fatalf("Cant cont strings: %v", err)
				}else if res != nil && res.Result != testValue.answer {
                        t.Errorf("ConcStrings(%s), expected %s", testValue.in, res.Result)
                }

        }
}

