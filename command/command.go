package command

import "fmt"

type Command interface {
	Execute()
}

// macroCommand ConcreteCommand
type macroCommand struct {
	commands []Command
}

func NewMacroCommand() *macroCommand {
	return &macroCommand{}
}

func (c *macroCommand) Execute() {
	for _, cmd := range c.commands {
		cmd.Execute()
	}
}
func (c *macroCommand) Append(cmd Command) {
	if cmd != c {
		c.commands = append(c.commands, cmd)
	}
}
func (c *macroCommand) Clear() {
	c.commands = nil
}
func (c *macroCommand) Undo() {
	c.commands = c.commands[:len(c.commands)-1]
}

// drawCommand ConcreteCommand
type drawCommand struct {
	drawable drawable
	position *position
}

func NewDrawCommand(drawable drawable, position *position) *drawCommand {
	return &drawCommand{drawable: drawable, position: position}
}

func (c *drawCommand) Execute() {
	c.drawable.draw(c.position.x, c.position.y)
}

// drawable is interface. Receiver.
type drawable interface {
	draw(x, y int)
}

// drawCanvas is struct. Receiver.
type drawCanvas struct {
	history *macroCommand
}

func NewDrawCanvas(history *macroCommand) *drawCanvas {
	return &drawCanvas{history: history}
}

func (d *drawCanvas) draw(x, y int) {
	fmt.Printf("Draw: %d.%d\n", x, y)
}

type position struct {
	x, y int
}

func NewPosition(x int, y int) *position {
	return &position{x: x, y: y}
}
