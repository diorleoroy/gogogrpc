package main

import (
	"context"
	"fmt"
	"net"
	"log"
	"github.com/diorleoroy/gogogrpc/gen/pb/gogogrpc"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		panic(fmt.Errorf("run: %w", err))
	}
}

func run() error {
	listenOn := ":20000"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return fmt.Errorf("Listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()

	gogogrpc.RegisterGreeterServer(server, &GreeterServerHandler{})
	log.Println("Listening on", listenOn)
	if err := server.Serve(listener); err !=nil {
		return fmt.Errorf("serve gRPC server: %w", err)
	}

	return nil
}

var _ gogogrpc.GreeterServer = &GreeterServerHandler{}

type GreeterServerHandler struct {
	gogogrpc.UnimplementedGreeterServer
}

// func (h *GreeterServerHandler) SayHello(
// 	ctx context.Context,req *gogogrpc.HelloRequest,
// ) (*gogogrpc.HelloReply, error) {
// 	fmt.Printf("hello",req.Name)
// 	return nil, nil
// }

func (h *GreeterServerHandler) SayHello(
	ctx context.Context, req *gogogrpc.HelloRequest,
) (*gogogrpc.HelloReply, error) {
return &gogogrpc.HelloReply{
	Message: fmt.Sprintln("hello", req.Name),
}, nil
}


