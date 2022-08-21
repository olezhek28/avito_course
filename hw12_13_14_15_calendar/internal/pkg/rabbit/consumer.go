package rabbit

import (
	"io"

	"github.com/olezhek28/avito_course/hw12_13_14_15_calendar/internal/config"
	"github.com/streadway/amqp"
)

const consumerName = ""

// Consumer ...
type Consumer interface {
	Consume() (<-chan amqp.Delivery, error)
	Close() error
}

type consumer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue

	queueName string

	closeFuncs []io.Closer
}

// NewConsumer ...
func NewConsumer(config *config.RabbitConsumer) (Consumer, error) {
	var closeFuncs []io.Closer

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

	err = ch.Qos(
		1,
		0,
		false,
	)
	if err != nil {
		return nil, err
	}

	return &consumer{
		connection: conn,
		channel:    ch,
		queue:      &q,
		queueName:  config.QueueName,
	}, nil
}

// Consume ...
func (c *consumer) Consume() (<-chan amqp.Delivery, error) {
	msgChan, err := c.channel.Consume(
		c.queueName,
		consumerName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return msgChan, nil
}

// Close ...
func (c *consumer) Close() error {
	for _, closeFunc := range c.closeFuncs {
		closeFunc.Close()
	}

	return nil
}
