package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	vessel "github.com/lwu1618/shipping/srv/vessel/proto/vessel"
)

type Vessel struct{}

func (e *Vessel) Handle(ctx context.Context, msg *vessel.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *vessel.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
