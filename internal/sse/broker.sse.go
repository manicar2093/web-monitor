package sse

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Broker struct {
	Notifier       chan interface{}
	newClient      chan chan interface{}
	closingClients chan chan interface{}
	clients        map[chan interface{}]bool
}

func NewBroker() (broker *Broker) {
	broker = &Broker{
		Notifier:       make(chan interface{}, 1),
		newClient:      make(chan chan interface{}),
		closingClients: make(chan chan interface{}),
		clients:        make(map[chan interface{}]bool),
	}

	go broker.listen()

	return
}
func (broker *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan := make(chan interface{})

	broker.newClient <- messageChan

	defer func() {
		broker.closingClients <- messageChan
	}()

	notify := r.Context().Done()

	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	for {
		fmt.Fprintf(w, "data: %s\n\n", <-messageChan)
		flusher.Flush()
	}
}

func (broker *Broker) listen() {

	for {
		select {
		case s := <-broker.newClient:
			broker.clients[s] = true
			log.Printf("Client added. %d registrated clients", len(broker.clients))
		case s := <-broker.closingClients:
			delete(broker.clients, s)
			log.Printf("Removed client. %d registreted clients", len(broker.clients))
		case event := <-broker.Notifier:
			for clientMessageChan, _ := range broker.clients {
				clientMessageChan <- event
			}
		}
	}
}

func (t Broker) Notify(data interface{}) {
	e, j := json.Marshal(data)
	if e != nil {
		log.Printf("error al parsear respuesta. Detalles: %v\n", e)
		return
	}
	t.Notifier <- j
}
