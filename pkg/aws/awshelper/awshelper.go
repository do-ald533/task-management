package awshelper

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Config() *aws.Config {
	var (
		AWS_ACCESS_KEY_ID     = os.Getenv("AWS_ACCESS_KEY_ID")
		AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
		AWS_REGION            = os.Getenv("AWS_REGION")
	)

	cred := credentials.NewStaticCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, "fake-session-token")
	conf := aws.NewConfig().WithRegion(AWS_REGION).WithCredentials(cred).WithEndpoint("http://localhost:4566")

	return conf
}

func NewAwsSession() *session.Session {
	conf := Config()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            *conf,
	}))

	return sess
}
