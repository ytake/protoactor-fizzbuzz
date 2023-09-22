package fizzbuzz

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/ytake/protoactor-fizzbuzz/message"
)

// Fizz is a struct
type Fizz struct {
	*RouteSlip
}

func NewFizz() actor.Actor {
	return &Fizz{RouteSlip: &RouteSlip{}}
}

// Receive is actor
func (f *Fizz) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *message.RouteSlip:
		if msg.Message.Number%3 == 0 {
			msg.Message.Text = "Fizz"
		}
		f.SendMessageToNextTask(context, msg.RouteSlip, msg.Message)
	}
}
