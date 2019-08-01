package main

import (
	"github.com/lwu1618/shipping/srv/consignment/handler"
	"github.com/lwu1618/shipping/srv/consignment/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	consignment "github.com/lwu1618/shipping/srv/consignment/proto/consignment"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	consignment.RegisterConsignmentHandler(service.Server(), new(handler.Consignment))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.consignment", service.Server(), new(subscriber.Consignment))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.consignment", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
