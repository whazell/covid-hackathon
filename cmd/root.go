package main

import (
	"covid/cmd/api"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "covid",
		Short: "Covid-19 Business Tracker API & CLI tool",
	}
	rootCmd.AddCommand(api.NewCommand())
	return rootCmd
}

func main() {
	if err := NewCommand().Execute(); err != nil {
		fmt.Printf("Could not run covid: %s\n", err.Error())
		os.Exit(1)
	}
}
