package models

import (
	"fmt"

	"github.com/juanhenaoparra/go-tting-started/shared"
)

// HTTPNotificationProvider is a message provider
type HTTPNotificationProvider struct {
	Segment string
}

// Send sends a message
func (h *HTTPNotificationProvider) Send(message Message) error {
	res, err := shared.SendPostJSON(h.Segment, message.Payload)
	if err != nil {
		return err
	}

	fmt.Printf("[200] [%s] notification '%s' sent to '%s' with response: %s\n", HTTPNotificationProviderType, message.ID, h.Segment, res)

	return nil
}
