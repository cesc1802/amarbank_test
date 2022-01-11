package cmd

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use:   "amarbank",
	Short: "",
	Long:  "",
}

func init() {
	root.AddCommand(serverCmd)

	serverCmd.Flags().StringP("mode", "m", "debug", "gin server mode")
	serverCmd.Flags().StringP("host", "a", "localhost", "")
	serverCmd.Flags().StringP("port", "p", "8088", "")
}

func Execute() error {
	if err := root.Execute(); err != nil {
		return err
	}
	return nil
}
