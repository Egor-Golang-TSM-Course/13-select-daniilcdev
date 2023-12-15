package chat

import (
	"crypto/rand"
	"fmt"
)

type server struct {
	users map[string]*client
	done  <-chan struct{}
	send  chan *envelop
}

func NewChatServer(done <-chan struct{}) *server {
	return &server{
		users: make(map[string]*client),
		done:  done,
		send:  make(chan *envelop),
	}
}

func (s *server) Start() {
	fmt.Println("starting chat server...")

	for {
		select {
		case <-s.done:
			for _, v := range s.users {
				v.onChatClosed()
				fmt.Printf("[server] user disconnected %s\n", v.name)
			}

			fmt.Printf("[server] closing send channel\n")

			close(s.send)
			clear(s.users)

			fmt.Println("chat server stopped")

			return
		case env := <-s.send:
			uname := s.users[env.uuid].name
			fmt.Printf("[server] message received from user (%s), dispatching...\n", uname)

			for k, v := range s.users {
				if k == env.uuid {
					continue
				}

				v.onMessage(env.msg)
			}
		}
	}
}

func (s *server) onClientJoined(cl *client) {
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4],
		b[4:6],
		b[6:8],
		b[8:10],
		b[10:],
	)

	fmt.Printf("client (id:%s, name:%s) entering chat\n", uuid, cl.name)

	s.users[uuid] = cl
	go cl.enterChat(uuid, s.send)
}
