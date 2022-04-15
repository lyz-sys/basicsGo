package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"

	mconfig "demo/config"

	"github.com/Shopify/sarama"
)

func AsyncProducer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = -1
	config.Net.MaxOpenRequests = 1
	config.Producer.Idempotent = true

	producer, err := sarama.NewAsyncProducer(mconfig.KafkaBrokerList, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                                  sync.WaitGroup
		enqueued, successes, producerErrors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

	var i int
ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: "testTopic1", Value: sarama.StringEncoder("message  " + strconv.Itoa(i))}
		i++
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}

		if i > 100 {
			signals <- os.Interrupt
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, producerErrors)
}

func AsyncProducerSelect() {
	producer, err := sarama.NewAsyncProducer(mconfig.KafkaBrokerList, nil)
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
	signal.Notify(signals, os.Kill)

	var enqueued, producerErrors, i int
ProducerLoop:
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: "testTopic1", Key: nil, Value: sarama.StringEncoder("message  " + strconv.Itoa(i))}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			producerErrors++
		case <-signals:
			break ProducerLoop
		}
		if i > 100 {
			signals <- os.Interrupt
		}
	}
	log.Printf("Enqueued: %d; errors: %d\n", enqueued, producerErrors)
}

func main() {
	//AsyncProducer()

	AsyncProducerSelect()
}
