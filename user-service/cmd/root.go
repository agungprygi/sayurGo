package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/rs/zerolog/log"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "user-service",
	Short: "User service",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Run(startCmd, nil)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .env)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".env")
		viper.SetConfigType("env")
	}
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Err(err).Msg("Error reading config file: " + viper.ConfigFileUsed())
	}
}
