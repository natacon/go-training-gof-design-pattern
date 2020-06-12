package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type node interface {
	Parse(context *context) error
	ToString() string
}

type programNode struct {
	commandListNode node
}

func NewProgramNode() *programNode {
	return &programNode{}
}

func (n *programNode) ToString() string {
	return fmt.Sprintf("[program %s]\n", n.commandListNode.ToString())
}

func (n *programNode) Parse(context *context) error {
	_ = context.skipToken("program")
	n.commandListNode = NewCommandListNode()
	err := n.commandListNode.Parse(context)
	return err
}

type CommandListNode struct {
	list []node
}

func (n *CommandListNode) ToString() string {
	var str string
	for i, node := range n.list {
		if i == len(n.list)-1 {
			str += node.ToString()
		} else {
			str += node.ToString() + " "
		}
	}
	return str
}

func (n *CommandListNode) Parse(context *context) error {
	for {
		if context.currentToken == "" {
			return fmt.Errorf("eissing 'end'")
		} else if context.currentToken == "end" {
			err := context.skipToken("end")
			if err != nil {
				return err
			}
			break
		} else {
			node := newCommandNode()
			err := node.Parse(context)
			if err != nil {
				return err
			}
			n.list = append(n.list, node)
		}
	}
	return nil
}

func NewCommandListNode() *CommandListNode {
	return &CommandListNode{}
}

type commandNode struct {
	node node
}

func (n *commandNode) ToString() string {
	return n.node.ToString()
}

func (n *commandNode) Parse(context *context) error {
	if context.currentToken == "repeat" {
		n.node = newRepeatCommandNode()
	} else {
		n.node = NewPrimitiveCommandNode()
	}
	err := n.node.Parse(context)
	return err
}

func newCommandNode() *commandNode {
	return &commandNode{}
}

type repeatCommandNode struct {
	number          int
	commandListNode node
}

func (n *repeatCommandNode) Parse(context *context) error {
	err := context.skipToken("repeat")
	n.number, err = context.currentNumber()
	context.nextToken()
	n.commandListNode = NewCommandListNode()
	err = n.commandListNode.Parse(context)
	return err
}

func (n *repeatCommandNode) ToString() string {
	return fmt.Sprintf("[repeat %d %s]", n.number, n.commandListNode.ToString())
}

func newRepeatCommandNode() *repeatCommandNode {
	return &repeatCommandNode{}
}

type PrimitiveCommandNode struct {
	name string
}

func (n *PrimitiveCommandNode) Parse(context *context) error {
	n.name = context.currentToken
	err := context.skipToken(n.name)
	if n.name != "go" && n.name != "right" && n.name != "left" {
		err = fmt.Errorf("%s is undefined", n.name)
	}
	return err
}

func (n *PrimitiveCommandNode) ToString() string {
	return n.name
}

func NewPrimitiveCommandNode() *PrimitiveCommandNode {
	return &PrimitiveCommandNode{}
}

type context struct {
	currentToken string
	tokens       []string
}

func NewContext(text string) *context {
	context := &context{
		tokens: strings.Fields(text),
	}
	context.nextToken()
	return context
}

func (c *context) skipToken(token string) (err error) {
	if token != c.currentToken {
		err = fmt.Errorf("warning: %s is expected, but %s is found", token, c.currentToken)
	}
	c.nextToken()
	return
}

func (c *context) nextToken() {
	if len(c.tokens) == 0 {
		c.currentToken = ""
	} else {
		// 最初のコマンドをcurrentとし、読み込み済みのコマンドは消すことで読む位置が進んでいることを表す。
		c.currentToken = c.tokens[0]
		c.tokens = c.tokens[1:]
	}
}

func (c *context) currentNumber() (int, error) {
	return strconv.Atoi(c.currentToken)
}
