package service

import "github.com/google/uuid"

const (
	randomTopicDefault = "random-topic"
)

func RandomTopic() string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return randomTopicDefault
	}
	return newUUID.String()
}
