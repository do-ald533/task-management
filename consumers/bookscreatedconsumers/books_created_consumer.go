package bookscreatedconsumers

import (
	"goapi/consumers"
	sqsconsumer "goapi/pkg/aws/sqs"
	"goapi/pkg/logging"
	"os"
)

type BooksCreatedConsumer struct {
	sqsConsumer *sqsconsumer.SqsConsumer
	stop        chan bool
}

func NewBooksCreatedConsumer() consumers.Consumer {
	var (
		queueURL        = os.Getenv("SQS_CREATED_BOOKS")
		batchSize int64 = 1
		waitTime  int64 = 20
	)

	baseConsumer := sqsconsumer.NewSqsConsumer(queueURL, batchSize, waitTime)
	return &BooksCreatedConsumer{
		sqsConsumer: baseConsumer,
		stop:        make(chan bool),
	}
}

func (c *BooksCreatedConsumer) Consume() error {
	logging.Info("Starting BooksCreatedConsumer")
	getMessage := func() (interface{}, error) {
		return c.sqsConsumer.GetMessage()
	}

	return consumers.ConsumeMessages(c, getMessage, c.stop)
}

func (c *BooksCreatedConsumer) Stop() {
	logging.Info("Stopping BooksCreatedConsumer")
	c.stop <- true
}
