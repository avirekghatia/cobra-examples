package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var printRootCmd = &cobra.Command{
	Use:   "printcobra",
	Short: "printcobra is a sample command in this experiment with Cobra",
	Long: `A sample cli command application to demonstrate
			the use of Cobra and for learning`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("In the run command in root.go")
	},
}

func Execute() {
	fmt.Println("inside the execute func")
	if err := printRootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
