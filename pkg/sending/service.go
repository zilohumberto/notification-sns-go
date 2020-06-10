package sending

import (
	"github.com/zilohumberto/notification-sns-go/pkg/sns/publish"
	"log"
)

type Service interface{
	SendPush(Notification) error
}

type Repository interface{
	SendPush(Notification) error
}

type service struct{
	r Repository
}

func NewService() Service {
	return &service{}
}

func (s *service) SendPush(n Notification) error{
	result, err := publish.Publish(&n.Message, &n.Topic)
	if err != nil{
		//log.Println(err)
		return err
	}
	log.Print(result)
	return nil
}
