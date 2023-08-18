package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/eventlog"
)

var keywordCmd = &cobra.Command{
	Use:   "keyword [keywords...]",
	Short: "Search Sysmon logs for keywords",
	Args:  cobra.MinimumNArgs(1),
	Run:   searchLogsForKeywords,
}

func main() {
	rootCmd := &cobra.Command{Use: "sysmonlogs"}
	rootCmd.AddCommand(keywordCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func searchLogsForKeywords(cmd *cobra.Command, args []string) {
	keywords := args

	// Open the System event log
	el, err := eventlog.Open("Microsoft-Windows-Sysmon/Operational")
	if err != nil {
		log.Fatal(err)
	}
	defer el.Close()

	// Read the event log records
	records, err := el.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	events := []map[string]interface{}{}
	for _, record := range records {
		// Skip event log errors
		if record.EventType == windows.EVENTLOG_ERROR_TYPE {
			continue
		}

		// Parse the event data
		eventData, err := record.String()
		if err != nil {
			log.Println("Error parsing event data:", err)
			continue
		}

		event := map[string]interface{}{
			"Time":    record.TimeGenerated.String(),
			"EventID": record.EventID,
			"Source":  record.Source,
			"Message": eventData,
		}

		// Check if any keyword matches the event message
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(eventData), strings.ToLower(keyword)) {
				events = append(events, event)
				break
			}
		}
	}

	// Print the matching log entries
	for _, event := range events {
		fmt.Println(event)
	}
}
