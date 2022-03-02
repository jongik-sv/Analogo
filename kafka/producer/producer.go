package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/Shopify/sarama"
)

// var kafkaClient sarama.AsyncProducer

func main() {

	producer, err := sarama.NewAsyncProducer([]string{"10.110.1.12:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, producerErrors int
ProducerLoop:
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: "c10", Key: nil, Value: sarama.StringEncoder("testing" + strconv.Itoa(enqueued))}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			producerErrors++
		case <-signals:
			break ProducerLoop
		}

		if enqueued > 100 {
			break
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, producerErrors)
}

func testProduce(topic string, limit int) <-chan struct{} {
	var produceDone = make(chan struct{})

	p, err := sarama.NewAsyncProducer([]string{"10.110.1.12:9092"}, sarama.NewConfig())
	if err != nil {
		return nil
	}

	go func() {
		defer close(produceDone)

		for i := 0; i < limit; i++ {
			msg := sarama.StringEncoder("testing" + strconv.Itoa(i))

			if err != nil {
				continue
			}
			select {
			case p.Input() <- &sarama.ProducerMessage{
				Topic: topic,
				Value: msg,
			}:
			case err := <-p.Errors():
				fmt.Printf("Failed to send message to kafka, err: %s, msg: %s\n", err, msg)
			}
		}
	}()

	return produceDone
}
