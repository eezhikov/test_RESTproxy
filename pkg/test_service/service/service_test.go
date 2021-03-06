package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
	"os"
	"testTask/pkg/test_service/config"
	pb "testTask/pkg/test_service/server"
	"testTask/pkg/test_service/service"
	"testing"
)

func TestConcStrings(t *testing.T) {
	var StringsRequest = []struct {
		in     *pb.ConcStringsRequest
		answer string
		err    error
	}{
		{
			&pb.ConcStringsRequest{
				FirstStr:  "qwerqwe",
				SecondStr: "qwerqwe",
			},
			"qwerqweqwerqwe",
			nil,
		},
		{
			&pb.ConcStringsRequest{
				FirstStr:  "",
				SecondStr: "qwerqwe",
			},
			"",
			errors.New("first string is empty"),
		},
		{
			&pb.ConcStringsRequest{
				FirstStr:  "qwerqwe",
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
		if err != nil && !errors.As(err, &testValue.err) {
			t.Fatalf("Cant cont strings: %v", err)
		} else if res != nil && res.Result != testValue.answer {
			t.Errorf("ConcStrings(%s), expected %s", testValue.in, res.Result)
		}
	}
}

func Test(t *testing.T) {
	cfg := config.New()
	url := fmt.Sprintf("http://%s:%d/strings", cfg.Host, cfg.RestPort)
	client := http.Client{}

	type req struct {
		FirstStr  string `json:"firstStr"`
		SecondStr string `json:"secondStr"`
	}

	var StringsReq = []struct {
		in     req
		answer string
	}{
		{
			in: req{
				FirstStr:  "qwer",
				SecondStr: "tyui",
			},
			answer: "{\"result\":\"qwertyui\"}",
		},
	}

	for _, testValue := range StringsReq {
		body, err := json.Marshal(testValue.in)
		if err != nil {
			t.Fatalf("Marshal error: %v", err)
			return
		}
		reqBody := bytes.NewReader(body)

		resp, err := client.Post(url, "application/json", reqBody)
		if err != nil {
			t.Fatalf("Request error: %v", err)
		}
		if resp == nil {
			t.Fatal("Nil response from REST service:")
		}
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return

		}
		if string(respBody) != testValue.answer {
			t.Errorf("Test(%s), expected %s", testValue.in, string(respBody))
		}
		resp.Body.Close()
	}
}
func BenchmarkTestService_ConcStrings(b *testing.B) {
	strReq := &pb.ConcStringsRequest{
		FirstStr:  "qwer",
		SecondStr: "tyui",
	}

	for i := 0; i < b.N; i++ {
		logger := zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			os.Stdout,
			zap.DebugLevel,
		))

		testService := service.New(logger)

		ctx := context.Background()
		_, err := testService.ConcStrings(ctx, strReq)
		if err != nil {
			b.Fatalf("Cant ConcStrings: %v", err)
		}
	}
}
