package rabbit

import (
	"io"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/streadway/amqp"
)

const (
	exchangeName = ""
	contentType  = "text/plain"
)

// Producer ...
type Producer interface {
	Publish(msg []byte) error
	Close() error
}

type producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue

	queueName string

	closeFuncs []io.Closer
}

// NewProducer ...
func NewProducer(config *config.RabbitProducer) (Producer, error) {
	closeFuncs := make([]io.Closer, 0, 2)

	conn, err := amqp.Dial(config.DSN)
	if err != nil {
		return nil, err
	}
	closeFuncs = append(closeFuncs, conn)

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	closeFuncs = append(closeFuncs, ch)

	q, err := ch.QueueDeclare(
		config.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &producer{
		connection: conn,
		channel:    ch,
		queue:      &q,
		queueName:  config.QueueName,
		closeFuncs: closeFuncs,
	}, nil
}

// Publish ...
func (p *producer) Publish(msg []byte) error {
	err := p.channel.Publish(
		exchangeName,
		p.queueName,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  contentType,
			Body:         msg,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

// Close ...
func (p *producer) Close() error {
	for _, closeFunc := range p.closeFuncs {
		closeFunc.Close()
	}

	return nil
}
