package notification

import (
	"github.com/zilohumberto/notification-sns-go/pkg/sending"
	"golang.org/x/net/context"
	"log"
)

type Server struct {

}

func (s *Server) SendNotification(ctx context.Context, in *Notification) (*Notification, error){
	log.Printf("\nTarget: %s\n Message: %s\n Subject: %s", in.Target, in.Message, in.Target)
	input := sending.Notification{
		Target: in.Target,
		Message: in.Message,
		Subject: in.Subject,
	}
	service := sending.NewService()
	err := service.SendPush(input)
	if err != nil{
		return &Notification{Target: "Push was not sent"}, err
	}
	return &Notification{Target: "Push sent"}, nil
}