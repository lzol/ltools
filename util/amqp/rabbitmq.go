package util

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

type receive func(d amqp.Delivery)

type Amqp struct {
	URI      string //amqp://guest:guest@localhost:5672/
	Exchange Exchange
	Queue    Queue
	conn     *amqp.Connection
	channel  *amqp.Channel
	ReconnectInternal int64	//reconnect when MQ lost,unit second
}

// direct|fanout|topic|x-custom
const EXCHANGE_TYPE_DIRECT string = "direct"
const EXCHANGE_TYPE_FANOUT string = "fanout"
const EXCHANGE_TYPE_TOPIC_ string = "topic"
const EXCHANGE_TYPE_XCUSTOM string = "x-custom"

type Exchange struct {
	Name       string     //exchange name
	Type       string     //exchange type
	RoutingKey string     //queue routing key
	Passive    bool       //exchange passive
	Durable    bool       //exchange durable
	AutoDelete bool       //exchange autodelete
	Internal   bool       //exchange internal
	NoWait     bool       //exchange nowait
	Arguments  amqp.Table //exchange arguments
}

type Queue struct {
	Name    string //queue name
	BindKey string //queue binding key

	Passive    bool       //queue passive
	Durable    bool       //queue durable
	Exclusive  bool       //queue exclusive
	AutoDelete bool       //queue autodelete
	NoWait     bool       //queue nowait
	Arguments  amqp.Table //queue arguments
}

func (a *Amqp) getChannel() (*amqp.Channel, error) {
	err := a.connect()
	if err != nil {
		return nil, err
	}

	channel, err := a.conn.Channel()
	if err != nil {
		log.Println("open channel fail:", err)
	}

	err = channel.ExchangeDeclare(
		a.Exchange.Name,       // name of the exchange
		a.Exchange.Type,       // type
		a.Exchange.Durable,    // durable
		a.Exchange.AutoDelete, // delete when complete
		a.Exchange.Internal,   // internal
		a.Exchange.NoWait,     // noWait
		a.Exchange.Arguments,              // arguments
	)
	if err != nil {
		log.Println("Exchange Declare: %s", err)
	}

	_, err = channel.QueueDeclare(
		a.Queue.Name,       // name of the queue
		a.Queue.Durable,    // durable
		a.Queue.AutoDelete, // delete when unused
		a.Queue.Exclusive,  // exclusive
		a.Queue.NoWait,     // noWait
		a.Queue.Arguments,                // arguments
	)
	if err != nil {
		log.Println("Queue Declare: %s", err)
	}
	return channel, err
}

func (a *Amqp) connect() (err error) {
	if a.URI == "" {
		log.Println("connect to MQ error:URI is empty")
	}
	if a.Exchange.Name == "" {
		log.Println("connect to MQ error:Exchange Name is empty")
	}
	if a.Queue.Name == "" {
		log.Println("connect to MQ error:Queue Name is empty")
	}
	conn, err := amqp.Dial(a.URI)
	if err != nil {
		log.Println("connect to MQ error:", err)
	}
	a.conn = conn

	if err == nil {
		log.Println("connect to MQ:", a.URI, " success")
	}
	return err
}

func (a *Amqp) Close() {
	a.conn.Close()
}

func (a *Amqp) Consume(f receive) {
RECONNECT:
	for {
		channel, err := a.getChannel()
		if err != nil {
			log.Println("connect to MQ error,reconnect")
			time.Sleep(time.Duration(a.ReconnectInternal)*time.Second)
			continue RECONNECT
		}
		a.channel = channel

		if a.channel != nil {
			msgs, err := a.channel.Consume(
				a.Queue.Name, // name
				"",           // consumerTag,
				true,         // noAck
				false,        // exclusive
				false,        // noLocal
				false,        // noWait
				nil,          // arguments
			)
			if err != nil {
				log.Println("Queue Consume: %s", err)
			}
			forever := make(chan bool)
			for {
				msg, ok := <-msgs
				if !ok {
					continue RECONNECT
				}
				go f(msg)
			}
			<-forever
		}
		defer a.Close()
	}
}

func (a *Amqp) Public(msg string) (err error) {
	channel, err := a.getChannel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	if err := channel.ExchangeDeclare(
		a.Exchange.Name,       // name of the exchange
		a.Exchange.Type,       // type
		a.Exchange.Durable,    // durable
		a.Exchange.AutoDelete, // delete when complete
		a.Exchange.Internal,   // internal
		a.Exchange.NoWait,     // noWait
		a.Exchange.Arguments,              // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	// Reliable publisher confirms require confirm.select support from the
	// connection.
	if err := channel.Confirm(false); err != nil {
		return fmt.Errorf("Channel could not be put into confirm mode: %s", err)
	}

	confirms := channel.NotifyPublish(make(chan amqp.Confirmation, 1))

	defer confirmOne(confirms)

	if err = channel.Publish(
		a.Exchange.Name,       // publish to an exchange
		a.Exchange.RoutingKey, // routing to 0 or more queues
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(msg),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
			// a bunch of application/implementation-specific fields
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}
	defer a.Close()
	return err
}
func confirmOne(confirms <-chan amqp.Confirmation) {
	log.Printf("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
	} else {
		log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
	}
}
