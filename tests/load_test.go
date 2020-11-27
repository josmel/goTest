package tests

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/josmel/br-seed-go/proto/consignment"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}
func Test_GetConsignment(t *testing.T) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewShippingServiceClient(conn)

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		t.Fatalf("Could not parse file: %v", err)
	}

	res, err := c.CreateConsignment(context.Background(), consignment)

	if err != nil {
		t.Fatalf("Could not greet: %v", err)
	}
	t.Logf("Created: %t", res.Created)

}
