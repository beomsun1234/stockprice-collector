package messagequeue

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type Messagequeue interface {
	SendMessage(msg []byte) error
}

type Kafka struct {
	Client *sarama.Client
}

func NewKafka(client *sarama.Client) Messagequeue {
	return &Kafka{
		Client: client,
	}
}

func (k *Kafka) SendMessage(msg []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := ctx.Err(); err != nil {
		log.Fatalln("kafka msg timeout")
		return err
	}

	producer, err := sarama.NewSyncProducerFromClient(*k.Client)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	sendMessage := k.convertSendMsg(msg)

	_, _, err = producer.SendMessage(sendMessage)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	log.Println("Message sent successfully")
	return nil
}

func (k *Kafka) convertSendMsg(msg []byte) *sarama.ProducerMessage {
	p_msg := &sarama.ProducerMessage{
		Topic: "stock",
		Value: sarama.ByteEncoder(msg),
	}

	return p_msg
}
