package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"mock-snap/pkg/checker"
	"mock-snap/pkg/exporter"
	"mock-snap/pkg/generator"
)

var (
	inputFlag     string
	outputDirFlag string
	styleFlag     string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "mock-snap",
		Short: "High-end Anti-AI-Slop business landing page preview generator",
		Long: `mock-snap is an autonomous Go CLI tool that reads business leads JSON and generates high-end, animated landing page previews powered by Tailwind CSS, GSAP 3.x, Lenis smooth scrolling, and an automated Impeccable Design Quality Audit pass.`,
		RunE: run,
	}

	rootCmd.Flags().StringVarP(&inputFlag, "input", "i", "", "Path to leads JSON file (e.g. './leads.json' or '/home/mlpc/dev/local-prospect/leads.json')")
	rootCmd.Flags().StringVarP(&outputDirFlag, "output-dir", "o", "./previews", "Output directory for HTML landing page previews")
	rootCmd.Flags().StringVarP(&styleFlag, "style", "s", "luxury_dark", "Theme style preset ('luxury_dark', 'warm_editorial', 'clean_architect')")

	_ = rootCmd.MarkFlagRequired("input")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	if strings.TrimSpace(inputFlag) == "" {
		return fmt.Errorf("input flag is required (e.g. --input './leads.json')")
	}

	startTime := time.Now()

	fmt.Printf("🔍 Loading lead records from '%s'...\n", inputFlag)
	leads, err := generator.LoadLeads(inputFlag)
	if err != nil {
		return fmt.Errorf("error loading leads: %w", err)
	}

	if len(leads) == 0 {
		fmt.Printf("⚠️ No leads found in '%s'.\n", inputFlag)
		return nil
	}

	fmt.Printf("🎨 Generating Anti-AI-Slop previews for %d leads (Style: '%s')...\n", len(leads), styleFlag)

	genInst := generator.NewGenerator()
	chkInst := checker.NewChecker()

	var results []exporter.ProcessedResult

	for _, lead := range leads {
		// 1. Render draft HTML from template
		draftHTML, err := genInst.RenderHTML(lead, styleFlag)
		if err != nil {
			return fmt.Errorf("error rendering template for lead '%s': %w", lead.Name, err)
		}

		// 2. Perform Impeccable Design Quality Audit & Auto-repair
		auditReport, cleanedHTML := chkInst.AuditHTML(draftHTML)

		// 3. Save preview file
		fileName := generator.SanitizeFilename(lead.Name)
		outputPath, err := exporter.SavePreview(outputDirFlag, fileName, cleanedHTML)
		if err != nil {
			return fmt.Errorf("error saving preview for lead '%s': %w", lead.Name, err)
		}

		results = append(results, exporter.ProcessedResult{
			Lead:        lead,
			OutputPath:  filepath.Clean(outputPath),
			AuditReport: auditReport,
		})
	}

	duration := time.Since(startTime)

	// 4. Render Lipgloss Summary Card
	summary := exporter.RenderTerminalSummary(results, styleFlag, filepath.Clean(outputDirFlag), duration)
	fmt.Println("\n" + summary)

	return nil
}
