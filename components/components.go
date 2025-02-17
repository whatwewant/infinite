package components

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/fzdwx/infinite/color"
	"github.com/fzdwx/infinite/style"
	"github.com/fzdwx/infinite/theme"
	"time"
)

type (
	/*
		Components, You can use these components directly:
			 	1. Input
			 	2. Selection
			 	3. Spinner
				4. Autocomplete
				5. Progress
		Or use them inline in your custom component,
		for how to embed them, you can refer to the implementation of `Confirm`.
	*/
	Components interface {
		tea.Model

		// SetProgram this method will be called back when the tea.Program starts.
		// please keep passing this method
		SetProgram(program *tea.Program)
	}
)

// NewAutocomplete constructor
func NewAutocomplete(suggester Suggester) *Autocomplete {
	return &Autocomplete{
		Suggester:            suggester,
		Completer:            DefaultCompleter(),
		Input:                NewInput(),
		KeyMap:               DefaultAutocompleteKeyMap(),
		ShowSelection:        true,
		ShouldNewSelection:   true,
		SelectionCreator:     DefaultSelectionCreator,
		SuggestionViewRender: NewLineSuggestionRender,
		//SuggestionViewRender: TabSuggestionRender,
	}
}

// NewInput constructor
func NewInput() *Input {
	c := &Input{
		Model:            textinput.New(),
		Status:           Focus,
		Prompt:           "> ",
		Placeholder:      "",
		BlinkSpeed:       DefaultBlinkSpeed,
		EchoMode:         EchoNormal,
		EchoCharacter:    '*',
		CharLimit:        0,
		QuitKey:          key.NewBinding(),
		PlaceholderStyle: style.New().Fg(color.Gray),
		PromptStyle:      style.New(),
		TextStyle:        style.New(),
		BackgroundStyle:  style.New(),
		CursorStyle:      style.New(),
	}

	return c
}

// NewPrintHelper constructor
func NewPrintHelper(program *tea.Program) *PrintHelper {
	return &PrintHelper{program: program}
}

// NewProgress constructor
func NewProgress() *Progress {
	p := &Progress{
		Id:              nextID(),
		Total:           100,
		Current:         0,
		PercentAgeFunc:  DefaultPercentAgeFunc,
		PercentAgeStyle: style.New().Inline(),
		Width:           defaultWidth,
		Full:            '█',
		FullColor:       "#7571F9",
		Empty:           '░',
		EmptyColor:      "#606060",
		ShowPercentage:  true,
		ShowCost:        true,
		prevAmount:      0,
		CostView:        DefaultCostView,
		TickCostDelay:   defaultTicKCostDelay,
	}

	return p
}

// NewSelection constructor
func NewSelection(choices []string) *Selection {

	items := slice.Map[string, SelectionItem](choices, func(idx int, item string) SelectionItem {
		return SelectionItem{idx, item}
	})

	c := &Selection{
		Choices:             items,
		Selected:            make(map[int]struct{}),
		CursorSymbol:        ">",
		UnCursorSymbol:      " ",
		CursorSymbolStyle:   theme.DefaultTheme.CursorSymbolStyle,
		ChoiceTextStyle:     theme.DefaultTheme.ChoiceTextStyle,
		Prompt:              "Please Selection your options:",
		PromptStyle:         theme.DefaultTheme.PromptStyle,
		HintSymbol:          "✓",
		HintSymbolStyle:     theme.DefaultTheme.MultiSelectedHintSymbolStyle,
		UnHintSymbol:        "✗",
		UnHintSymbolStyle:   theme.DefaultTheme.UnHintSymbolStyle,
		quited:              false,
		DisableOutPutResult: false,
		PageSize:            5,
		Keymap:              DefaultMultiKeyMap,
		Help:                help.New(),
		RowRender:           DefaultRowRender,
		EnableFilter:        true,
		FilterInput:         NewInput(),
		FilterFunc:          DefaultFilterFunc,
		ShowHelp:            true,
	}

	return c
}

// NewSpinner constructor
func NewSpinner() *Spinner {
	c := &Spinner{
		Model:               spinner.New(),
		Shape:               Line,
		ShapeStyle:          theme.DefaultTheme.SpinnerShapeStyle,
		Prompt:              "Loading...",
		DisableOutPutResult: false,
		Status:              Normal,
	}
	return c
}

const (
	DefaultBlinkSpeed = time.Millisecond * 530
)

// Status About the state of the Component
type Status int

const (
	// Focus only use Input
	Focus Status = iota
	// Blur only use Input
	Blur
	// Quit component
	Quit
	// Normal ignore it
	Normal
)

// CursorMode describes the behavior of the cursor.
type CursorMode int

const (
	CursorBlink CursorMode = iota
	CursorStatic
	CursorHide
)

// String returns the cursor mode in a human-readable format. This method is
// provisional and for informational purposes only.
func (c CursorMode) String() string {
	return [...]string{
		"blink",
		"static",
		"hidden",
	}[c]
}

func (c CursorMode) Map() textinput.CursorMode {
	switch c {
	case CursorBlink:
		return textinput.CursorBlink
	case CursorStatic:
		return textinput.CursorStatic
	case CursorHide:
		return textinput.CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%d", c))
}

func newCursorMode(other textinput.CursorMode) CursorMode {
	switch other {
	case textinput.CursorBlink:
		return CursorBlink
	case textinput.CursorStatic:
		return CursorStatic
	case textinput.CursorHide:
		return CursorHide
	}

	panic(fmt.Sprintf("unknow cursorMode :%s", other))
}

// EchoMode sets the Input behavior of the text Input field.
type EchoMode int

const (
	// EchoNormal displays text as is. This is the default behavior.
	EchoNormal EchoMode = iota
	// EchoPassword displays the EchoCharacter mask instead of actual
	// characters.  This is commonly used for password fields.
	EchoPassword
	// EchoNone displays nothing as characters are entered. This is commonly
	// seen for password fields on the command line.
	EchoNone
)
