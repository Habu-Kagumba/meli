package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	proto "github.com/Habu-Kagumba/meli/consignment-service/proto/consignment"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*proto.Consignment, error) {
	var consignment *proto.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	cmd.Init()

	client := proto.NewShippingService("go.micro.srv.consignment", client.DefaultClient)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not create consignment(s): %v", err)
	}

	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &proto.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
