package config

import (
	"context"
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
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

	path, _ := filepath.Abs(".")
	log.Printf("Loading config from: %s\n", path)

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
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

	err = check(cfg)
	if err != nil {
		return err
	}

	ctx := context.WithValue(cmd.Context(), configContextKey, cfg)
	cmd.SetContext(ctx)
	return nil
}

func check(cfg *Configuration) error {
	if cfg.DBHost == "" {
		return errors.New("DB_HOST is required")
	}
	if cfg.DBPort == 0 {
		return errors.New("DB_PORT is required")
	}
	if cfg.DBUser == "" {
		return errors.New("DB_USER is required")
	}
	if cfg.DBPassword == "" {
		return errors.New("DB_PASSWORD is required")
	}
	if cfg.DBName == "" {
		return errors.New("DB_NAME is required")
	}
	if cfg.WebServerPort < 1024 || cfg.WebServerPort > 65535 {
		return errors.New("WEB_SERVER_PORT must be between 49152 and 65535")
	}
	if cfg.GRPCServerPort < 49152 || cfg.GRPCServerPort > 65535 {
		return errors.New("GRPC_SERVER_PORT must be between 49152 and 65535")
	}
	return nil
}

func GetConfig(ctx context.Context) *Configuration {
	cfg, ok := ctx.Value(configContextKey).(*Configuration)
	if !ok {
		log.Fatalf("Failed to get config from context")
	}
	return cfg
}
