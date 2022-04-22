package cmd

import (
	"github.com/spf13/cobra"
	"golang-k8s/api/productapi"
)

var productApiCmd = &cobra.Command{
	Use:   "productapi",
	Short: "product api",
	Long:  `product api`,
	RunE:  productapi.Init,
}

func init() {
	RootCmd.AddCommand(productApiCmd)
}
