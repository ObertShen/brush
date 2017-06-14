package kafka

import (
	"log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
)

var (
	logger          = log.New(os.Stderr, "[srama]", log.LstdFlags)
	bigdataProducer *Producer
)

const defaultTopic = "bigdata"

// GetProducer 获得 bigdataProducer
func GetProducer() *Producer {
	if bigdataProducer == nil {
		bigdataProducer = NewProducer()
	}

	return bigdataProducer
}

// Producer kafka producer类
type Producer struct {
	conn sarama.SyncProducer
}

// NewProducer 创建 Producer
func NewProducer() *Producer {
	sarama.Logger = logger

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Wait for all in-sync replicas to ack the message
	config.Producer.Retry.Max = 10                   // Retry up to 10 times to produce the message
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(strings.Split("go.sna.com:9092", ","), config)
	if err != nil {
		logger.Printf("Failed to produce message: %s", err.Error())
		logger.Panic(err)
	}

	return &Producer{producer}
}

// SendMessage 向kafaka对应的topic发送信息
func (kp *Producer) SendMessage(topic, message string) (err error) {
	if topic == "" {
		topic = defaultTopic
	}

	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Partition = int32(-1)

	msg.Key = sarama.StringEncoder("key")
	msg.Value = sarama.ByteEncoder(message)

	partition, offset, err := kp.conn.SendMessage(msg)
	if err != nil {
		logger.Println("Failed to produce message: ", err)
		return err
	}

	logger.Printf("partition=%d, offset=%d\n", partition, offset)
	return nil
}
