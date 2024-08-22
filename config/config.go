package config

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type (
	configContextKeyValue string

	Configuration struct {
		DBHost         string `mapstructure:"DB_HOST"`
		DBPort         int    `mapstructure:"DB_PORT"`
		DBUser         string `mapstructure:"DB_USER"`
		DBPassword     string `mapstructure:"DB_PASSWORD"`
		DBName         string `mapstructure:"DB_NAME"`
		WebServerPort  int    `mapstructure:"WEB_SERVER_PORT"`
		GRPCServerPort int    `mapstructure:"GRPC_SERVER_PORT"`
	}
)

var configContextKey configContextKeyValue = "config-context-key"

func LoadConfig(cmd *cobra.Command, _ []string) error {
	var cfg *Configuration
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	ctx := context.WithValue(cmd.Context(), configContextKey, cfg)
	cmd.SetContext(ctx)
	return nil
}

func GetConfig(ctx context.Context) *Configuration {
	return ctx.Value(configContextKey).(*Configuration)
}
