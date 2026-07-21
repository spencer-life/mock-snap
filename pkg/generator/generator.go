package generator

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"strings"

	"mock-snap/pkg/models"
)

//go:embed templates/*.html
var templateFS embed.FS

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func LoadLeads(filePath string) ([]models.Lead, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read input leads file: %w", err)
	}

	var leads []models.Lead
	if err := json.Unmarshal(data, &leads); err != nil {
		return nil, fmt.Errorf("failed to parse leads JSON: %w", err)
	}

	return leads, nil
}

func (g *Generator) RenderHTML(lead models.Lead, style string) (string, error) {
	if style == "" {
		style = "luxury_dark"
	}

	templatePath := fmt.Sprintf("templates/%s.html", style)
	tmplContent, err := templateFS.ReadFile(templatePath)
	if err != nil {
		// Fallback to luxury_dark if style not found
		templatePath = "templates/luxury_dark.html"
		tmplContent, err = templateFS.ReadFile(templatePath)
		if err != nil {
			return "", fmt.Errorf("failed to read embedded template: %w", err)
		}
	}

	funcMap := template.FuncMap{
		"slice": func(s string, start, end int) string {
			if len(s) == 0 {
				return "B"
			}
			if end > len(s) {
				end = len(s)
			}
			return s[start:end]
		},
	}

	tmpl, err := template.New("landing").Funcs(funcMap).Parse(string(tmplContent))
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, lead); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func SanitizeFilename(name string) string {
	name = strings.ToLower(name)
	replacer := strings.NewReplacer(" ", "_", "/", "_", "\\", "_", ":", "", "'", "", "\"", "", "..", "", "~", "")
	clean := replacer.Replace(name)
	clean = strings.Trim(clean, "._")
	if clean == "" {
		return "lead_preview.html"
	}
	return clean + "_preview.html"
}
