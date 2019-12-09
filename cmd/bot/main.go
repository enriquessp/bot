package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/what-is-bot/bot/internal"

	"github.com/spf13/cobra"
	"github.com/what-is-bot/bot/internal/postgres"
	"github.com/what-is-bot/bot/internal/slack"
)

var cmd = &cobra.Command{
	Use:   "bot",
	Short: "What is bot is your friendly onboarding helper for Slack",
	Run: func(cmd *cobra.Command, args []string) {
		slack.Start(internal.NewController(postgres.NewAnswerProvider(), nil))
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.SetEnvPrefix("WHAT_IS_BOT")
	viper.AutomaticEnv()
}
