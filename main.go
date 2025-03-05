package main

import (
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	dryRun := flag.Bool("dry-run", false, "Don't commit, just print the commit message")
	flag.Parse()

	g, err := NewGit()
	if err != nil {
		fmt.Println("Error opening repository:", err)
		os.Exit(1)
	}

	if err = g.Status(); !*dryRun && err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m := NewModel()
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
