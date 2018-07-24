package main

import (
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

	grpcClient := pb.NewTransportClient(conn)

	streamClient := &client.Client{}
	log.Println("Starting client ... ")
	wg := &sync.WaitGroup{}
	wg.Add(9)
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-1")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-2")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-3")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-4")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-5")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-6")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-7")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-8")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	go func() {
		// rpc
		if err := streamClient.Stream(ctx, grpcClient, []byte("client-hello-9")); err != nil {
			log.Println(err)
			wg.Done()
		}
	}()
	wg.Wait()
}
