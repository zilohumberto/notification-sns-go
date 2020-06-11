package main

import (
	notification "github.com/zilohumberto/notification-sns-go/pkg/http/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := notification.Server{}
	log.Println("Start server")
	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	log.Println("Estamos comandando la liga")
}