package message

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
)

// RouteSlip is a message
type RouteSlip struct {
	RouteSlip []*actor.PID
	Message   *FizzBuzz
}

type FizzBuzz struct {
	Number int64
	Text   string
}

func (f *FizzBuzz) String() string {
	return fmt.Sprintf("%d: %s", f.Number, f.Text)
}
