/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pos-goexpert-desafio-clean-arch",
	Short: "Clean Architecture",
	Long:  `Desafio sobre clean architecture da PÓS gradução go-expert.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	host, _ := rootCmd.Flags().GetString("db-host")
	port, _ := rootCmd.Flags().GetInt("db-port")
	user, _ := rootCmd.Flags().GetString("db-user")
	pass, _ := rootCmd.Flags().GetString("db-password")
	dbName, _ := rootCmd.Flags().GetString("db-name")

	db, err := database.NewDB(host, port, user, pass, dbName)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %+v", err)
	}
	ctx := context.WithValue(context.Background(), database.DBContextKey, db)

	err = rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().String("db-host", "localhost", "MySQL server hostname")
	rootCmd.Flags().Int("db-port", 3306, "MySQL server port")
	rootCmd.Flags().String("db-name", "goexpert", "MySQL server database name")
	rootCmd.Flags().String("db-user", "root", "MySQL server username")
	rootCmd.Flags().String("db-password", "root", "MySQL server user password")
}
