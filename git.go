package main

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/go-git/go-git/v5"
	"strings"
	"text/template"
)

const templateString = `{{.Type}}{{if .Scope}}({{.Scope}}){{end}}{{if .BreakingChange}}!{{end}}: {{.Description}}{{if .Body}}

{{.Body}}{{end}}`

var textTemplate = template.Must(template.New("commit").Parse(templateString))

type Commit struct {
	Type           string
	BreakingChange bool
	Scope          string
	Description    string
	Body           string
}

func (c *Commit) Message() string {
	var buf strings.Builder
	err := textTemplate.Execute(&buf, c)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return ""
	}
	return buf.String()
}

func (c *Commit) Commit() {
	repo, err := git.PlainOpen(".")
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return
	}

	w, err := repo.Worktree()
	if err != nil {
		fmt.Println("Error getting worktree:", err)
		return
	}

	fmt.Println("Repository status before commit:")
	status, err := w.Status()
	if err != nil {
		fmt.Println("Error getting status:", err)
	}
	fmt.Println(status)

	confirm := false
	if err = huh.NewConfirm().Title("Commit?").Affirmative("Yes").Negative("Cancel").Value(&confirm).Run(); err != nil {
		return
	}
	if !confirm {
		return
	}

	hash, err := w.Commit(c.Message(), &git.CommitOptions{})
	if err != nil {
		fmt.Println("Error getting worktree:", err)
		return
	}

	println("Committed:")

	obj, err := repo.CommitObject(hash)
	if err != nil {
		fmt.Println("Error getting commit object:", err)
		return
	}

	fmt.Println(obj)
}
