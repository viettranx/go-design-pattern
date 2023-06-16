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

type ServiceMemento struct {
	value    Value
	cmdStack *CommandStack
}

func (s *UndoableService) DoAdd() {
	s.addCmd.Execute()
	s.cmdStack.Push(s.addCmd)
}

func (s *UndoableService) DoSub() {
	s.subCmd.Execute()
	s.cmdStack.Push(s.subCmd)
}

func (s *UndoableService) Save() ServiceMemento {
	return ServiceMemento{
		value:    *s.value, // value of a pointer
		cmdStack: &CommandStack{current: s.cmdStack.current},
	}
}

func (s *UndoableService) Restore(m ServiceMemento) {
	*s.value = m.value
	s.cmdStack = m.cmdStack
}

func (s *UndoableService) GetValue() int { return s.value.val }

func (s *UndoableService) Undo() {
	if cmd, err := s.cmdStack.Pop(); err == nil {
		cmd.Undo()
	}
}

type Caretaker struct {
	mementos []ServiceMemento
}

func (c *Caretaker) AddMemento(m ServiceMemento) { c.mementos = append(c.mementos, m) }

func (c *Caretaker) Restore(s *UndoableService, idx int) { s.Restore(c.mementos[idx]) }

func (c *Caretaker) Size() int { return len(c.mementos) }

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
	caretaker := Caretaker{}

	service := NewService(10, 2, 1)
	fmt.Println(service.GetValue()) // 10

	service.DoAdd()
	service.DoAdd()
	fmt.Println(service.GetValue()) // 14

	// Save state
	caretaker.AddMemento(service.Save()) // state at index 0

	service.DoSub()
	fmt.Println(service.GetValue()) // 13

	// Save state
	caretaker.AddMemento(service.Save()) // state at index 1

	service.Undo()
	fmt.Println(service.GetValue()) // 14

	// Save state
	caretaker.AddMemento(service.Save()) // state at index 2

	service.Undo()
	service.Undo()
	fmt.Println(service.GetValue()) // 10

	// Save state
	caretaker.AddMemento(service.Save()) // state at index 3

	service.Undo()
	fmt.Println(service.GetValue()) // 10 - nothing changed

	// Restore state
	for i := 0; i < caretaker.Size(); i++ {
		fmt.Println("Restored state at index:", i)
		caretaker.Restore(&service, i)
		fmt.Println("Value:", service.GetValue())

		service.Undo()
		fmt.Println("Value after undo:", service.GetValue())
	}
}
