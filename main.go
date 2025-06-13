// fgo - Find latest compatible Go module version for your Go version
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: fgo install/get {module-path}")
		os.Exit(1)
	}

	packageName := os.Args[2]

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
	command := os.Args[1]

	for _, version := range versions {
		full := fmt.Sprintf("%s@%s", packageName, version)
		cmd := exec.Command("go", command, full)
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
