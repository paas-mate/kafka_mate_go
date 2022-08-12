package service

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"kafka_mate_go/pkg/config"
	"time"
)

func HealthCheck() error {
	topic := RandomTopic()
	err := CreateTopic(topic, 1)
	if err != nil {
		return err
	}
	logrus.Infof("create topic %s success", topic)
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{config.KafkaAddr},
		Topic:     topic,
		GroupID:   topic,
		Partition: 0,
	})
	reader.SetOffset(0)
	defer reader.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := kafka.DialLeader(ctx, "tcp", config.KafkaAddr, topic, 0)
	if err != nil {
		return err
	}
	logrus.Infof("Create conn %s success ", topic)
	defer conn.Close()
	err = conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		return err
	}
	_, err = conn.Write([]byte(topic))
	if err != nil {
		return err
	}
	logrus.Info("producer write message success")
	consumerCtx, consumerCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer consumerCancel()
	message, err := reader.ReadMessage(consumerCtx)
	if err != nil {
		return err
	}
	if string(message.Value) != topic {
		return errors.New("consume msg not equal")
	}
	err = DeleteTopic(topic)
	if err != nil {
		return err
	}
	return nil
}
