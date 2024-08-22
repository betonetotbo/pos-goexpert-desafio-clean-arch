/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "Faz a migração da base de dados",
	Long:    `Faz a migração da base de dados criando as tabelas necessárias`,
	PreRunE: database.PrepareDatabase,
	RunE: func(cmd *cobra.Command, args []string) error {
		return database.RunMigrations(cmd.Context())
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
