package system

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pterm/pterm"
)

var (
	Prefix     string
	MirrorsDir string
)

func init() {
	Prefix = os.Getenv("TERMUX_PREFIX")
	if Prefix == "" {
		Prefix = "/data/data/com.termux/files/usr"
	}
	MirrorsDir = filepath.Join(Prefix, "etc", "termux", "mirrors")
}

func CheckEnvironment() {
	if _, err := os.Stat(MirrorsDir); os.IsNotExist(err) {
		pterm.Error.Printf("This script must be run inside Termux.\nCould not locate mirrors directory at: %s\n", MirrorsDir)
		os.Exit(1)
	}

	if _, err := exec.LookPath("apt"); err != nil {
		pterm.Error.Println("Cannot change mirrors since apt is not installed.")
		os.Exit(1)
	}

	pkgManager := os.Getenv("TERMUX_APP_PACKAGE_MANAGER")
	if pkgManager == "pacman" {
		pterm.Warning.Println("This script only changes mirrors for apt.")
		result, _ := pterm.DefaultInteractiveConfirm.WithDefaultText("Do you want to continue?").Show()
		if !result {
			os.Exit(0)
		}
	}
}

func ApplyMirror(path string) {
	spinner, _ := pterm.DefaultSpinner.
		WithSequence("⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷").
		Start("Updating chosen_mirrors symbolic link...")
	
	chosenPath := filepath.Join(Prefix, "etc", "termux", "chosen_mirrors")

	if _, err := os.Lstat(chosenPath); err == nil {
		os.Remove(chosenPath)
	}

	err := os.Symlink(path, chosenPath)
	if err != nil {
		spinner.Fail(fmt.Sprintf("Failed to link mirror: %v", err))
		os.Exit(1)
	}
	spinner.Success("Symbolic link updated.")

	pterm.Info.Println("Verifying configuration and updating Termux package indexes...")

	cmd := exec.Command("pkg", "--check-mirror", "update")
	cmd.Env = append(os.Environ(), "TERMUX_APP_PACKAGE_MANAGER=apt")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		pterm.Error.Printf("Failed to update packages: %v\n", err)
		os.Exit(1)
	}

	pterm.Println()
	pterm.DefaultBox.
		WithTitle("All Done!").
		WithTitleTopCenter().
		WithRightPadding(2).
		WithLeftPadding(2).
		Println(pterm.Green("Repository mirror has been updated successfully.\nTermux packages are ready to go!"))
}
