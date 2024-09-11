package main

import (
	"fmt"

	"github.com/danielhookx/eventbus"
)

func calculator(a int, b int) {
	fmt.Printf("%d\n", a+b)
}

func main() {
	bus := eventbus.New()
	bus.Subscribe("calculator", calculator)
	bus.Publish("calculator", 10, 20)
}
