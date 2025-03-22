package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"poridhictl/internal/ignite"

	"github.com/spf13/cobra"
)

var (
	nodeName  string
	cpus      int
	memory    string
	diskSize  string
	imageOCI  string
	enableSSH bool
)

func generateUID() (string, error) {
	bytes := make([]byte, 8) // 8 bytes = 16 hex characters
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate UID: %v", err)
	}
	return hex.EncodeToString(bytes), nil
}

var launchCmd = &cobra.Command{
	Use:   "launch vm",
	Short: "Launch a new VM",
	Long:  `Launch a new virtual machine with specified configuration using Weave Ignite`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if args[0] != "vm" {
			return fmt.Errorf("only 'vm' argument is supported")
		}

		// Generate unique ID for the node
		nodeUID, err := generateUID()
		if err != nil {
			return fmt.Errorf("failed to generate UID: %v", err)
		}

		// Create VM configuration
		vmConfig := &ignite.VMConfig{
			Name:      nodeName,
			UID:       nodeUID,
			CPUs:      cpus,
			Memory:    memory,
			DiskSize:  diskSize,
			ImageOCI:  imageOCI,
			EnableSSH: enableSSH,
		}

		// Launch the VM
		ip, err := ignite.LaunchVM(vmConfig)
		if err != nil {
			return fmt.Errorf("failed to launch VM: %v", err)
		}

		fmt.Printf("Successfully launched VM:\n")
		fmt.Printf("Name: %s\n", nodeName)
		fmt.Printf("ID: %s\n", nodeUID)
		fmt.Printf("IP: %s\n", ip)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)

	// Add flags
	launchCmd.Flags().StringVarP(&nodeName, "name", "n", "", "Name of the VM (required)")
	launchCmd.Flags().IntVarP(&cpus, "cpus", "c", 2, "Number of CPUs")
	launchCmd.Flags().StringVarP(&memory, "memory", "m", "1GB", "Memory size")
	launchCmd.Flags().StringVarP(&diskSize, "disk-size", "d", "3GB", "Disk size")
	launchCmd.Flags().StringVarP(&imageOCI, "image", "i", "shajalahamedcse/only-k3-go:v1.0.10", "OCI image")
	launchCmd.Flags().BoolVarP(&enableSSH, "ssh", "s", true, "Enable SSH")

	// Mark required flags
	launchCmd.MarkFlagRequired("name")
}
