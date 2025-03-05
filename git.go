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

type Git struct {
	repo     *git.Repository
	worktree *git.Worktree
}

func NewGit() (*Git, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}

	worktree, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	return &Git{
		repo:     repo,
		worktree: worktree,
	}, nil
}

func (g *Git) Status() error {
	status, err := g.worktree.Status()
	if err != nil {
		fmt.Println("Error getting status:", err)
	}

	emptyWorktree := false
	for _, s := range status {
		if s.Staging == git.Modified || s.Staging == git.Added {
			emptyWorktree = true
			break
		}
	}

	if !emptyWorktree {
		return fmt.Errorf("repository contains unstaged changes. Please stage or discard them before committing")
	}

	return nil
}

func (g *Git) Commit(c *Commit) {
	confirm := false
	if err := huh.NewConfirm().Title("Commit?").Affirmative("Yes").Negative("Cancel").Value(&confirm).Run(); err != nil {
		return
	}
	if !confirm {
		return
	}

	hash, err := g.worktree.Commit(c.Message(), &git.CommitOptions{})
	if err != nil {
		fmt.Println("Error during commit:", err)
		return
	}

	obj, err := g.repo.CommitObject(hash)
	if err != nil {
		fmt.Println("Error getting commit object:", err)
		return
	}

	fmt.Println("Committed:")
	fmt.Println(obj)
}
