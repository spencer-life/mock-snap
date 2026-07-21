package checker

import (
	"strings"
	"testing"
)

func TestChecker_AuditHTML(t *testing.T) {
	t.Run("Valid Impeccable HTML Template Passes 100%", func(t *testing.T) {
		validHTML := `<!DOCTYPE html>
		<html>
		<head>
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<link href="https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@700&display=swap" rel="stylesheet">
			<script src="https://unpkg.com/lenis@1.1.18/dist/lenis.min.js"></script>
			<script src="https://cdn.jsdelivr.net/npm/gsap@3.12.5/dist/ScrollTrigger.min.js"></script>
		</head>
		<body style="background: #0b0f19; color: #f8fafc;">
			<a href="tel:4805550199">Call Now</a>
			<script>
				const lenis = new Lenis({ autoRaf: true });
				lenis.on('scroll', ScrollTrigger.update);
			</script>
		</body>
		</html>`

		chk := NewChecker()
		report, cleaned := chk.AuditHTML(validHTML)

		if !report.IsCompliant() {
			t.Errorf("expected 100%% compliance, got %.1f%%. Violations: %v", report.ComplianceScore, report.Violations)
		}
		if !report.HasLenisGSAP {
			t.Errorf("expected HasLenisGSAP to be true")
		}
		if !report.HasViewport {
			t.Errorf("expected HasViewport to be true")
		}
		if !report.HasCustomFont {
			t.Errorf("expected HasCustomFont to be true")
		}
		if !report.HasTelLink {
			t.Errorf("expected HasTelLink to be true")
		}
		if cleaned != validHTML {
			t.Errorf("cleaned HTML modified compliant template unexpectedly")
		}
	})

	t.Run("Flawed HTML Template Triggers Auto-Repair", func(t *testing.T) {
		flawedHTML := `<!DOCTYPE html>
		<html>
		<head>
			<link href="https://fonts.googleapis.com/css2?family=Plus+Jakarta+Sans:wght@700&display=swap" rel="stylesheet">
			<script src="https://unpkg.com/lenis@1.1.18/dist/lenis.min.js"></script>
			<script src="https://cdn.jsdelivr.net/npm/gsap@3.12.5/dist/ScrollTrigger.min.js"></script>
		</head>
		<body style="background: #000000; color: #ffffff;">
			<a href="tel:4805550199">Call Now</a>
			<script>
				const lenis = new Lenis({ autoRaf: true });
				lenis.on('scroll', ScrollTrigger.update);
			</script>
		</body>
		</html>`

		chk := NewChecker()
		report, cleaned := chk.AuditHTML(flawedHTML)

		if strings.Contains(cleaned, "#000000") || strings.Contains(cleaned, "#ffffff") {
			t.Errorf("auto-repair failed to replace pure black and pure white inline colors")
		}
		if !report.HasViewport {
			t.Errorf("expected auto-repair to inject missing viewport tag")
		}
	})
}
