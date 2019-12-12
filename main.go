package main

import (
	"https://github.com/KHYehor/architecture-lab4/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
}
