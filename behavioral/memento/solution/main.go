package main

import "fmt"

// Originator & Memento have to be in the same package

type Editor struct {
	str string
}

func (e *Editor) TypeMore(s string) { e.str += s }
func (e *Editor) Content() string   { return e.str }
func (e *Editor) Save() Memento     { return Memento{content: e.str} }
func (e *Editor) Restore(m Memento) { e.str = m.content }

type Memento struct {
	// store all originator's attributes
	content string
}

type Caretaker struct {
	mementos []Memento
}

func (c *Caretaker) AddMemento(m Memento) { c.mementos = append(c.mementos, m) }

func (c *Caretaker) Restore(editor *Editor, idx int) { editor.Restore(c.mementos[idx]) }

func (c *Caretaker) Size() int { return len(c.mementos) }

func NewEditor(s string) Editor {
	return Editor{s}
}

func main() {
	editor := NewEditor("")
	caretaker := &Caretaker{}

	editor.TypeMore("hello")
	editor.TypeMore(" world")

	caretaker.AddMemento(editor.Save())

	editor.TypeMore(". I am Memento")
	caretaker.AddMemento(editor.Save())

	editor.TypeMore(", a design pattern lets you save and ")
	editor.TypeMore("restore the previous state")

	for i := 0; i < caretaker.Size(); i++ {
		caretaker.Restore(&editor, i)
		fmt.Println("Restored content:", editor.Content())
	}
}
