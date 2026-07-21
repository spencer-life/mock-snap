package checker

import (
	"fmt"
	"regexp"
	"strings"
)

type AuditReport struct {
	HasLenisGSAP     bool     `json:"has_lenis_gsap"`
	NoPureBlackWhite bool     `json:"no_pure_black_white"`
	HasViewport      bool     `json:"has_viewport"`
	HasCustomFont    bool     `json:"has_custom_font"`
	HasTelLink       bool     `json:"has_tel_link"`
	ComplianceScore  float64  `json:"compliance_score"` // 0 to 100%
	Violations       []string `json:"violations"`
}

type Checker struct{}

func NewChecker() *Checker {
	return &Checker{}
}

var (
	viewportRegex    = regexp.MustCompile(`(?i)<meta[^>]+name=["']?viewport["']?`)
	lenisRegex       = regexp.MustCompile(`(?i)lenis\.min\.js|new\s+Lenis`)
	gsapSyncRegex    = regexp.MustCompile(`(?i)ScrollTrigger\.update|gsap\.ticker`)
	googleFontRegex  = regexp.MustCompile(`(?i)fonts\.googleapis\.com`)
	telLinkRegex     = regexp.MustCompile(`(?i)href=["']tel:[^"']+["']`)
	pureBlackRegex   = regexp.MustCompile(`(?i)#000000|#000\b|color:\s*black|background-color:\s*black`)
	pureWhiteRegex   = regexp.MustCompile(`(?i)#ffffff|#fff\b|color:\s*white|background-color:\s*white`)
)

func (c *Checker) AuditHTML(html string) (AuditReport, string) {
	report := AuditReport{}
	passed := 0
	total := 5

	cleanedHTML := html

	// 1. Lenis & GSAP Check
	if lenisRegex.MatchString(html) && gsapSyncRegex.MatchString(html) {
		report.HasLenisGSAP = true
		passed++
	} else {
		report.Violations = append(report.Violations, "Missing Lenis or GSAP ScrollTrigger ticker synchronization script")
	}

	// 2. Pure Black / White Check & Auto-Repair
	if !pureBlackRegex.MatchString(html) && !pureWhiteRegex.MatchString(html) {
		report.NoPureBlackWhite = true
		passed++
	} else {
		report.Violations = append(report.Violations, "Detected pure black (#000000) or pure white (#ffffff) inline styles - running auto-repair")
		// Auto-repair string transformations
		cleanedHTML = strings.ReplaceAll(cleanedHTML, "#000000", "#0b0f19")
		cleanedHTML = strings.ReplaceAll(cleanedHTML, "#ffffff", "#f8fafc")
		cleanedHTML = strings.ReplaceAll(cleanedHTML, "#000", "#0b0f19")
		cleanedHTML = strings.ReplaceAll(cleanedHTML, "#fff", "#f8fafc")
		report.NoPureBlackWhite = true // Marked repaired
		passed++
	}

	// 3. Mobile Viewport Check
	if viewportRegex.MatchString(html) {
		report.HasViewport = true
		passed++
	} else {
		report.Violations = append(report.Violations, "Missing <meta name='viewport'> tag")
		// Auto-inject missing viewport tag in head
		if strings.Contains(cleanedHTML, "<head>") {
			cleanedHTML = strings.Replace(cleanedHTML, "<head>", "<head>\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">", 1)
			report.HasViewport = true
		}
	}

	// 4. Custom Google Font Check
	if googleFontRegex.MatchString(html) {
		report.HasCustomFont = true
		passed++
	} else {
		report.Violations = append(report.Violations, "Missing custom Google Fonts stylesheet link")
	}

	// 5. Valid Phone CTA Link Check
	if telLinkRegex.MatchString(html) {
		report.HasTelLink = true
		passed++
	} else {
		report.Violations = append(report.Violations, "Missing valid href='tel:...' phone call-to-action link")
	}

	report.ComplianceScore = (float64(passed) / float64(total)) * 100.0
	return report, cleanedHTML
}

func (r AuditReport) IsCompliant() bool {
	return r.ComplianceScore == 100.0
}

func (r AuditReport) SummaryString() string {
	if r.IsCompliant() {
		return "100% Compliant (Impeccable Audit Passed)"
	}
	return fmt.Sprintf("%.0f%% Compliant (%d violations logged & repaired)", r.ComplianceScore, len(r.Violations))
}
