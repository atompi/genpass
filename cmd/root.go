/*
Copyright © 2023 Atom Pi <coder.atompi@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/atompi/genpass/cmd/options"
	"github.com/atompi/genpass/pkg/handle"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "genpass [length]",
	Short: "A tool for generating random passwords",
	Long: `A tool for generating random passwords.
It basically conforms to the OWASP strong password specification.`,
	Version: options.Version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		opts := options.NewOptions(args)
		handle.Handle(opts)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().BoolP("all", "a", false, "with all types, represent -lns")
	rootCmd.PersistentFlags().BoolP("letter", "l", false, "with letters")
	rootCmd.PersistentFlags().BoolP("number", "n", false, "with numbers")
	rootCmd.PersistentFlags().BoolP("symbol", "s", false, "with symbols")

	viper.BindPFlag("all", rootCmd.PersistentFlags().Lookup("all"))
	viper.BindPFlag("letter", rootCmd.PersistentFlags().Lookup("letter"))
	viper.BindPFlag("number", rootCmd.PersistentFlags().Lookup("number"))
	viper.BindPFlag("symbol", rootCmd.PersistentFlags().Lookup("symbol"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {}
