package cmd

import (
	"github.com/spf13/cobra"
)

// showCmd represents the show command
func newShowCmd() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "short comment for show subcommand",
		Long:  "long comment for show subcommand",
		RunE: func(cmd *cobra.Command, args []string) error {
			i, err := cmd.Flags().GetInt("integer")

			if err != nil {
				return err
			}

			b, err := cmd.Flags().GetBool("boolean")

			if err != nil {
				return err
			}

			s, err := cmd.Flags().GetString("string")

			if err != nil {
				return err
			}

			cui.Outputln("Integer option value:", i)
			cui.Outputln("Boolean option value:", b)
			cui.Outputln(" String option value:", s)

			return nil
		},
	}

	showCmd.Flags().IntP("integer", "i", 0, "integer option")
	showCmd.Flags().BoolP("boolean", "b", false, "boolean option")
	showCmd.Flags().StringP("string", "s", "", "string option")

	return showCmd
}
