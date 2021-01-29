package main

import (
	"client/protoapi"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

//根据实际情况填写grpc地址
const address  = "127.0.0.1:9001"

func main() {
	newsClient := protoapi.GetNewsClient(address)

	// GetNewsInfoById调用
	resp01, err := newsClient.GetNewsInfoById(context.Background(),&protoapi.IdReq{Id:22})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp01)

	resp02, err := newsClient.GetNews(context.Background(),&empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp02)

}
