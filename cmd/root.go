package main

import (
	"covid/cmd/api"
	"covid/cmd/create"
	"covid/cmd/process"
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
	rootCmd.AddCommand(create.NewCommand())
	rootCmd.AddCommand(process.NewCommand())
	return rootCmd
}

func main() {
	if err := NewCommand().Execute(); err != nil {
		fmt.Printf("Could not run covid: %s\n", err.Error())
		os.Exit(1)
	}
}
