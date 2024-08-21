package commands

import (
	"context"

	"github.com/bep/simplecobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newVersionCmd() simplecobra.Commander {
	return &comando{
		name:  "version",
		short: "show texcoco version",
		long:  "show texcoco version",
		run: func(ctx context.Context, cmder *simplecobra.Commandeer, rootCmd *rootComando, args []string) error {
			rootCmd.Printf("v%s\n", viper.GetString("version"))

			return nil
		},
		withc: func(cmd *cobra.Command, r *rootComando) {
			cmd.ValidArgsFunction = cobra.NoFileCompletions
		},
	}
}
