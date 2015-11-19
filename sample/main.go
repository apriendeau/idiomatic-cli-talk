package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "boom",
	Short:         "make an explosive entrance",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(c *cobra.Command, args []string) error {
		// oh no! lets fake a error
		return errors.New("BOOM!! I SAY!")
	},
}

func main() {
	if cmd, err := RootCmd.ExecuteC(); err != nil {
		fmt.Println(cmd.Usage())
		log.Fatal(err) // <- all errors come here
	}
}
