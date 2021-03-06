package handler

import (
	"context"
	"github.com/3Rivers/order/model"

	log "github.com/micro/go-micro/v2/logger"

	order "github.com/3Rivers/order/proto/order"
)

type Order struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Order) Call(ctx context.Context, req *order.Request, rsp *order.Response) error {
	log.Info("Received Order.Call request")
	//rsp.Msg = "Hello " + req.Name
	orderObj := &model.Order{}
	orderObj.ID = req.Id
	err := orderObj.GetOrderInfo()
	if err != nil {
		return err
	}
	rsp.User = orderObj.UserName
	rsp.Goods = orderObj.Goods
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Order) Stream(ctx context.Context, req *order.StreamingRequest, stream order.Order_StreamStream) error {
	log.Infof("Received Order.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&order.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Order) PingPong(ctx context.Context, stream order.Order_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&order.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
