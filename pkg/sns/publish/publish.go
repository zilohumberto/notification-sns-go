package publish

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// Publish send through sns a notification
func Publish(msgPtr *string, topicPtr *string) (sns.PublishOutput, error) {
	// urlEndpoint is added only for test purpose
	urlEndpoint := "http://localhost:4100"
	// credentials take it from ENVIRONMENT
	creds := credentials.NewEnvCredentials()
	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint: &urlEndpoint,
			Credentials: creds,
		},
	}))

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:  msgPtr,
		TopicArn: topicPtr,
	})
	if err != nil {
		return sns.PublishOutput{}, err
	}
	return *result, nil
}
