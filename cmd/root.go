/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"cliTest/config"
	"cliTest/events"
	sentryService "cliTest/services/sentry"
	"github.com/spf13/cobra"
)

var (
	ProductId     int
	UserId        int
	configuration config.Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "buy_event <send_type>",
	Short:   configuration.Description,
	Run:     events.Actions,
	Args:    cobra.MinimumNArgs(1),
	Version: configuration.Version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	configuration = *config.GetInstance()
	sentryService.SentryService.Init()
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().IntVarP(&UserId, "user_id", "u", UserId, "User id")
	rootCmd.Flags().IntVarP(&ProductId, "product_id", "p", ProductId, "Product id")

}
