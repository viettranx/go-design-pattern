package main

import "fmt"

type Tool struct {
	isUsing bool
}

func (t *Tool) CanUse() bool { return !t.isUsing }
func (t *Tool) Using()       { t.isUsing = true }
func (t *Tool) Done()        { t.isUsing = false }

type ProblemSolver interface {
	Solve(tool *Tool)
}

type Worker struct{}
type Engineer struct{}

func (w Worker) Solve(tool *Tool) {
	fmt.Println("Worker is using tool")
	tool.Using()
}

func (eng Engineer) Solve(tool *Tool) {
	fmt.Println("Engineer is using tool")
	tool.Using()
}

func main() {
	shareTool := &Tool{isUsing: false}

	worker := Worker{}
	engineer := Engineer{}

	if shareTool.CanUse() {
		worker.Solve(shareTool)
	}

	if shareTool.CanUse() {
		engineer.Solve(shareTool)
	}

	// engineer need to wait for worker finish her/his job,
	// but we don't have any notifier for them.
	// Furthermore, If more users want to use tool, we need a queue or something like that.
}
