package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	consignment "github.com/lwu1618/shipping/srv/consignment/proto/consignment"
)

type Consignment struct{}

func (e *Consignment) Handle(ctx context.Context, msg *consignment.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *consignment.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
