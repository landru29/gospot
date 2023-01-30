package main

import (
	"context"
	"os"
	"strings"

	"github.com/landru29/gospot/internal/app"
	"github.com/spf13/cobra"
)

func rootCommand(ctx context.Context, quitSignal chan os.Signal) *cobra.Command {
	var (
		configurationFolder string
		listToLaunch        []string
	)

	conf := &app.Config{}

	cmd := &cobra.Command{
		Use:   "gospot",
		Short: "gospot",
		Long:  "gospot main command",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return conf.LoadConfig(configurationFolder, "conf")
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			application, cleanup, err := setupApplication(conf, apiList(listToLaunch))
			if err != nil {
				return err
			}

			application.Start(ctx, quitSignal)

			cleanup()

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(
		&configurationFolder,
		"configuration-folder",
		"c",
		".",
		"configuration folder")

	cmd.Flags().StringSliceVarP(
		&listToLaunch,
		"launch",
		"l",
		[]string{"client", "api"},
		"servers to launch ['client', 'api']",
	)

	return cmd
}

type apiList []string

func (a apiList) Has(elt string) bool {
	for _, inside := range a {
		if strings.EqualFold(inside, elt) {
			return true
		}
	}

	return false
}
