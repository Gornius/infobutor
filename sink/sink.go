package sink

import (
	"sync"

	"github.com/gornius/infobutor/message"
	"github.com/gornius/infobutor/sender"
)

// sink is struct that holds its senders and relays messages to its senders
type Sink struct {
	Name    string
	Token   string
	Senders []sender.Sender
}

// sends messages to all senders of a sink
// TODO: implement async error handling
func (ch *Sink) Send(message *message.Message) error {
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
