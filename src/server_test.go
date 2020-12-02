package main

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/josmel/br-seed-go/proto/consignment"
)

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	pb.RegisterShippingServiceServer(server, &service{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestDepositServer_Deposit(t *testing.T) {
	tests := []struct {
		name   string
		weight int32
		res    *pb.Response
	}{
		{
			"invalid request ",
			11,
			nil,
		},
		{
			"valid request ",
			34,
			&pb.Response{Created: true},
		},
	}

	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewShippingServiceClient(conn)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := &pb.Consignment{
				Description: tt.name,
				Weight:      tt.weight,
			}

			response, err := client.CreateConsignment(ctx, request)

			if response != nil {
				if response.Created != true {
					t.Error("response: expected", true, "received", response.Created)
				}
			}

			if err != nil {
				t.Error("response: expected")
			}

		})
	}
}
