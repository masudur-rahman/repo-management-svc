package main

import (
	"fmt"
	"log"
	"net"

	"github.com/masudur-rahman/repo-management-svc/internal/gRPC/service"
	"google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(9000)
	grpcServer := grpc.NewServer()

	rpSvcImpl := service.NewRepositoryServiceGrpcImpl()
	service.RegisterRepositoryServiceServer(grpcServer, rpSvcImpl)

	if err := grpcServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return lis
}
