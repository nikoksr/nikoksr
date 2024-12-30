package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

type User struct {
	Name  string
	Email string
}

type Project struct {
	Path        string
	Description string
}

type Profile struct {
	EnableEmojis bool
	Motto        string
	User         User
	Projects     []Project
}

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string) error {
	_, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	renderer, err := NewRenderer()
	if err != nil {
		return fmt.Errorf("create renderer: %w", err)
	}

	if len(args) < 1 {
		return errors.New("no args given")
	}

	cmd := normalizeInput(args[0])
	switch cmd {
	case "profile":
		if len(args) < 2 {
			return errors.New("missing file-path for profile output")
		}

		outputPath := normalizeInput(args[1])
		profile := myGitHubProfile()

		err = renderer.RenderProfileToFile(outputPath, profile)
	default:
		err = fmt.Errorf("unknown command: %s", cmd)
	}

	return err
}

func normalizeInput(input string) string {
	return strings.TrimSpace(input)
}

func myGitHubProfile() Profile {
	return Profile{
		EnableEmojis: true,
		Motto:        "Driven by curiosity. Always learning, always coding.",
		User: User{
			Name:  "Niko",
			Email: "i694aqo23@mozmail.com",
		},
		Projects: []Project{
			{
				Path:        "nikoksr/notify",
				Description: "A dead simple Go library for sending notifications to various messaging services.",
			},
			{
				Path:        "nikoksr/konfetty",
				Description: "Zero-dependency, type-safe and powerful post-processing for your existing config solution in Go.",
			},
			{
				Path:        "nikoksr/assert-go",
				Description: "Zero-dependency, idiomatic Go assertion library focused on crystal-clear failure messages and thoughtful source context. Inspired by [Tiger Style](https://github.com/tigerbeetle/tigerbeetle/blob/main/docs/TIGER_STYLE.md#safety).",
			},
			{
				Path:        "nikoksr/typeid-zig",
				Description: "Type-safe, K-sortable, globally unique identifier inspired by Stripe IDs implemented in Zig.",
			},
		},
	}
}
