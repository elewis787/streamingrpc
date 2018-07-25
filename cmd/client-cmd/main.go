package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/net/context"

	"github.com/elewis787/streamingrpc/client"
	"github.com/elewis787/streamingrpc/pb"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
		os.Exit(1)
	}()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("127.0.0.1:8787", opts...)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	ts := pb.NewTransportClient(conn)

	cli := &client.Client{
		Transport: ts,
	}

	hc := pb.NewHelloClient(conn)

	helloCli := &client.HelloClient{
		HC: hc,
	}

	// TCP - MULTIPLEXING
	// 1 TCP conenction
	// 6 virtual streams
	//
	// stream 1
	wg := &sync.WaitGroup{}
	wg.Add(6)

	go cli.Stream(ctx, []byte(fmt.Sprintf("client-hello-1")))
	go cli.Stream(ctx, []byte(fmt.Sprintf("client-hello-3")))
	go cli.Stream(ctx, []byte(fmt.Sprintf("client-hello-3")))

	go helloCli.Stream(ctx, []byte(fmt.Sprintf("client-hello-4")))
	go helloCli.Stream(ctx, []byte(fmt.Sprintf("client-hello-5")))
	go helloCli.Stream(ctx, []byte(fmt.Sprintf("client-hello-6")))
	wg.Wait()
}
