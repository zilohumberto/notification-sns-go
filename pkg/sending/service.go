package sending

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/zilohumberto/notification-sns-go/pkg/sns/publish"
	"log"
)

type Service interface{
	SendPush(Notification) error
	SendTopic(Notification) error
}

type service struct{
	msgStructure string
}

func NewService() Service {
	return &service{ "string"}
}

// SendPush load the input to send a SNS notification
func (s *service) SendPush(n Notification) error{
	input := sns.PublishInput{
		Message:  &n.Message,
		TargetArn: &n.Target,
		MessageStructure: &s.msgStructure,
		Subject: &n.Subject,
	}
	result, err := publish.Publish(input)
	if err != nil{
		return err
	}
	log.Print(result)
	return nil
}

// SendTopic load the input to send a SNS notification
func (s *service) SendTopic(n Notification) error{
	input := sns.PublishInput{
		Message:  &n.Message,
		TopicArn: &n.Topic,
		MessageStructure: &s.msgStructure,
		Subject: &n.Subject,
	}
	result, err := publish.Publish(input)
	if err != nil{
		return err
	}
	log.Print(result)
	return nil
}
