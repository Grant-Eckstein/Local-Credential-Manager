package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func init() {
	Config.AddCommand(New)
}

var (
	New = &cobra.Command{
		Use:   "save [profile]",
		Args:  cobra.ExactArgs(1),
		Short: "Save a new profile",
		Long:  "Save a new profile",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]

			if err := viper.ReadInConfig(); err == nil {
				if debug {
					log.Println("Successfully read in existing configuration file: ", viper.ConfigFileUsed())
				}
			} else {
				log.Fatalln("Error reading viper config: ", err)
			}
			// Check if file exists
			if h, err := os.UserHomeDir(); err == nil {
				config_file := filepath.Join(h, ".ssh", "config")

				if debug {
					log.Printf("Checking SSH Config at: %s\n", config_file)
				}

				if _, err := os.Stat(config_file); err == nil {
					log.Printf("File exists, creating SSH profile entry\n")

					if data, err := ioutil.ReadFile(config_file); err == nil {
						viper.Set(name, string(data))
						if err := viper.WriteConfig(); err == nil {
							log.Println("Successfully wrote new profile!")
						} else {
							log.Fatalln("Error writing to profiler: ", err)
						}
					} else {
						log.Fatalln("Error reading ssh config: ", err)
					}
				}
			} else {
				log.Fatalln("Error finding ssh directory: ", err)
			}

		},
	}
)
