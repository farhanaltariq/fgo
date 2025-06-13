// fgoinstall - Find latest compatible Go module version for your Go version
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: fgoinstall {module-path}")
		os.Exit(1)
	}

	packageName := os.Args[1]

	fmt.Println("📦 Fetching versions for:", packageName)

	// Get all versions
	cmd := exec.Command("go", "list", "-m", "-versions", packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ Error listing versions:", err)
		fmt.Println(string(output))
		return
	}

	// Parse versions
	parts := strings.Fields(string(output))
	if len(parts) < 2 {
		fmt.Println("⚠️ No versions found.")
		return
	}
	versions := parts[1:]

	// Sort newest to oldest
	sort.Slice(versions, func(i, j int) bool {
		return strings.Compare(versions[i], versions[j]) > 0
	})

	fmt.Println("🔍 Testing versions for compatibility with your Go version...")

	for _, version := range versions {
		full := fmt.Sprintf("%s@%s", packageName, version)
		cmd := exec.Command("go", "install", full)
		cmd.Stderr = new(bytes.Buffer)
		cmd.Stdout = new(bytes.Buffer)

		err := cmd.Run()
		if err == nil {
			fmt.Printf("✅ Compatible: %s\n", full)
			return
		}
		fmt.Printf("❌ Not compatible: %s\n", full)
	}

	fmt.Println("❌ No compatible version found.")
}

func getGoVersion() (major int, minor int, err error) {
	cmd := exec.Command("go", "version")
	output, err := cmd.Output()
	if err != nil {
		return 0, 0, fmt.Errorf("failed to get go version: %w", err)
	}

	// Example: go version go1.11 linux/amd64
	parts := strings.Fields(string(output))
	if len(parts) < 3 {
		return 0, 0, fmt.Errorf("unexpected output: %s", output)
	}

	verStr := strings.TrimPrefix(parts[2], "go")
	verParts := strings.Split(verStr, ".")
	if len(verParts) < 2 {
		return 0, 0, fmt.Errorf("unexpected version format: %s", verStr)
	}

	major, err = strconv.Atoi(verParts[0])
	if err != nil {
		return 0, 0, err
	}
	minor, err = strconv.Atoi(verParts[1])
	if err != nil {
		return 0, 0, err
	}
	return major, minor, nil
}

func tryInstall(pkg, version string) error {
	full := fmt.Sprintf("%s@%s", pkg, version)

	major, minor, err := getGoVersion()
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	var cmdStr string

	if major == 1 && minor < 17 {
		// Use go get for versions < 1.17
		cmdStr = fmt.Sprintf("go get %s", full)
		cmd = exec.Command("go", "get", full)
	} else {
		// Use go install for 1.17+
		cmdStr = fmt.Sprintf("go install %s", full)
		cmd = exec.Command("go", "install", full)
	}

	fmt.Println("👉 Running:", cmdStr)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("command failed: %v\nstderr: %s", err, stderr.String())
	}
	return nil
}

