package event

import (
	"encoding/json"
	"fmt"
	"order_service/internal/app"

	"order_service/internal/delivery/event/dto"
	"order_service/pkg/logger"
	"order_service/pkg/queue"
	"sync"

	"order_service/pkg/errors"
)

type Consumer struct {
	app *app.OrderUseCase
	sub queue.Subscriber
	ts  []string
	ms  chan queue.Message
	er  chan error

	l logger.Logger
}

func NewConsumer(app *app.OrderUseCase, sub queue.Subscriber, l logger.Logger, topics []string) *Consumer {
	return &Consumer{
		app: app,
		sub: sub,
		ts:  topics,
		ms:  make(chan queue.Message),
		er:  make(chan error),

		l: l,
	}
}

func (s *Consumer) Run(wg *sync.WaitGroup) {
	const op = errors.Op("event.Consumer.Run")

	defer wg.Done()
	w := &sync.WaitGroup{}

	for _, t := range s.ts {
		w.Add(1)
		go func(topic string) {

			defer w.Done()
			msgs, err := s.sub.Subscribe(topic)
			if err != nil {
				s.er <- errors.Wrap(op, err)
			}
			s.l.Debug(string(op), fmt.Sprintf("subscribe topic: %s", topic))

			for msg := range msgs {
				s.ms <- msg
			}
		}(t)
	}

	w.Add(1)
	go s.parser(w)

	s.l.Debug(string(op), "started")
	w.Wait()

}

func (s *Consumer) parser(wg *sync.WaitGroup) {
	const op = errors.Op("event.Consumer.parser")

	s.l.Debug(string(op), "started")
	defer wg.Done()
	for {
		select {
		case msg := <-s.ms:
			switch msg.Topic() {
			case "deposite.confirmed":
				go func(m queue.Message) {
					d := &dto.Deposit{}
					if err := json.Unmarshal(m.Data(), d); err != nil {
						s.l.Error(string(op), errors.Wrap(err, op).Error())
						s.ackMsg(m)
						return
					}

					s.l.Debug(string(op), d.String())

					if err := s.handleConfirmedDeposite(d); err != nil {

						// if NotFound error return, means that order is not exists or already processed
						// so we can ack message
						if errors.ErrorCode(err) == errors.ErrNotFound {
							s.ackMsg(m)
						}
						return
					}

					s.ackMsg(m)

				}(msg)
			}
		case err := <-s.er:
			s.handleError(errors.Wrap(err, op))
		}
	}

}

func (s *Consumer) handleError(err error) {
	s.l.Error("event.Consumer", err.Error())
}

func (s *Consumer) handleConfirmedDeposite(d *dto.Deposit) error {
	if err := s.app.SetDepositeVolume(d.UserID, d.OrderId, d.DepositeId, d.Volume); err != nil {
		return err
	}
	return nil
}

func (s *Consumer) ackMsg(m queue.Message) {
	const op = errors.Op("event.Consumer.ackMsg")
	if err := m.Ack(); err != nil {
		s.l.Error(string(op), errors.Wrap(err, fmt.Sprintf("msgId: '%s'", m.Id())).Error())
		return
	}
	s.l.Debug(string(op), fmt.Sprintf("messageId: '%s' acknowleged", m.Id()))
	return
}
