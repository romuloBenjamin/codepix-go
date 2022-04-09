package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	"github.com/romuloBenjamin/codepix-go/application/grpc/pb"
	"github.com/romuloBenjamin/codepix-go/application/usecase"
	"github.com/romuloBenjamin/codepix-go/infraestructure/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	Pixrepository := repository.PixKeyRepositoryDb{Db: database}
	PixUseCase := usecase.PixUseCase{PixKeyRepository: Pixrepository}
	pixGrpcService := NewPixGrpcService(PixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("connot start grpc server", err)
	}
	log.Printf("grpx has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("connot start grpc server", err)
	}
}
