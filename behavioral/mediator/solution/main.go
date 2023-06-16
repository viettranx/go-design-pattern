package main

import (
	"fmt"
	"sync"
)

type Tool struct {
	isUsing bool
}

func (t *Tool) CanUse() bool { return !t.isUsing }
func (t *Tool) Using()       { t.isUsing = true }
func (t *Tool) Done()        { t.isUsing = false }

type Component interface {
	AllowUsingTool(tool *Tool)
	Solve()
}

type Worker struct {
	mediator Mediator
}
type Engineer struct {
	mediator Mediator
}

func (w Worker) doJob(tool *Tool) {
	fmt.Println("Worker is using tool")
	w.mediator.NotifyDone(w)
}

func (w Worker) Solve() {
	fmt.Println("Worker needs to use tool to solve problem. He/She is asking mediator...")
	w.mediator.RegisterUsingTool(w)
}

func (w Worker) AllowUsingTool(tool *Tool) {
	go w.doJob(tool)
}

func (eng Engineer) design(tool *Tool) {
	fmt.Println("Engineer is using tool")
	eng.mediator.NotifyDone(eng)
}

func (eng Engineer) Solve() {
	fmt.Println("Engineer needs to use tool to solve problem. He/She is asking mediator...")
	eng.mediator.RegisterUsingTool(eng)
}

func (eng Engineer) AllowUsingTool(tool *Tool) {
	go eng.design(tool)
}

type Mediator interface {
	RegisterUsingTool(c Component)
	NotifyDone(fromComp Component)
	Wait()
}

type SimpleMediator struct {
	tool      *Tool
	queue     []Component
	lock      sync.Mutex
	waitGroup sync.WaitGroup
}

func (m *SimpleMediator) allocToolForComponent(tool *Tool, c Component) {
	tool.Using()
	c.AllowUsingTool(m.tool)
}

func (m *SimpleMediator) RegisterUsingTool(c Component) {
	m.waitGroup.Add(1)

	m.lock.Lock()
	defer m.lock.Unlock()

	if m.tool.CanUse() {
		// No one using it, you are allowed to use it
		m.allocToolForComponent(m.tool, c)
		return
	}

	// add component into queue
	m.queue = append(m.queue, c)
}

func (m *SimpleMediator) NotifyDone(fromComp Component) {
	m.waitGroup.Done()

	m.lock.Lock()
	defer m.lock.Unlock()

	fmt.Println("Component done its job")

	m.tool.Done()

	if len(m.queue) == 0 {
		return
	}

	nextComp := m.queue[0] // enqueue
	m.queue = append(make([]Component, 0), m.queue[1:]...)

	m.allocToolForComponent(m.tool, nextComp)
}

func (m *SimpleMediator) Wait() {
	m.waitGroup.Wait()
}

func NewMediator(tool *Tool) Mediator {
	return &SimpleMediator{
		tool:      tool,
		lock:      sync.Mutex{},
		waitGroup: sync.WaitGroup{},
	}
}

func main() {
	shareTool := &Tool{isUsing: false}
	mediator := &SimpleMediator{tool: shareTool}

	components := []Component{
		Worker{mediator}, Worker{mediator},
		Worker{mediator}, Engineer{mediator}, Engineer{mediator},
	}

	for _, c := range components {
		c.Solve()
	}

	mediator.Wait()
	fmt.Println("all components have done their jobs")
}
