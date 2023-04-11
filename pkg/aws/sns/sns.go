package sns

import (
	"encoding/json"
	"goapi/pkg/aws/awshelper"
	"goapi/pkg/aws/sns/snsactions"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSService interface {
	Publish(data interface{}, action snsactions.SNS_ACTIONS) error
	Setup(topicARN string, options ...func(*sns.PublishInput) error) SNSService
}

type SNS struct {
	client   *sns.SNS
	topicARN string
	options  []func(*sns.PublishInput) error
}

func NewSNSService() SNSService {
	session := awshelper.NewAwsSession()

	return &SNS{
		client:   sns.New(session),
		topicARN: "",
	}
}

func (s *SNS) Setup(topicARN string, options ...func(*sns.PublishInput) error) SNSService {
	s.topicARN = topicARN
	s.options = options
	return s
}

func (s *SNS) Publish(data interface{}, action snsactions.SNS_ACTIONS) error {
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	input := &sns.PublishInput{
		Message:  aws.String(string(dataJSON)),
		TopicArn: aws.String(s.topicARN),
	}

	input.MessageAttributes = make(map[string]*sns.MessageAttributeValue)

	input.MessageAttributes["action"] = &sns.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(string(action)),
	}

	for _, option := range s.options {
		if err := option(input); err != nil {
			return err
		}
	}

	_, err = s.client.Publish(input)
	return err
}

func WithAttributes(attributes map[string]string) func(*sns.PublishInput) error {
	return func(input *sns.PublishInput) error {
		if input.MessageAttributes == nil {
			input.MessageAttributes = make(map[string]*sns.MessageAttributeValue)
		}

		for k, v := range attributes {
			input.MessageAttributes[k] = &sns.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String(v),
			}
		}
		return nil
	}
}
