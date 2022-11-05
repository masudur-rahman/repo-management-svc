package service

import (
	"context"
	"log"
	"strconv"

	"github.com/masudur-rahman/repo-management-svc/internal/gRPC/domain"
)

type RepositoryServiceGrpcImpl struct{}

func NewRepositoryServiceGrpcImpl() *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{}
}

func (serviceImpl *RepositoryServiceGrpcImpl) Add(ctx context.Context, in *domain.Repository) (*AddRepositoryResponse, error) {
	log.Println("Received request for adding repository with id " + strconv.FormatInt(in.Id, 10))

	// Logic to persist to database or storage.
	log.Println("Repository persisted to the storage")

	return &AddRepositoryResponse{
		AddedRepository: in,
		Error:           nil,
	}, nil
}

func (serviceImpl *RepositoryServiceGrpcImpl) mustEmbedUnimplementedRepositoryServiceServer() {
	//TODO implement me
	panic("implement me")
}
