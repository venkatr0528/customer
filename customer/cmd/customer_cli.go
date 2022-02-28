package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{

	Use: "start",

	Short: "This command will show you a hello world message",

	Long: `Welcome in start command, this cmd will display to you a hello world message.`,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Hello world, start have been called!")

	},
}
