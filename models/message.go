package models

import (
	"errors"
	"fmt"

	"github.com/juanhenaoparra/go-tting-started/shared"
)

var (
	// ErrEmptyPayload is an error for empty payload
	ErrEmptyPayload = errors.New("empty payload")
	// ErrNoMessagesLeft is an error for no messages left
	ErrNoMessagesLeft = errors.New("no messages left")
	// ErrEmptyReceiverDirection is an error for empty receiver direction
	ErrEmptyReceiverDirection = errors.New("empty receiver direction")
)

// MessageQueue is a simple message queue
type MessageQueue struct {
	Messages []Message
}

// Message is a simple message
type Message struct {
	ID           string                   `json:"id"`
	Payload      map[string]any           `json:"payload"`
	ReceiverType NotificationProviderType `json:"receiver_type"`
	handler      NotificationProvider     `json:"-"`
}

// NewMessageQueue creates a new message queue
func NewMessageQueue() *MessageQueue {
	return &MessageQueue{
		Messages: make([]Message, 0, 100),
	}
}

// Push adds a new message to the queue
func (mq *MessageQueue) Push(payload map[string]any, receiverType NotificationProviderType, receiverDirection string) error {
	if len(payload) == 0 {
		return ErrEmptyPayload
	}

	if receiverDirection == "" {
		return ErrEmptyReceiverDirection
	}

	isValid := ValidNotificationProviders[receiverType]
	if !isValid {
		return ErrUnknownNotificationProvider
	}

	provider, err := NewNotificationProvider(receiverType, receiverDirection)
	if err != nil {
		return fmt.Errorf("creating '%s' notification provider: %w", receiverType, err)
	}

	m := Message{
		ID:           shared.NewID("msg"),
		Payload:      payload,
		ReceiverType: receiverType,
		handler:      provider,
	}

	mq.Messages = append(mq.Messages, m)

	return nil
}

// Pop removes the last message from the queue and returns it
func (mq *MessageQueue) Pop() (*Message, error) {
	queueLength := len(mq.Messages)

	if queueLength == 0 {
		return nil, ErrNoMessagesLeft
	}

	m := mq.Messages[queueLength-1]
	mq.Messages = mq.Messages[0 : queueLength-1]

	return &m, nil
}
