package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"strings"

	"./engine"
)

// printCommand - command for priniting messages to console
type printCommand struct {
	arg string
}

// Execute - run command
func (pc *printCommand) Execute(loop engine.Handler) {
	fmt.Println(pc.arg)
}

// hashCommand - command for hashing string with sha1 and print to console
type hashCommand struct {
	arg string
}

// Execute - run command
func (hc *hashCommand) Execute(loop engine.Handler) {
	h := sha1.New()
	h.Write([]byte(hc.arg))
	bs := h.Sum(nil)
	loop.Post(&printCommand{arg: string(bs)})
}

// parse - parsing messages to commands structs
func parse(commandline string) engine.Command {
	parts := strings.Fields(commandline)
	switch parts[0] {
	case "print":
		return &printCommand{arg: parts[1]}
	case "sha1":
		return &hashCommand{arg: parts[1]}
	default:
		return &printCommand{arg: "Syntax Error Unexpected command"}
	}
}

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open("./commands.txt"); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			eventLoop.Post(parse(scanner.Text()))
		}
	}
	eventLoop.AwaitFinish()
}
