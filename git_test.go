package main

import (
	"testing"
)

func TestCommit_GetCommitMessage(t *testing.T) {
	tests := []struct {
		name           string
		commit         Commit
		expectedOutput string
	}{
		// Test: Type and Description
		{
			name: "Type and description",
			commit: Commit{
				Type:           "feat",
				BreakingChange: false,
				Scope:          "",
				Description:    "Add user login feature",
				Body:           "",
			},
			expectedOutput: "feat: Add user login feature",
		},
		// Test: All fields with Scope without breaking change
		{
			name: "Type, scope, and description",
			commit: Commit{
				Type:           "fix",
				BreakingChange: false,
				Scope:          "auth",
				Description:    "Fix login issue",
				Body:           "",
			},
			expectedOutput: "fix(auth): Fix login issue",
		},
		// Test: All fields with Scope and breaking change
		{
			name: "Breaking change with scope",
			commit: Commit{
				Type:           "feat",
				BreakingChange: true,
				Scope:          "auth",
				Description:    "Introduce SSO login",
				Body:           "",
			},
			expectedOutput: "feat(auth)!: Introduce SSO login",
		},
		// Test: All fields including Body
		{
			name: "Complete commit message",
			commit: Commit{
				Type:           "feat",
				BreakingChange: true,
				Scope:          "auth",
				Description:    "Introduce SSO login",
				Body:           "This enables single sign-on functionality for external providers.",
			},
			expectedOutput: `feat(auth)!: Introduce SSO login

This enables single sign-on functionality for external providers.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			message := tt.commit.Message()
			if message != tt.expectedOutput {
				t.Errorf("Message() = %q, want %q", message, tt.expectedOutput)
			}
		})
	}
}
