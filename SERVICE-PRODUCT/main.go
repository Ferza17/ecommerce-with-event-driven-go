package main

import (
	"runtime"

	"github.com/Ferza17/event-driven-product-service/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
