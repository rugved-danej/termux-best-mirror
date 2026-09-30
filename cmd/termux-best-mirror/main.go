package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/pterm/pterm"
	
	"github.com/termux/termux-best-mirror/internal/mirror"
	"github.com/termux/termux-best-mirror/internal/system"
	"github.com/termux/termux-best-mirror/internal/tester"
)

func main() {
	pterm.DefaultCenter.Println(pterm.Cyan(pterm.DefaultHeader.WithFullWidth().Sprint("Termux Best Mirror")))
	pterm.Println()

	system.CheckEnvironment()

	regions := mirror.GetRegions()
	if len(regions) == 0 {
		pterm.Error.Println("No mirror folders found.")
		os.Exit(1)
	}

	var regionOptions []string
	regionMap := make(map[string]string)
	for _, r := range regions {
		display := mirror.FormatRegionName(r)
		regionOptions = append(regionOptions, display)
		regionMap[display] = r
	}

	selectedDisplay, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select Mirror Region to Test").
		WithOptions(regionOptions).
		WithMaxHeight(12).
		Show()

	selectedFolder := regionMap[selectedDisplay]
	pterm.Success.Printf("Region: %s\n", pterm.Cyan(selectedDisplay))

	mirrors := mirror.GetMirrors(selectedFolder)
	if len(mirrors) == 0 {
		pterm.Error.Println("No mirrors found in the selected region.")
		os.Exit(1)
	}

	pterm.Info.Printf("Found %d mirrors to test.\n", len(mirrors))

	results := tester.TestMirrors(mirrors)

	if len(results) == 0 {
		pterm.Error.Println("No mirrors responded successfully.")
		os.Exit(1)
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Success != results[j].Success {
			return results[i].Success
		}
		if results[i].Duration == results[j].Duration {
			return results[i].Name < results[j].Name
		}
		return results[i].Duration < results[j].Duration
	})

	var resultOptions []string
	resultMap := make(map[string]mirror.Result)

	for _, r := range results {
		desc := mirror.GetMirrorDescription(r.Path)
		
		termWidth := pterm.GetTerminalWidth()
		maxDescLen := termWidth - len(r.Name) - 27
		if maxDescLen < 0 {
			desc = ""
		} else if len(desc) > maxDescLen {
			runes := []rune(desc)
			if len(runes) > maxDescLen {
				desc = string(runes[:maxDescLen]) + "..."
			}
		}

		var display string
		if !r.Success {
			if desc == "" {
				display = fmt.Sprintf("%s (Failed after %.3fs)", pterm.Red(r.Name), r.Duration.Seconds())
			} else {
				display = fmt.Sprintf("%s - %s (Failed after %.3fs)", pterm.Red(r.Name), pterm.Gray(desc), r.Duration.Seconds())
			}
		} else {
			var color func(a ...interface{}) string
			if r.Duration < 200*time.Millisecond {
				color = pterm.Green
			} else if r.Duration < 500*time.Millisecond {
				color = pterm.Yellow
			} else {
				color = pterm.Red
			}

			if desc == "" {
				display = fmt.Sprintf("%s (%.3fs)", color(r.Name), r.Duration.Seconds())
			} else {
				display = fmt.Sprintf("%s - %s (%.3fs)", color(r.Name), pterm.Gray(desc), r.Duration.Seconds())
			}
		}

		resultOptions = append(resultOptions, display)
		resultMap[display] = r
	}

	print("\033[H\033[2J")
	pterm.DefaultCenter.Println(pterm.Cyan(pterm.DefaultHeader.WithFullWidth().Sprint("Termux Best Mirror")))
	pterm.Println()

	selectedMirrorDisplay, _ := pterm.DefaultInteractiveSelect.
		WithDefaultText("Select Your Preferred Mirror").
		WithOptions(resultOptions).
		WithMaxHeight(12).
		Show()

	selectedMirror := resultMap[selectedMirrorDisplay]
	pterm.Success.Printf("Selected: %s\n", pterm.Cyan(selectedMirror.Name))

	system.ApplyMirror(selectedMirror.Path)
}
