package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	m := NewModel()

	_, err := tea.NewProgram(m).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	m.commit.Commit()
}
