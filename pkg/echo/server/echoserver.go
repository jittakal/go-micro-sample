//go:generate protoc -I ../service --go_out=plugins=grpc:../../../ ../service/echo.proto

package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/jittakal/go-micro-sample/pkg/echo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "0.0.0.0:50051"
)

type server struct{}

func (s *server) DoEcho(ctx context.Context, in *pb.Request) (*pb.Reply, error) {
	return &pb.Reply{Message: in.Message}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	fmt.Println(lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
