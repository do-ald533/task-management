package sqs

import (
	"goapi/pkg/aws/awshelper"

	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsConsumer struct {
	service       *sqs.SQS
	queueURL  string
	batchSize int64
	waitTime  int64
	stop      chan bool
}

func NewSqsConsumer(queueURL string, batchSize int64, waitTime int64) *SqsConsumer {
	session := awshelper.NewAwsSession()

	return &SqsConsumer{
		service:       sqs.New(session),
		queueURL:  queueURL,
		batchSize: batchSize,
		waitTime:  waitTime,
		stop:      make(chan bool),
	}
}

func (c *SqsConsumer) GetMessage() (interface{}, error) {
	result, err := c.service.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &c.queueURL,
		MaxNumberOfMessages: &c.batchSize,
		WaitTimeSeconds:     &c.waitTime,
	})

	if err != nil {
		return nil, err
	}

	if len(result.Messages) > 0 {
		return result.Messages[0], nil
	}

	return nil, nil
}

func (c *SqsConsumer) DeleteMessage(receiptHandle *string) error {
	_, err := c.service.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &c.queueURL,
		ReceiptHandle: receiptHandle,
	})

	return err
}

func (c *SqsConsumer) Stop() {
	c.stop <- true
}
