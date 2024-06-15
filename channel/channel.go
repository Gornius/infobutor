package channel

import (
	"sync"

	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender"
)

type Channel struct {
	Name    string
	Token   string
	Senders []sender.Sender
}

// TODO: implement async error handling and write tests
func (ch *Channel) Send(message *message.Message) error {
	wg := sync.WaitGroup{}
	for _, s := range ch.Senders {
		s := s
		wg.Add(1)
		go func() {
			s.Send(*message)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
