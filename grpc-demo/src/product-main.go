package main

import (
	"fmt"
	"net"

	"github.com/rohitk12r/src/service"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("err")
	}
	server := service.RegisterGRPCProto()
	fmt.Println("Start GRPC server")
	if err := server.Serve(lis); err != nil {
		fmt.Errorf("failed to serve: %v", err)
	}
}
