package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/juanhenaoparra/go-tting-started/app"
	"github.com/juanhenaoparra/go-tting-started/models"
	"github.com/juanhenaoparra/go-tting-started/shared/respond"
)

var (
	defaultPort = 8001
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		fmt.Println("error creating new app: ", err.Error())
		return
	}

	setupRoutes(a.ActiveQueue, a.Router)

	fmt.Println("starting app on port: ", defaultPort)

	err = a.Start(defaultPort)
	if err != nil {
		fmt.Printf("error starting app on port%d: %s", defaultPort, err.Error())
	}

}

func setupRoutes(queue *models.MessageQueue, router *chi.Mux) {
	router.Post("/new", func(w http.ResponseWriter, r *http.Request) {
		payload := make(map[string]any)

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			respond.Err(w, err)
			return
		}

		receiverType := models.NotificationProviderType(r.URL.Query().Get("receiver_type"))
		receiverDirection := r.URL.Query().Get("receiver_direction")

		message, err := queue.Push(payload, receiverType, receiverDirection)
		if err != nil {
			reqErr := respond.NewRequestError(http.StatusBadRequest, err.Error())
			respond.Err(w, reqErr)
			return
		}

		respond.JSON(w, http.StatusOK, message)
	})

	router.Put("/send", func(w http.ResponseWriter, r *http.Request) {
		m, err := queue.Pop()
		if errors.Is(err, models.ErrNoMessagesLeft) {
			respond.Err(w, respond.NewRequestError(http.StatusNotFound, err.Error()))
			return
		}

		if err != nil {
			respond.Err(w, err)
			return
		}

		err = m.Send()
		if err != nil {
			respond.Err(w, err)
			return
		}

		respond.JSON(w, http.StatusOK, m)
	})

	router.Put("/send/bulk", func(w http.ResponseWriter, r *http.Request) {
		sentMessages := make([]models.Message, 0)

		for {
			m, err := queue.Pop()
			if err != nil {
				break
			}

			err = m.Send()
			if err != nil {
				respond.Err(w, err)
				return
			}

			sentMessages = append(sentMessages, *m)
		}

		if len(sentMessages) == 0 {
			respond.Err(w, respond.NewRequestError(http.StatusNotFound, "no messages were sent"))
			return
		}

		respond.JSON(w, http.StatusOK, sentMessages)
	})
}
