package main

type Editor struct {
	str string
}

func (e *Editor) TypeMore(s string) { e.str += s }
func (e *Editor) Content() string   { return e.str }

func NewEditor(s string) Editor {
	return Editor{s}
}

func main() {
	editor := NewEditor("")

	editor.TypeMore("hello")
	editor.TypeMore(" world")

	// how can I save current content and restore previous contents
	// WITHOUT know details of editor (all attributes or properties)
	// AND we can't do it outside the editor
}
