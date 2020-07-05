package config

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
	Config.AddCommand(Load)
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
					config_file := filepath.Join(h, ".ssh", "config")
					profile_contents := viper.GetString(name)
					if err := ioutil.WriteFile(config_file, []byte(profile_contents), 500); err == nil {
						fmt.Println("Successfully loaded new profile!")
					} else {
						log.Fatalln(err)
					}
				} else {
					log.Fatalln(err)
				}
			} else {
				log.Fatalln(err)
			}
		},
	}
)
