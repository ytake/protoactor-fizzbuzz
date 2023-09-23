package main

import (
	"fmt"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/stream"
	"github.com/ytake/protoactor-fizzbuzz/command"
	"github.com/ytake/protoactor-fizzbuzz/fizzbuzz"
	"github.com/ytake/protoactor-fizzbuzz/message"
)

const rangeTo = 100

func main() {
	system := actor.NewActorSystem()
	p := stream.NewTypedStream[*message.FizzBuzz](system)
	go func() {
		router := system.Root.Spawn(actor.PropsFromProducer(func() actor.Actor {
			return &fizzbuzz.SlipRouter{
				Fizz:      system.Root.Spawn(actor.PropsFromProducer(fizzbuzz.NewFizz)),
				Buzz:      system.Root.Spawn(actor.PropsFromProducer(fizzbuzz.NewBuzz)),
				RouteSlip: &fizzbuzz.RouteSlip{},
				EndStep:   p.PID()}
		}))
		for v := range [rangeTo]int{} {
			system.Root.Send(router, &command.Say{Number: int64(v + 1)})
		}
	}()
	for range [rangeTo]int{} {
		fmt.Println(<-p.C())
	}
}
