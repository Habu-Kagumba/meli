package main

import (
	"fmt"

	proto "github.com/Habu-Kagumba/meli/consignment-service/proto/consignment"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*proto.Consignment) (*proto.Consignment, error)
	GetAll() []*proto.Consignment
}

// Repository simulates the use of a real store
type Repository struct {
	consignments []*proto.Consignment
}

func (repo *Repository) Create(consignment *proto.Consignment) (*proto.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*proto.Consignment {
	return repo.consignments
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *proto.Consignment, resp *proto.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = consignment

	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *proto.GetRequest, resp *proto.Response) error {
	consignments := s.repo.GetAll()

	resp.Consignments = consignments

	return nil
}

func main() {
	repo := &Repository{}

	srv := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
