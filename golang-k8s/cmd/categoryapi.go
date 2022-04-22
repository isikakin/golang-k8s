package cmd

import (
	"github.com/spf13/cobra"
	"golang-k8s/api/categoryapi"
)

var categoryApiCmd = &cobra.Command{
	Use:   "categoryapi",
	Short: "category api",
	Long:  `category api`,
	RunE:  categoryapi.Init,
}

func init() {
	RootCmd.AddCommand(categoryApiCmd)
}
