// Copyright © 2018 NAME HERE shane.siebken@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/shanesiebken/csv-util/csv"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "csv-util",
	Short: "A collection of .csv file utilities",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		pwd, err := os.Executable()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		log.Print(pwd)
		// Search config in home directory or current directory with name "config" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(pwd)
		viper.AddConfigPath("./")

		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run() {
	records, err := csv.Read(viper.GetString("source"))
	if err != nil {
		log.Fatalf("Unable to read source file: %s", err)
	}
	records, err = csv.DoMapping(records)
	if err != nil {
		log.Fatalf("Error performing field mapping: %s", err)
	}
	records, err = csv.DoConcatenation(records)
	if err != nil {
		log.Fatalf("Error performing field concatenation: %s", err)
	}
	records = csv.DoDrop(records)
	err = csv.Write(viper.GetString("destination"), records)
	if err != nil {
		log.Fatalf("Error writing destination file: %s", err)
	}
}
