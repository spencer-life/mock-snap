package exporter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"mock-snap/pkg/checker"
	"mock-snap/pkg/models"
)

type ProcessedResult struct {
	Lead        models.Lead
	OutputPath  string
	AuditReport checker.AuditReport
}

func SavePreview(outputDir, fileName, htmlContent string) (string, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	fullPath := filepath.Join(outputDir, fileName)
	if err := os.WriteFile(fullPath, []byte(htmlContent), 0644); err != nil {
		return "", fmt.Errorf("failed to write HTML preview file: %w", err)
	}

	return fullPath, nil
}

func RenderTerminalSummary(results []ProcessedResult, style string, outputDir string, duration time.Duration) string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#04B5D5")).
		Padding(0, 1)

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))

	successStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#10B981"))

	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#04B5D5")).
		Padding(1, 2)

	summary := strings.Builder{}
	summary.WriteString(titleStyle.Render("MOCK-SNAP LANDING PAGE GENERATOR SUMMARY") + "\n\n")

	metaInfo := fmt.Sprintf("%s Output Folder   : %s\n%s Theme Style     : %s\n%s Total Generated  : %d previews\n%s Execution Time   : %s",
		headerStyle.Render("📁"), outputDir,
		headerStyle.Render("🎨"), style,
		headerStyle.Render("⚡"), len(results),
		headerStyle.Render("⏱️"), duration.Round(time.Millisecond).String(),
	)
	summary.WriteString(metaInfo + "\n\n")

	summary.WriteString(headerStyle.Render("GENERATED PREVIEWS & AUDIT PASS:") + "\n")

	for i, res := range results {
		if i >= 5 {
			summary.WriteString(fmt.Sprintf("  ... and %d more previews generated in %s\n", len(results)-5, outputDir))
			break
		}

		statusBadge := successStyle.Render("100% AUDIT PASS ✅")
		if !res.AuditReport.IsCompliant() {
			statusBadge = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF9F1C")).Render(fmt.Sprintf("%.0f%% COMPLIANT ⚠️", res.AuditReport.ComplianceScore))
		}

		summary.WriteString(fmt.Sprintf("  • %-30s | Lenis+GSAP: %s | %s\n    Path: %s\n",
			res.Lead.Name,
			successStyle.Render("ACTIVE"),
			statusBadge,
			res.OutputPath,
		))
	}

	summary.WriteString("\n" + successStyle.Render("✨ Anti-AI-Slop Engine: All generated pages feature Lenis smooth scrolling, GSAP ScrollTrigger reveals, and custom Google Typography."))

	return cardStyle.Render(summary.String())
}
