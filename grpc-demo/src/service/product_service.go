package service

import (
	"context"
	"errors"

	pb "github.com/rohitk12r/src/grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.SearchProductServer
}

func (grpcService *server) FindProductById(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	productList := findProduct()
	for _, value := range productList {
		if value.ProductId == req.ProductId {
			return value, nil
		}
	}
	return nil, errors.New("No data found")
}

func findProduct() []*pb.ProductResponse {
	var productList []*pb.ProductResponse
	product1 := &pb.ProductResponse{
		ProductId:   1,
		ProductName: "Laptop",
		Description: "Laptop is single device which is perform multiple operation",
	}
	product2 := &pb.ProductResponse{
		ProductId:   2,
		ProductName: "Monitor",
		Description: "Monitor is a display",
	}
	product3 := &pb.ProductResponse{
		ProductId:   3,
		ProductName: "Printer",
		Description: "Printer is machine to print the pages.",
	}
	productList = append(productList, product1, product2, product3)
	return productList
}

func RegisterGRPCProto() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterSearchProductServer(s, &server{})
	return s
}
