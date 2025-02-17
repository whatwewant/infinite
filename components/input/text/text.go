package text

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/fzdwx/infinite/components"
)

type Text struct {
	inner   *components.Input
	startUp *components.StartUp
}

func New(ops ...Option) *Text {
	inner := components.NewInput()
	i := &Text{inner: inner, startUp: components.NewStartUp(inner)}

	i.inner.QuitKey = key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "quit input text"))

	i.Apply(ops...)

	return i
}

// Apply options on Text
func (i *Text) Apply(ops ...Option) *Text {
	if len(ops) > 0 {
		for _, option := range ops {
			option(i)
		}
	}
	return i
}

func (i *Text) Display() error {
	return i.startUp.Start()
}

// Focus sets the Focus state on the model. When the model is in Focus it can
// receive keyboard input and the cursor will be hidden.
func (i *Text) Focus() {
	i.inner.Focus()
}

// Blur removes the Focus state on the model.  When the model is blurred it can
// not receive keyboard input and the cursor will be hidden.
func (i *Text) Blur() {
	i.inner.Blur()
}

// Quit input
func (i *Text) Quit() {
	i.inner.Quit()
}

// Value returns the value of the text input.
func (i *Text) Value() string {
	return i.inner.Value()
}

// Focused returns the focus state on the model.
func (i *Text) Focused() bool {
	return i.inner.Focused()
}

// Cursor returns the cursor position.
func (i *Text) Cursor() int {
	return i.inner.Cursor()
}

// Blink returns whether or not to draw the cursor.
func (i *Text) Blink() bool {
	return i.inner.Blink()
}

// SetCursor moves the cursor to the given position. If the position is
// out of bounds the cursor will be moved to the start or end accordingly.
func (i *Text) SetCursor(pos int) {
	i.inner.SetCursor(pos)
}

// CursorMode returns the model's cursor mode. For available cursor modes, see
// type CursorMode.
func (i *Text) CursorMode() components.CursorMode {
	return i.inner.CursorMode()
}

// SetCursorMode sets the model's cursor mode. This method returns a command.
//
// For available cursor modes, see type CursorMode.
func (i *Text) SetCursorMode(model components.CursorMode) {
	i.inner.SetCursorMode(model)
}

// CursorStart moves the cursor to the start of the input field.
func (i *Text) CursorStart() {
	i.inner.CursorStart()
}

// CursorEnd moves the cursor to the end of the input field.
func (i *Text) CursorEnd() {
	i.inner.CursorEnd()
}

// Reset sets the input to its default state with no input. Returns whether
// or not the cursor blink should reset.
func (i *Text) Reset() bool {
	return i.inner.Reset()
}
