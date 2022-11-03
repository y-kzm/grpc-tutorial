package main

import (
	"context"
	"log"
	"time"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"com.example/grpc-sample/echo"
)

const addr = "localhost:52000"

func main() {
	argv := os.Args
	// EchoService サーバーへ接続する
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := echo.NewEchoServiceClient(conn)

	// Echo メソッドを呼び出す
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &echo.EchoRequest{Message: argv[1]})
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}
	log.Printf("Received from server: %s", r.GetMessage())
}
