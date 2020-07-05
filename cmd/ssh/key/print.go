package key

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Key.AddCommand(Display)
}

var (
	Display = &cobra.Command{
		Use:   "print [profile]",
		Args:  cobra.ExactArgs(1),
		Short: "Print out configuration file",
		Long:  "Print out configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			cfg := viper.Get(name)

			fmt.Printf("%s\n", cfg)
		},
	}
)
