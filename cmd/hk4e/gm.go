package main

import (
	"context"

	"hk4e/gm/app"

	"github.com/spf13/cobra"
)

// GMCmd
func GMCmd() *cobra.Command {
	var cfg string
	c := &cobra.Command{
		Use:   "gm",
		Short: "gm server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.Run(context.Background(), cfg)
		},
	}
	c.Flags().StringVar(&cfg, "config", "application.toml", "config file")
	return c
}
