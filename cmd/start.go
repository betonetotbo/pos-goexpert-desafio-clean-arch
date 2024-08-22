/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/config"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/server"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/utils"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia a aplicação",
	Long:  `Inicia a aplicação levantando o servidor HTTP`,
	PreRunE: utils.MultiRun(
		database.PrepareDatabase,
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.GetConfig(cmd.Context())
		server.ListenAndServe(cmd.Context(), cfg.WebServerPort, cfg.GRPCServerPort)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
