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
	messages []Message
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
		messages: make([]Message, 0, 100),
	}
}

// Push adds a new message to the queue
func (mq *MessageQueue) Push(payload map[string]any, receiverType NotificationProviderType, receiverDirection string) (*Message, error) {
	if len(payload) == 0 {
		return nil, ErrEmptyPayload
	}

	if receiverDirection == "" {
		return nil, ErrEmptyReceiverDirection
	}

	isValid := ValidNotificationProviders[receiverType]
	if !isValid {
		return nil, ErrUnknownNotificationProvider
	}

	provider, err := NewNotificationProvider(receiverType, receiverDirection)
	if err != nil {
		return nil, fmt.Errorf("creating '%s' notification provider: %w", receiverType, err)
	}

	m := Message{
		ID:           shared.NewID("msg"),
		Payload:      payload,
		ReceiverType: receiverType,
		handler:      provider,
	}

	mq.messages = append(mq.messages, m)

	return &m, nil
}

// Pop removes the last message from the queue and returns it
func (mq *MessageQueue) Pop() (*Message, error) {
	queueLength := len(mq.messages)

	if queueLength == 0 {
		return nil, ErrNoMessagesLeft
	}

	m := mq.messages[0]
	mq.messages = mq.messages[1:]

	return &m, nil
}

// Send sends a message
func (m *Message) Send() error {
	return m.handler.Send(*m)
}
