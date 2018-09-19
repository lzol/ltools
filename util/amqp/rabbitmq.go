package amqp

import (
	"github.com/streadway/amqp"
	"log"
)

type receive func(d amqp.Delivery)

type Amqp struct {
	URI       string //amqp://guest:guest@localhost:5672/
	Exchange  Exchange
	Queue     Queue
	conn *amqp.Connection

}

type Exchange struct {
	Exchange   string     //exchange name
	Type       string     //exchange type
	Passive    bool       //exchange passive
	Durable    bool       //exchange durable
	AutoDelete bool       //exchange autodelete
	Internal   bool       //exchange internal
	NoWait     bool       //exchange nowait
	Arguments  amqp.Table //exchange arguments
}

type Queue struct {
	Queue      string     //queue name
	Passive    bool       //queue passive
	Durable    bool       //queue durable
	Exclusive  bool       //queue exclusive
	AutoDelete bool       //queue autodelete
	NoWait     bool       //queue nowait
	Arguments  amqp.Table //queue arguments
}

func(a *Amqp) connect() {
	if a.URI == ""{
		log.Println("connect to MQ error:URI is empty")
	}
	conn, err := amqp.Dial(a.URI)
	if err!=nil {
		log.Println("connect to MQ error:",err)
	}
	a.conn = conn
	defer a.conn.Close()
}

func (a *Amqp)Consume(f receive,queue Queue){
	if a.conn == nil{
		a.connect()
	}


}

func Public(msg string)(err error){
	return nil
}


