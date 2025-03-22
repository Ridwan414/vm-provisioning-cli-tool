package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all VMs",
	Long:  `Display all running virtual machines in a tabular format`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Run ignite ps command
		output, err := exec.Command("sudo", "ignite", "ps").Output()
		if err != nil {
			return fmt.Errorf("failed to list VMs: %v", err)
		}

		// Split output into lines
		lines := strings.Split(string(output), "\n")

		// Create table data
		tableData := [][]string{
			// Header
			{"VM ID", "NAME", "IMAGE", "KERNEL", "SIZE", "CPUS", "MEMORY", "CREATED", "STATUS", "IP"},
		}

		// Add VM data
		for i, line := range lines {
			// Skip header and empty lines
			if i == 0 || len(strings.TrimSpace(line)) == 0 {
				continue
			}

			// Parse the line into fields
			parts := strings.Fields(line)
			if len(parts) < 14 {
				continue // Skip malformed lines
			}

			vmData := []string{
				parts[0],                  // VM ID
				parts[len(parts)-1],       // NAME
				parts[1],                  // IMAGE
				parts[2],                  // KERNEL
				parts[3] + " " + parts[4], // SIZE
				parts[5],                  // CPUS
				parts[6] + " " + parts[7], // MEMORY
				parts[8] + " " + parts[9], // CREATED
				parts[10],                 // STATUS
				parts[12],                 // IPS
			}
			tableData = append(tableData, vmData)
		}

		// Create and configure table
		table := pterm.TableData(tableData)

		// Print table with styling
		err = pterm.DefaultTable.
			WithHasHeader().
			WithBoxed(true).
			WithHeaderStyle(pterm.NewStyle(pterm.FgLightCyan)).
			WithData(table).
			Render()

		if err != nil {
			return fmt.Errorf("failed to render table: %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
