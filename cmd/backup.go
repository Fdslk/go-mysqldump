package cmd

import (
	"github.com/dumpsql/mysqldump"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(DBBackupcmd)
}

var DBBackupcmd = &cobra.Command{
	Use:   "backup",
	Short: "backup database dumps",
	Long:  "create mysql dumps for devloping based on local environment",
	Run: func(cmd *cobra.Command, args []string) {
		mysqldump.DBBackup()
	},
}
