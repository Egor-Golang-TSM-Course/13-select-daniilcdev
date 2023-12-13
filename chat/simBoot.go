package chat

import (
	"fmt"
)

func Chat(done <-chan struct{}) {
	s := NewChatServer(done)

	for i := 0; i < 5; i++ {
		cl := NewClient(fmt.Sprintf("User #%d", i+1))
		s.onClientJoined(cl)
	}

	go s.Start()
}
