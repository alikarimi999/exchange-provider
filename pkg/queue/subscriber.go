package queue

import "sync"

type Subscriber interface {
	Run(wg *sync.WaitGroup)
	Subscribe(t string) (<-chan Message, error)
}

type Message interface {
	Id() string
	Topic() string
	Data() []byte
	Ack() error
}

func NewSubscriber(url string, queue string, ts []*Topic) Subscriber {

	return newRabbitSubscriber(url, queue, ts)
}
