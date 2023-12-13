package chat

import (
	"fmt"
	"math/rand"
	"time"
)

type client struct {
	name       string
	connClosed chan struct{}
}

func NewClient(name string) *client {
	return &client{name: name, connClosed: make(chan struct{})}
}

func (cl *client) enterChat(uuid string, send chan<- *envelop) {
	fmt.Printf("%s joined to chat\n", cl.name)

	go func() {
		defer func() {
			if recover() != nil {
				fmt.Println("error: connection closed")
			}
		}()

		msg := fmt.Sprintf("Hi! My name is %s and I've sent this message at %v", cl.name, time.Now().Format(time.UnixDate))

		send <- &envelop{
			uuid: uuid,
			msg:  msg,
		}

		time.Sleep(time.Duration(1000<<10 + (100 + rand.Int63()%1000)))
	}()

	<-cl.connClosed

	fmt.Printf("[client] %s connection closed\n", cl.name)

}

func (cl *client) onChatClosed() {
	close(cl.connClosed)
}

func (cl *client) onMessage(msg string) {
	fmt.Printf("%s received msg : \"%s\"\n", cl.name, msg)
}
