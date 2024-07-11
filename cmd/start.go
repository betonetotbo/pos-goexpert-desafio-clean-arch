/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Inicia a aplicação",
	Long:  `Inicia a aplicação levantando o servidor HTTP`,
	RunE: func(cmd *cobra.Command, args []string) error {
		httpPort, e := cmd.Flags().GetInt("port")
		if e != nil {
			return e
		}
		grpcPort, e := cmd.Flags().GetInt("grpc")
		if e != nil {
			return e
		}
		server.ListenAndServeHTTP(httpPort, grpcPort)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("port", "p", 8080, "Porta do servidor HTTP")
	startCmd.Flags().IntP("grpc", "g", 50000, "Porta do servidor gRPC")
}
