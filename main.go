package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	dryRun := pflag.BoolP("dry-run", "d", false, "Don't commit, just print the commit message")
	message := pflag.StringP("message", "m", "", "Commit message")
	help := pflag.BoolP("help", "h", false, "Show help")
	pflag.Parse()

	if *help {
		pflag.Usage()
		os.Exit(0)
	}

	g, err := NewGit()
	if err != nil {
		fmt.Println("Error opening repository:", err)
		os.Exit(1)
	}

	if err = g.Status(); !*dryRun && err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := NewModel(message)
	_, err = tea.NewProgram(m).Run()
	if err != nil {
		os.Exit(1)
	}

	if *dryRun {
		fmt.Println("Commit message:")
		fmt.Println(m.commit.Message())
	} else {
		g.Commit(m.commit)
	}
}
