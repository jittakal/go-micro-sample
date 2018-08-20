package client

import (
	"fmt"
	"log"
	"time"

	pb "github.com/jittakal/go-micro-sample/pkg/echo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address        = ":50051"
	defaultMessage = "Hello, World!"
)

// DoEcho print echo message reply from DoEcho service
func DoEcho(message string) error {
	// Setup a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := pb.NewEchoClient(conn)

	r, err := c.DoEcho(ctx, &pb.Request{Message: message})
	if err != nil {
		return err
	}
	log.Printf("Echo message is `%s`", r.Message)
	fmt.Println(r.Message)
	return nil
}
