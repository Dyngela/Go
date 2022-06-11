package main

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"os"
	"clientmq/events"
	"time"
)

func main()  {

	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to rabbitmq !")
	log.Println("Listening for message")

	consumer, err := events.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp091.Connection, error)  {
	var count int64
	var backOff = 1 * time.Second

	for {
		c, err := amqp091.Dial("amqp://guest:guest@rabbitMQ")
		if err != nil {
			log.Println("Rabbitmq isn't ready yet")
			count++
		} else {
			return c, nil
		}

		if count > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		log.Println("Backing off from trying to connect to rabbitmq")
		time.Sleep(backOff)
		continue
	}
}
