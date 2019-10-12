package server

import (
	"context"
	"github.com/micro/go-micro/registry"
	"testing"

	pb "github.com/micro/go-micro/server/proto"
	"github.com/micro/go-micro/util/log"
)

type MyServer struct{}

func (s *MyServer) Handle(ctx context.Context, req *pb.HandleRequest, rsp *pb.HandleResponse) error {
	log.Log("Received Test.Handle request")
	return nil
}

func (s *MyServer) Subscribe(ctx context.Context, req *pb.SubscribeRequest, rsp *pb.SubscribeResponse) error {
	log.Log("Received Test.Handle request")
	return nil
}

func TestNewRpcHandler(t *testing.T) {
	newRpcHandler(&MyServer{})
}