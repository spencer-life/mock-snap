<p align="center">
  <img src="banner.jpg" alt="mock-snap banner" width="700">
</p>

<h1 align="center">mock-snap</h1>

<p align="center">
  <strong>Autonomous Go CLI & UX design engine generating high-end, animated Anti-AI-Slop landing page previews.</strong><br>
  Powered by Tailwind CSS, GSAP 3.x, Lenis smooth scrolling, and automated design audit passes.
</p>

<p align="center">
  <a href="https://github.com/spencer-life/mock-snap/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
  <a href="https://golang.org"><img src="https://img.shields.io/badge/Go-1.26+-00ADD8.svg?style=flat&logo=go" alt="Go Version"></a>
  <img src="https://img.shields.io/badge/Design-Anti--AI--Slop-7d56f4.svg" alt="Anti-AI-Slop">
  <img src="https://img.shields.io/badge/Animations-GSAP%20%2B%20Lenis-04b5d5.svg" alt="GSAP + Lenis">
</p>

<p align="center">
  <a href="#-installation--usage">Installation</a> •
  <a href="#-features--aesthetics">Features</a> •
  <a href="#-theme-presets">Presets</a> •
  <a href="#-automated-impeccable-design-quality-audit">Audit Engine</a>
</p>

---

## 🎨 Features & Aesthetics

- **Anti-AI-Slop Aesthetics**:
  - Asymmetrical 2-column hero layouts.
  - Border-separated service list rows (strictly NO generic 3-card icon grids).
  - Paired Google Fonts (`Plus Jakarta Sans` for headings, `Outfit` or `Instrument Sans` for body text).
  - Zero raw `#000000` / `#ffffff` inline colors—uses warm tinted neutrals (`slate-950` / `zinc-100`).
- **Interactive Micro-Animations**:
  - Lenis smooth scroll (`autoRaf: true`).
  - GSAP ScrollTrigger reveals (staggered opacity/transform-y entries).
- **Automated Impeccable Design Quality Audit Pass (`pkg/checker`)**:
  - Validates Lenis/GSAP script synchronization, mobile viewport meta tags, custom Google Fonts links, valid phone CTA links (`href="tel:..."`), and color safety before writing HTML to disk.

---

## 🎭 Theme Presets

- **`luxury_dark` (Default)**: Dark glassmorphism, violet/cyan glows, Plus Jakarta Sans + Outfit.
- **`warm_editorial`**: Dark espresso background, warm amber accents, Instrument Sans.
- **`clean_architect`**: Deep slate background, sky blue accents, architectural grid lines.

---

## 📦 Installation & Usage

### Install via Go
```bash
go install github.com/spencer-life/mock-snap@latest
```

### Install AI Agent Skill (Cross-Agent)
```bash
npx skills add spencer-life/mock-snap
```

---

## 🛠️ CLI Flags

```bash
mock-snap --input "<path_to_leads.json>" [flags]
```

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--input` | `-i` | Path to leads JSON file | Required |
| `--output-dir` | `-o` | Output directory for HTML previews | `./previews` |
| `--style` | `-s` | Theme preset (`luxury_dark`, `warm_editorial`, `clean_architect`) | `luxury_dark` |

---

## 💡 Example

```bash
mock-snap --input /home/mlpc/dev/local-prospect/leads.json --output-dir ./previews --style luxury_dark
```

---

## 🧪 Testing

Run the Go test suite:
```bash
go test -v ./...
```
