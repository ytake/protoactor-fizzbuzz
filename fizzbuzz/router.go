package fizzbuzz

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/ytake/protoactor-fizzbuzz/command"
	"github.com/ytake/protoactor-fizzbuzz/message"
)

type RouteSlip struct{}

type FizzBuzz struct {
	RouteSlip []*actor.PID
	Message   interface{}
}

// SendMessageToNextTask is a method
func (_ *RouteSlip) SendMessageToNextTask(context actor.Context, routeSlip []*actor.PID, msg *command.Say) {
	nextTask := routeSlip[0]
	newSlip := routeSlip[1:]
	if len(newSlip) == 0 {
		context.Send(nextTask, msg)
		return
	}
	context.Send(nextTask, &message.RouteSlip{
		RouteSlip: newSlip,
		Message:   msg,
	})
}

// SlipRouter is a sample
type SlipRouter struct {
	Fizz    *actor.PID
	Buzz    *actor.PID
	EndStep *actor.PID
	*RouteSlip
}

func (state *SlipRouter) createRouteSlip() []*actor.PID {
	routeSlip := make([]*actor.PID, 0)
	routeSlip = append(routeSlip, state.Fizz, state.Buzz)
	return append(routeSlip, state.EndStep)
}

func (state *SlipRouter) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *command.Say:
		state.SendMessageToNextTask(context, state.createRouteSlip(), msg)
	}
}
