package main

import (
	"context"
	"errors"
	"fmt"
	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"log"
)

func main() {
	// create a Dapr service server
	fmt.Println("start...")
	s := daprd.NewService(":50051")
	//if err != nil {
	//	log.Fatalf("failed to start the server: %v", err)
	//}
	if err := s.AddServiceInvocationHandler("echo", echoHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}
	dc, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	// 保存一个状态
	err = dc.SaveState(ctx, "statestore", "hello", []byte("stop222"))
	fmt.Println("in...", err)
	if err != nil {
		fmt.Println(err)
	}

	// 获取一个状态
	item, err := dc.GetState(ctx, "statestore", "hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(item.Value))
	// start the server
	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func echoHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("invocation parameter required")
		return
	}
	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
