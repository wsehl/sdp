package main

import "fmt"

type Command interface {
	Execute()
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("light on")
}

func (l *Light) Off() {
	fmt.Println("light off")
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(c Command) {
	i.command = c
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	invoker := &Invoker{}

	light := &Light{}

	lightOnCmd := &LightOnCommand{light: light}
	lightOffCmd := &LightOffCommand{light: light}

	invoker.SetCommand(lightOnCmd)
	invoker.ExecuteCommand()

	invoker.SetCommand(lightOffCmd)
	invoker.ExecuteCommand()
}
