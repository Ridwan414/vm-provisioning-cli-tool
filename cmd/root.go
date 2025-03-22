package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "poridhictl",
	Short: "poridhictl is a CLI tool for managing VM provisioning",
	Long: `A CLI tool that allows you to provision and manage virtual machines 
using Weave Ignite through a simple command line interface.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
