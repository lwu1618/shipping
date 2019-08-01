package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	consignment "github.com/lwu1618/shipping/srv/consignment/proto/consignment"
)

type Consignment struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Consignment) Call(ctx context.Context, req *consignment.Request, rsp *consignment.Response) error {
	log.Log("Received Consignment.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Consignment) Stream(ctx context.Context, req *consignment.StreamingRequest, stream consignment.Consignment_StreamStream) error {
	log.Logf("Received Consignment.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&consignment.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Consignment) PingPong(ctx context.Context, stream consignment.Consignment_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&consignment.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
