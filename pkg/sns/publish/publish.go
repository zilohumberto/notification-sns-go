package publish

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"fmt"
)
func Publish(msgPtr *string, topicPtr *string) (sns.PublishOutput, error) {
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Message:  msgPtr,
		TopicArn: topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		return sns.PublishOutput{}, err
	}
	return *result, nil
}
