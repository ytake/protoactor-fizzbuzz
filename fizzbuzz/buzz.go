package fizzbuzz

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/ytake/protoactor-fizzbuzz/message"
)

// Buzz is a struct
type Buzz struct {
	*RouteSlip
}

func NewBuzz() actor.Actor {
	return &Buzz{RouteSlip: &RouteSlip{}}
}

// Receive is actor
func (b *Buzz) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *message.RouteSlip:
		if msg.Message.Number%5 == 0 {
			msg.Message.Text = msg.Message.Text + "Buzz"
		}
		b.SendMessageToNextTask(context, msg.RouteSlip, msg.Message)
	}
}
