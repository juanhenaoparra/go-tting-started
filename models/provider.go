package models

import (
	"errors"
)

const (
	// HTTPNotificationProviderType is a type of message provider for HTTP
	HTTPNotificationProviderType NotificationProviderType = "http"
	// EmailNotificationProviderType is a type of message provider for email
	EmailNotificationProviderType NotificationProviderType = "email"
	// SMSNotificationProviderType is a type of message provider for SMS
	SMSNotificationProviderType NotificationProviderType = "sms"
)

var (
	// ValidNotificationProviders is a map of valid notification providers
	ValidNotificationProviders = map[NotificationProviderType]bool{
		HTTPNotificationProviderType:  true,
		EmailNotificationProviderType: true,
		SMSNotificationProviderType:   false,
	}

	// ErrUnknownNotificationProvider is an error for unknown notification provider
	ErrUnknownNotificationProvider = errors.New("unknown notification provider")
)

// NotificationProviderType is a type of message provider
type NotificationProviderType string

// NotificationProvider is a message provider
type NotificationProvider interface {
	Send(message Message) error
}

// NewNotificationProvider creates a new message provider
func NewNotificationProvider(t NotificationProviderType, direction string) (NotificationProvider, error) {
	switch t {
	case HTTPNotificationProviderType:
		return &HTTPNotificationProvider{
			Segment: direction,
		}, nil
	case EmailNotificationProviderType:
		return &EmailNotificationProvider{
			Email: direction,
		}, nil
	default:
		return nil, ErrUnknownNotificationProvider
	}
}
