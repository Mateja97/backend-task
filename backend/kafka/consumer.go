package kafka

import (
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type KafkaConsumer struct {
	topic    string
	consumer sarama.Consumer
	Wg       sync.WaitGroup
}

func (kC *KafkaConsumer) Init(kafkaBrokers []string, kafkaTopic string) error {
	kC.topic = kafkaTopic
	var err error
	kC.consumer, err = sarama.NewConsumer(kafkaBrokers, sarama.NewConfig())
	if err != nil {
		return err
	}
	return nil
}

func (kC *KafkaConsumer) ConsumeMessage(ch chan *sarama.ConsumerMessage) {
	partitions, err := kC.consumer.Partitions(kC.topic)
	if err != nil {
		log.Println("[ERROR] Partitions list failed")
		return
	}
	for _, p := range partitions {
		pc, err := kC.consumer.ConsumePartition(kC.topic, int32(p), sarama.OffsetOldest)
		if err != nil {
			log.Printf("[ERROR] Consume partition %d failed", p)
			return
		}

		kC.Wg.Add(1)
		defer pc.Close()
		go func(pc sarama.PartitionConsumer) {
			defer kC.Wg.Done()
			for {
				msg := <-pc.Messages()
				ch <- msg
			}

		}(pc)
		kC.Wg.Wait()
	}
}

func (kC *KafkaConsumer) Stop() error {
	log.Println("Subscriber stopped")
	if err := kC.consumer.Close(); err != nil {
		log.Println("[ERROR] Consumer close failed")
		return err
	}
	return nil
}