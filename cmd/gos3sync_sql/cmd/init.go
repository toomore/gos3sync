package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/toomore/gos3sync/syncdb"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "init a new sync folder",
	Long:  `init a new sync folder, create db.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called", args, cmd.Flags().Lookup("name").Value)
		if len(args) == 1 {
			log.Println(cmd.CommandPath())
			syncdb.Init(args[0])
		}
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().BoolP("name", "n", false, "Need name")

}
