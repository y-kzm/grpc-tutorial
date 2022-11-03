package main

import (
	"context"
	"log"

	"com.example/grpc-sample/echo"
)

// EchoService を実装するサーバーの構造体
type server struct {
	echo.UnimplementedEchoServiceServer // とりあえず Not implemented の実装を入れておく
}

// EchoServiceServer インタフェースの Echo メソッドの実装（本物の Echo 実装）
func (s *server) Echo(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	log.Printf("Received from client: %v", in.GetMessage())
	return &echo.EchoResponse{Message: "Hello! " + in.GetMessage()}, nil
}
