package key

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func init() {
	Key.AddCommand(Load)
}

var (
	Load = &cobra.Command{
		Use:   "load [profile name]",
		Args:  cobra.ExactArgs(1),
		Short: "Write existing profile to current ssh configuration",
		Long:  "Write existing profile to current ssh configuration",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			if err := viper.ReadInConfig(); err == nil {
				if h, err := os.UserHomeDir(); err == nil {
					private_key_file := filepath.Join(h, ".ssh", "id_rsa")
					public_key_file := filepath.Join(h, ".ssh", "id_rsa.pub")
					auth_keys := filepath.Join(h, ".ssh", "authorized_keys")
					known_hosts := filepath.Join(h, ".ssh", "known_hosts")
					profile_contents := viper.GetString(name)
				} else {
					log.Fatalln(err)
				}
			} else {
				log.Fatalln(err)
			}
		},
	}
)
