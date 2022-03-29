package main

import (
	"context"
	"fmt"
	"log"

	mconfig "test-demo/config"

	"github.com/Shopify/sarama"
)

type consumerGroupHandler struct {
}

func (consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}
func (consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s, partition = %d", msg.Value, msg.Timestamp, msg.Topic, msg.Partition)
		session.MarkMessage(msg, "")
	}
	return nil
}

func SaramaConsumerGroup() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.MaxVersion                    // specify appropriate version
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // 未找到组消费位移的时候从哪边开始消费

	group, err := sarama.NewConsumerGroup(mconfig.KafkaBrokerList, "my-group1", config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	// Track errors
	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()
	fmt.Println("Consumed start")
	// Iterate over consumer sessions.
	ctx := context.Background()
	for {
		topics := []string{"testTopic1"}
		handler := consumerGroupHandler{}

		// `Consume` should be called inside an infinite loop, when a
		// server-side rebalance happens, the consumer session will need to be
		// recreated to get the new claims
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	SaramaConsumerGroup()
}
