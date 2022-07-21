package queue

import (
	"errors"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabMsg struct {
	d     amqp.Delivery
	topic string
	data  []byte
}

func (m *rabMsg) Id() string {
	return m.d.MessageId
}

func (m *rabMsg) Topic() string {
	return m.topic
}

func (m *rabMsg) Data() []byte {
	return m.data
}

func (m *rabMsg) Ack() error {
	return m.d.Ack(false)
}

type rabbitSubscriber struct {
	ch       *amqp.Channel
	queue    string
	topics   []*Topic
	delivery map[string]chan Message
}

func newRabbitSubscriber(url, q string, topics []*Topic) *rabbitSubscriber {

	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	r := &rabbitSubscriber{
		ch:       ch,
		queue:    q,
		topics:   topics,
		delivery: make(map[string]chan Message),
	}

	for _, t := range topics {
		r.delivery[t.RoutingKey] = make(chan Message)
	}

	return r
}

func (s *rabbitSubscriber) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for _, t := range s.topics {
		err := s.ch.ExchangeDeclare(t.Exchange, "topic", true, false, false, false, nil)
		if err != nil {
			panic(err)
		}
		_, err = s.ch.QueueDeclare(s.queue, true, false, false, false, nil)
		if err != nil {
			panic(err)
		}
		err = s.ch.QueueBind(s.queue, t.RoutingKey, t.Exchange, false, nil)
		if err != nil {
			panic(err)
		}
	}

	msgs, err := s.ch.Consume(s.queue, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		select {
		case s.delivery[msg.RoutingKey] <- &rabMsg{
			d:     msg,
			topic: msg.RoutingKey,
			data:  msg.Body,
		}:
		default:
		}
	}

}

func (s *rabbitSubscriber) Subscribe(t string) (<-chan Message, error) {
	if _, ok := s.delivery[t]; !ok {
		return nil, errors.New("topic not found")
	}

	return s.delivery[t], nil
}
