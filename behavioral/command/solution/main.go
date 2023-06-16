package main

import (
	"errors"
	"fmt"
)

type Command interface {
	Execute()
	Undo()
}

type CommandAdd struct {
	v     *Value
	param int
}

type CommandSub struct {
	v     *Value
	param int
}

func (cmd CommandAdd) Execute() { cmd.v.Add(cmd.param) }
func (cmd CommandAdd) Undo()    { cmd.v.Sub(cmd.param) }

func (cmd CommandSub) Execute() { cmd.v.Sub(cmd.param) }
func (cmd CommandSub) Undo()    { cmd.v.Add(cmd.param) }

type Value struct {
	val int
}

func (v *Value) Add(n int) { v.val += n }
func (v *Value) Sub(n int) { v.val -= n }
func (v *Value) Val() int  { return v.val }

func NewValue(v int) Value {
	return Value{v}
}

type CommandNode struct {
	cmd  Command
	next *CommandNode
}

type CommandStack struct {
	current *CommandNode
}

func (cmdStack *CommandStack) Push(cmd Command) {
	cmdStack.current = &CommandNode{cmd: cmd, next: cmdStack.current}
}

func (cmdStack *CommandStack) Pop() (Command, error) {
	if cmdNode := cmdStack.current; cmdNode != nil {
		cmdStack.current = cmdNode.next
		return cmdNode.cmd, nil
	}

	return nil, errors.New("no command in stack")
}

type UndoableService struct {
	value    *Value
	addCmd   Command
	subCmd   Command
	cmdStack *CommandStack
}

func (s UndoableService) DoAdd() {
	s.addCmd.Execute()
	s.cmdStack.Push(s.addCmd)
}

func (s UndoableService) DoSub() {
	s.subCmd.Execute()
	s.cmdStack.Push(s.subCmd)
}

func (s UndoableService) GetValue() int { return s.value.val }

func (s UndoableService) Undo() {
	if cmd, err := s.cmdStack.Pop(); err == nil {
		cmd.Undo()
	}
}

func NewService(initValue int, incrStep int, decrStep int) UndoableService {
	value := NewValue(initValue)

	addCmd := CommandAdd{v: &value, param: incrStep}
	subCmd := CommandSub{v: &value, param: decrStep}

	return UndoableService{
		value:    &value,
		addCmd:   addCmd,
		subCmd:   subCmd,
		cmdStack: &CommandStack{},
	}
}

func main() {
	service := NewService(10, 2, 1)
	fmt.Println(service.GetValue()) // 10

	service.DoAdd()
	service.DoAdd()
	fmt.Println(service.GetValue()) // 14

	service.DoSub()
	fmt.Println(service.GetValue()) // 13

	service.Undo()
	fmt.Println(service.GetValue()) // 14

	service.Undo()
	service.Undo()
	fmt.Println(service.GetValue()) // 10

	service.Undo()
	fmt.Println(service.GetValue()) // 10 - nothing changed
}
