package main

import context "context"

type numbersService struct {
	X int64
	Y int64
}

func NewNumbersService() NumbersStorageServer {
	return &numbersService{
		X: 0,
		Y: 0,
	}
}

func (ns *numbersService) SetNumbers(ctx context.Context, req *SetNumbersRequest) (res *SetNumbersResponse, err error) {
	ns.X = req.GetX()
	ns.Y = req.GetY()

	return &SetNumbersResponse{}, nil
}

func (ns *numbersService) GetNumbers(ctx context.Context, req *GetNumbersRequest) (*GetNumbersResponse, error) {
	return &GetNumbersResponse{
		X: ns.X,
		Y: ns.Y,
	}, nil
}
