package bookscreatedconsumers

import (
	"goapi/pkg/logging"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func (c *BooksCreatedConsumer) ProcessMessage(msg interface{}) error {
	message := msg.(*sqs.Message)

	logging.Info("Processing SQS Message:", *message.MessageId)
	logging.Info("SQS Body:", *message.Body)

	// A lógica de processamento da mensagem relacionada ao pedido vai aqui
	// ...

	// Após processar a mensagem, delegue a remoção da mensagem ao consumer base
	return c.sqsConsumer.DeleteMessage(message.ReceiptHandle)
}
