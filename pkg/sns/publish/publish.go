package publish

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// Publish send through sns a notification
func Publish(input sns.PublishInput) (*sns.PublishOutput, error) {
	// urlEndpoint is added only for test purpose
	urlEndpoint := "http://127.0.0.1:4575"
	// credentials take it from ENVIRONMENT
	_credentials := credentials.NewEnvCredentials()
	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Endpoint: &urlEndpoint,
			Credentials: _credentials,
		},
	}))

	svc := sns.New(sess)
	result, err := svc.Publish(&input)
	if err != nil {
		return nil, err
	}
	return result, nil
}
