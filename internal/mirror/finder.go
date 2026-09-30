package mirror

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/termux/termux-best-mirror/internal/system"
)

func GetRegions() []string {
	var regions []string
	entries, err := os.ReadDir(system.MirrorsDir)
	if err != nil {
		return regions
	}

	regions = append(regions, "all")
	for _, e := range entries {
		if e.IsDir() {
			regions = append(regions, e.Name())
		}
	}
	return regions
}

func FormatRegionName(region string) string {
	switch region {
	case "all":
		return "All mirrors"
	case "asia":
		return "Asia (excl. Chinese Mainland & Russia)"
	case "chinese_mainland":
		return "Chinese Mainland"
	case "europe":
		return "Europe"
	case "north_america":
		return "North America"
	case "oceania":
		return "Oceania"
	case "russia":
		return "Russia"
	default:
		return strings.Title(strings.ReplaceAll(region, "_", " "))
	}
}

func GetMirrors(region string) []string {
	var mirrors []string
	searchDir := system.MirrorsDir
	if region != "all" {
		searchDir = filepath.Join(system.MirrorsDir, region)
	}

	filepath.WalkDir(searchDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if !strings.HasSuffix(d.Name(), ".dpkg-old") && !strings.HasSuffix(d.Name(), ".dpkg-new") && !strings.HasSuffix(d.Name(), "~") {
				mirrors = append(mirrors, path)
			}
		}
		return nil
	})
	return mirrors
}

func GetMirrorDescription(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) >= 2 {
		parts := strings.SplitN(lines[1], " ", 2)
		if len(parts) > 1 {
			return parts[1]
		}
	}
	return ""
}

func GetMirrorURL(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "MAIN=") {
			parts := strings.Split(line, `"`)
			if len(parts) >= 3 {
				url := parts[1]
				if !strings.HasSuffix(url, "/") {
					url += "/"
				}
				return url
			}
		}
	}
	return ""
}
