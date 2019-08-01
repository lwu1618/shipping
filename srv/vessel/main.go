package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/lwu1618/shipping/srv/vessel/handler"
	"github.com/lwu1618/shipping/srv/vessel/subscriber"

	vessel "github.com/lwu1618/shipping/srv/vessel/proto/vessel"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	vessel.RegisterVesselHandler(service.Server(), new(handler.Vessel))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.vessel", service.Server(), new(subscriber.Vessel))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.vessel", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
