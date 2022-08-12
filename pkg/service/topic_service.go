package service

import (
	"github.com/segmentio/kafka-go"
	"kafka_mate_go/pkg/config"
	"net"
	"strconv"
)

func CreateTopic(topic string, partitionNumber int) error {
	conn, err := kafka.Dial("tcp", config.KafkaAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	controller, err := conn.Controller()
	if err != nil {
		return err
	}
	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()
	topicConfigs := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partitionNumber,
		ReplicationFactor: 1,
	}
	return controllerConn.CreateTopics(topicConfigs)
}

func DeleteTopic(topic string) error {
	conn, err := kafka.Dial("tcp", config.KafkaAddr)
	if err != nil {
		return err
	}
	defer conn.Close()
	controller, err := conn.Controller()
	if err != nil {
		return err
	}
	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()
	return controllerConn.DeleteTopics(topic)
}
