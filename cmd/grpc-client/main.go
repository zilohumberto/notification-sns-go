package main

import (
	"log"

	"github.com/zilohumberto/notification-sns-go/pkg/http/grpc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	push := notification.Notification{
		Target: "Hello From Client!",
		Message: "Message",
		Subject: "Subject",
	}
	response, err := c.SendNotification(context.Background(), &push)
	if err != nil {
		log.Fatalf("Error when calling SendNotification: %s", err)
	}
	log.Printf("Response from server: %s", response.Target)
}
