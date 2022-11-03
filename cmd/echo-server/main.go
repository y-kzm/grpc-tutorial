package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"com.example/grpc-sample/echo"
)

const port = 52000

func main() {
	// TCP ポートをオープンできるか確認
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC サーバーを生成し、EchoService サーバーの実装を登録する
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &server{})

	// gRPC サーバーを稼働開始
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
