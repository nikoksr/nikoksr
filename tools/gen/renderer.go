package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
)

const (
	templatesGlobPattern = "templates/*.tmpl"
	templateProfile      = "profile.md.tmpl"
)

//go:embed "templates/*"
var templatesFS embed.FS

type Renderer struct {
	tmpl *template.Template
}

func NewRenderer() (*Renderer, error) {
	tmpl, err := template.ParseFS(templatesFS, templatesGlobPattern)
	if err != nil {
		return nil, err
	}

	return &Renderer{tmpl: tmpl}, nil
}

func (r *Renderer) RenderProfile(w io.Writer, profile Profile) error {
	return r.tmpl.ExecuteTemplate(w, templateProfile, profile)
}

func (r *Renderer) RenderProfileToFile(outPath string, profile Profile) error {
	if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer file.Close()

	return r.RenderProfile(file, profile)
}
