package protoapi

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
)

// GetNewsClient get news client.
func GetNewsClient(address string) (newsClient NewsClient) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("not connect address: %s , err: %s",address,err))
	}

	newsClient = NewNewsClient(conn)

	return newsClient
}
