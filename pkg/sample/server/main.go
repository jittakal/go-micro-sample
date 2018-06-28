//go:generate protoc -I ../service --go_out=plugins=grpc:../../../ ../service/sample.proto

package main

import "context"

const (
	port = ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{message: "Hello " + in.Name}, nil
}

func main() {

}
