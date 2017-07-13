package main

import (
	"fmt"
	"strconv"
)

// Command define the type that can execute something
type Command interface {
	execute(string)
}

// HelloCommand that print hello world when executed
type HelloCommand struct {
	name string
}

func (c *HelloCommand) execute(newName string) {
	fmt.Println("Hello world " + c.name)
	c.name = newName
}

// ExecuteCommands execute list of command.
func ExecuteCommands(commands ...Command) {
	for i, command := range commands {
		command.execute(strconv.Itoa(i))
	}
}

func main() {
	var helloAnjanaCmd = new(HelloCommand)
	helloAnjanaCmd.name = "Anjana"

	var helloMichelleCmd = new(HelloCommand)
	helloMichelleCmd.name = "Michelle"

	commands := []Command{
		helloAnjanaCmd,
		helloMichelleCmd}

	ExecuteCommands(commands...)
	ExecuteCommands(commands...)
}
