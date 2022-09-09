package main

import (
	"runtime"

	"github.com/Ferza17/event-driven-api-gateway/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Run()
}
