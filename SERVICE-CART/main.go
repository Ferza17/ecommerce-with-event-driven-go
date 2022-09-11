package main

import (
	"runtime"

	"github.com/Ferza17/event-driven-cart-service/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
