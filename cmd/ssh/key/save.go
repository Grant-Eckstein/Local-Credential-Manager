package key

import (
	"github.com/spf13/cobra"
)

func init() {
	Key.AddCommand(New)
}

var (
	New = &cobra.Command{
		Use:   "save [profile]",
		Args:  cobra.ExactArgs(1),
		Short: "Save a new profile",
		Long:  "Save a new profile",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)
