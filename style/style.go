package style

import (
	"github.com/charmbracelet/lipgloss"
)

func New() *Style {
	return newStyle()
}

func newStyle() *Style {
	style := lipgloss.NewStyle()
	return &Style{inner: style}
}

type Style struct {
	inner lipgloss.Style
}

func (s *Style) Bold() *Style {
	s.inner = s.inner.Bold(true)
	return s
}

func (s *Style) Italic() *Style {
	s.inner = s.inner.Italic(true)
	return s
}

func (s *Style) Fg(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.Foreground(color)
	return s
}

func (s *Style) Bg(color lipgloss.TerminalColor) *Style {
	s.inner = s.inner.Background(color)
	return s
}

func (s *Style) Render(str string) string {
	return s.inner.Render(str)
}

func (s *Style) Inner() lipgloss.Style {
	return s.inner.Copy()
}
