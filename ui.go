package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 80

var (
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().Foreground(lipgloss.Color("212"))
	s.Help = lg.NewStyle().Foreground(lipgloss.Color("240"))
	return &s
}

var commitTypes = []huh.Option[string]{
	huh.NewOption("feat - a new feature", "feat"),
	huh.NewOption("fix - a bug fix", "fix"),
	huh.NewOption("build - changes that affect the build system or external dependencies", "build"),
	huh.NewOption("chore - changes to the build process or auxiliary tools and libraries", "chore"),
	huh.NewOption("ci - changes to our CI configuration files and scripts", "ci"),
	huh.NewOption("docs - documentation only changes", "docs"),
	huh.NewOption("perf - a code change that improves performance", "perf"),
	huh.NewOption("refactor - a code change that neither fixes a bug nor adds a feature", "refactor"),
	huh.NewOption("revert - reverts a previous commit", "revert"),
	huh.NewOption("style - changes that do not affect the meaning of the code", "style"),
	huh.NewOption("test - adding missing tests or correcting existing tests", "test"),
}

type Model struct {
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int

	commit *Commit
}

func NewModel(msg *string) Model {
	const descriptionLength = 60
	const bodyLength = 1000

	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)
	m.commit = &Commit{}

	if msg != nil {
		trunc := min(descriptionLength, len(*msg))
		m.commit.Description = (*msg)[:trunc]
	}

	title := func(s string, i int) string {
		return fmt.Sprintf("%s %d characters remaining", s, i)
	}

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Key("type").
				Options(commitTypes...).
				Title("Commit Type").
				Description("What type of commit is this?").
				Value(&m.commit.Type),

			huh.NewInput().
				Key("scope").
				Title("Scope (optional)").
				Description("What is the commit affecting?").
				Value(&m.commit.Scope),

			huh.NewConfirm().
				Key("breaking").
				Title("Breaking Change?").
				Description("Does this commit introduce a breaking change?").
				Value(&m.commit.BreakingChange),
		),
		huh.NewGroup(
			huh.NewInput().
				Key("description").
				Title("Short Description").
				Description(title("What is the commit about at a high level?", descriptionLength)).
				DescriptionFunc(func() string {
					return title("What is the commit about at a high level?", descriptionLength-len(m.commit.Description))
				}, &m.commit.Description).
				CharLimit(descriptionLength).
				Value(&m.commit.Description).
				Validate(func(s string) error {
					if len(s) == 0 {
						return fmt.Errorf("description cannot be empty")
					}
					return nil
				}),

			huh.NewText().
				Key("body").
				Title("Body").
				Description(title("What is the commit about in more detail?", bodyLength)).
				DescriptionFunc(func() string {
					return title("What is the commit about in more detail?", bodyLength-len(m.commit.Body))
				}, &m.commit.Body).
				CharLimit(1000).
				Lines(10).
				Value(&m.commit.Body),
		),
	).
		WithWidth(maxWidth).
		WithShowHelp(false).
		WithShowErrors(false)
	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Interrupt
		case "esc":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	footer := lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(m.form.Help().ShortHelpView(m.form.KeyBinds())),
		lipgloss.WithWhitespaceChars("/"),
		lipgloss.WithWhitespaceForeground(indigo),
	)

	return m.styles.Base.Render(m.form.View() + "\n" + footer)
}
