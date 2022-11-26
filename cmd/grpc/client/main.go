package main

import (
	"context"
	"fmt"

	"github.com/masudur-rahman/repo-management-svc/internal/gRPC/domain"
	"github.com/masudur-rahman/repo-management-svc/internal/gRPC/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := "localhost:9000"
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := service.NewRepositoryServiceClient(conn)
	for i := range [10]int{} {
		rpModel := domain.Repository{
			Id:        int64(i),
			Name:      "gRPC-demo",
			UserID:    1245,
			IsPrivate: true,
		}

		resp, err := client.Add(context.Background(), &rpModel)
		if err != nil {
			panic("was not able to insert record, " + err.Error())
		}

		fmt.Println("Record Inserted..")
		fmt.Println(resp)
		fmt.Println("====================================")
	}

}
