package main

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// TestSysmonLogsParsing tests the Sysmon logs parsing functionality
func TestSysmonLogsParsing(t *testing.T) {
	// Test keywords
	keywords := []string{"error", "warning"}

	// Set up command-line arguments
	os.Args = []string{"sysmonlogs", "keyword", "error", "warning"}

	// Capture stdout for testing
	var output bytes.Buffer
	log.SetOutput(&output)

	// Execute the command
	cmd := &cobra.Command{}
	searchLogsForKeywords(cmd, keywords)

	// Verify output contains expected log entries
	outputStr := output.String()
	for _, keyword := range keywords {
		if !strings.Contains(outputStr, keyword) {
			t.Errorf("Expected log entry containing keyword '%s' not found in output", keyword)
		}
	}
}

// TestMain runs the test suite
func TestMain(m *testing.M) {
	// Run tests
	exitCode := m.Run()

	// Exit with the appropriate code
	os.Exit(exitCode)
}
