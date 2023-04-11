package consumers

import (
	"goapi/consumers"
	"goapi/consumers/bookscreatedconsumers"
	"sync"
)

type ConsumerManager struct {
	consumers []consumers.Consumer
	wg        sync.WaitGroup
}

func NewConsumerManager() *ConsumerManager {
	return &ConsumerManager{
		consumers: make([]consumers.Consumer, 0),
	}
}
func (cm *ConsumerManager) AddConsumer(consumer consumers.Consumer) {
	cm.consumers = append(cm.consumers, consumer)
}

func (cm *ConsumerManager) RegisterConsumers() {
	booksConsumer := bookscreatedconsumers.NewBooksCreatedConsumer()
	cm.AddConsumer(booksConsumer)

	// Adicione outros consumers aqui
}

func (cm *ConsumerManager) StartAll() {
	for _, consumer := range cm.consumers {
		cm.wg.Add(1)
		go func(consumer consumers.Consumer) {
			consumer.Consume()
			defer cm.wg.Done()
		}(consumer)
	}
}

func (cm *ConsumerManager) StopAll() {
	for _, consumer := range cm.consumers {
		consumer.Stop()
	}
	cm.wg.Wait()
}
