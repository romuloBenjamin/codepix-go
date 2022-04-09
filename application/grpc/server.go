package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:#{port}")
	listner, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start")
	}

	log.Printf("gRPC start port %d", port)
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("not start grpc", err)
	}
}
