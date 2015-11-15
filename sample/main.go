package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "boom",
		Short: "make an explosive entrance",
		RunE: func(c *cobra.Command, args []string) error {
			println("make an explosive entrance")
			return nil
		},
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
