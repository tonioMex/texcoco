package commands

import (
	"context"
	"runtime"

	"github.com/bep/simplecobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newEnvCmd() simplecobra.Commander {
	return &comando{
		name:  "env",
		short: "display texcoco environment variables",
		long:  "display texcoco environment variables",
		run: func(ctx context.Context, cmder *simplecobra.Commandeer, rootCmd *rootComando, args []string) error {
			rootCmd.Printf("GOES=%q\n", runtime.GOOS)
			rootCmd.Printf("GOARCH=%q\n", runtime.GOARCH)
			rootCmd.Printf("GOVERSION=%q\n", runtime.Version())
			rootCmd.Println()
			rootCmd.Println("--- config.yaml ---")

			for _, key := range viper.AllKeys() {
				rootCmd.Printf("%s: %v\n", key, viper.Get(key))
			}

			return nil
		},
		withc: func(cmd *cobra.Command, r *rootComando) {
			cmd.ValidArgsFunction = cobra.NoFileCompletions
		},
	}
}
