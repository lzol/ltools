package util

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type receive func(d amqp.Delivery)

type Amqp struct {
	URI      string //amqp://guest:guest@localhost:5672/
	Exchange Exchange
	Queue    Queue
	conn     *amqp.Connection
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
	Name       string     //queue name
	BindKey    string     //queue binding key

	Passive    bool       //queue passive
	Durable    bool       //queue durable
	Exclusive  bool       //queue exclusive
	AutoDelete bool       //queue autodelete
	NoWait     bool       //queue nowait
	Arguments  amqp.Table //queue arguments
}

func (a *Amqp) connect() {
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
	log.Println("connect to MQ:", a.URI, " success")
}

func (a *Amqp)Close(){
	a.conn.Close()
}

func (a *Amqp) Consume(f receive) {
	if a.conn == nil {
		a.connect()
	}
	defer a.conn.Close()
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
		nil,                   // arguments
	)
	if err != nil {
		log.Println("Exchange Declare: %s", err)
	}

	queue, err := channel.QueueDeclare(
		a.Queue.Name,       // name of the queue
		a.Queue.Durable,    // durable
		a.Queue.AutoDelete, // delete when unused
		a.Queue.Exclusive,  // exclusive
		a.Queue.NoWait,     // noWait
		nil,                // arguments
	)
	if err != nil {
		log.Println("Queue Declare: %s", err)
	}

	err = channel.QueueBind(
		queue.Name,      // name of the queue
		a.Queue.BindKey, // bindingKey
		a.Exchange.Name, // sourceExchange
		a.Queue.NoWait,  // noWait
		nil,             // arguments
	)
	if err != nil {
		log.Println("Queue Bind: %s", err)
	}

	msgs, err := channel.Consume(
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
	for d := range msgs {
		go f(d)
	}
	<-forever
}

func (a *Amqp) Public(msg string) (err error) {
	if a.conn == nil {
		a.connect()
	}
	channel, err := a.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	if err := channel.ExchangeDeclare(
		a.Exchange.Name, // name
		a.Exchange.Type, // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // noWait
		nil,             // arguments
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
		a.Exchange.Name,   // publish to an exchange
		a.Exchange.RoutingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
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
