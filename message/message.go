package message

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/ytake/protoactor-fizzbuzz/command"
)

// RouteSlip is a message
type RouteSlip struct {
	RouteSlip []*actor.PID
	Message   *command.Say
}
