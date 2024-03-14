package models

import (
	"fmt"

	"github.com/juanhenaoparra/go-tting-started/shared"
)

// EmailNotificationProvider is a message provider
type EmailNotificationProvider struct {
	Email string
}

// Send sends a message
func (e *EmailNotificationProvider) Send(message Message) error {
	res, err := shared.SendPostJSON(e.Email, message.Payload)
	if err != nil {
		return err
	}

	fmt.Printf("[200] [%s] notification '%s' sent to '%s' with response: %s\n", EmailNotificationProviderType, message.ID, e.Email, res)

	return nil
}
