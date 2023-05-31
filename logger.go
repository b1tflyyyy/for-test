package main

import (
	context "context"
	fmt "fmt"
	"time"
)

type loggerService struct {
	service NumbersStorageServer
}

func NewLoggerService(service NumbersStorageServer) NumbersStorageServer {
	return &loggerService{
		service: service,
	}
}

func (ls *loggerService) SetNumbers(ctx context.Context, req *SetNumbersRequest) (res *SetNumbersResponse, err error) {
	defer func(start time.Time) {
		fmt.Printf("req=%v err=%v took=%v\n\n", req.String(), err, time.Since(start))
	}(time.Now())

	return ls.service.SetNumbers(ctx, req)
}

func (ls *loggerService) GetNumbers(ctx context.Context, req *GetNumbersRequest) (res *GetNumbersResponse, err error) {
	defer func(start time.Time) {
		fmt.Printf("err=%v res=%v\ntook=%v\n\n", err, string(res.GetX())+" "+string(res.GetY()), time.Since(start))
	}(time.Now())

	return ls.service.GetNumbers(ctx, req)
}
