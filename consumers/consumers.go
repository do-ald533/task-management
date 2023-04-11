package consumers

import "time"

type Consumer interface {
	ProcessMessage(interface{}) error
	Consume() error
	Stop()
}

func ConsumeMessages(c Consumer, getMessage func() (interface{}, error), stop chan bool) error {
	for {
		select {
		case <-stop:
			return nil
		default:
			msg, err := getMessage()
			if err != nil {
				return err
			}
			if msg != nil {
				if err := c.ProcessMessage(msg); err != nil {
					return err
				}
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}
}
