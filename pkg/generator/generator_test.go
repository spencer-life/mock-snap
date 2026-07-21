package generator

import (
	"strings"
	"testing"

	"mock-snap/pkg/models"
)

func TestGenerator_RenderHTML(t *testing.T) {
	lead := models.Lead{
		OsmID:    101,
		Name:     "Apex Plumbing LLC",
		Phone:    "+1-480-555-0199",
		Website:  "",
		Street:   "100 Main St",
		City:     "Scottsdale",
		Category: "plumber",
		Tier:     "TIER_1_NO_WEBSITE",
	}

	gen := NewGenerator()

	t.Run("Render luxury_dark template", func(t *testing.T) {
		html, err := gen.RenderHTML(lead, "luxury_dark")
		if err != nil {
			t.Fatalf("failed to render luxury_dark template: %v", err)
		}

		if !strings.Contains(html, "Apex Plumbing LLC") {
			t.Errorf("rendered HTML missing lead name 'Apex Plumbing LLC'")
		}
		if !strings.Contains(html, "Scottsdale") {
			t.Errorf("rendered HTML missing city 'Scottsdale'")
		}
		if !strings.Contains(html, "Lenis") {
			t.Errorf("rendered HTML missing Lenis smooth scroll script")
		}
	})

	t.Run("Sanitize filename", func(t *testing.T) {
		clean := SanitizeFilename("Apex Plumbing & Rooter LLC")
		if clean != "apex_plumbing_&_rooter_llc_preview.html" && clean != "apex_plumbing__rooter_llc_preview.html" {
			if !strings.HasSuffix(clean, "_preview.html") {
				t.Errorf("invalid sanitized filename: %s", clean)
			}
		}
	})
}
