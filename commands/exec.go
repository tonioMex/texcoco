package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/bep/simplecobra"
)

func newExec() (*simplecobra.Exec, error) {
	rootCmd := &rootComando{
		comandos: []simplecobra.Commander{
			newEnvCmd(),
			newVersionCmd(),
		},
	}

	return simplecobra.New(rootCmd)
}

func Execute(args []string) error {
	exec, err := newExec()
	if err != nil {
		return err
	}

	cmder, err := exec.Execute(context.Background(), args)
	if err != nil {
		if err == errors.New("help requested") {
			cmder.CobraCommand.Help()
			fmt.Println()
			return nil
		}

		if simplecobra.IsCommandError(err) {
			cmder.CobraCommand.Help()
			fmt.Println()
		}
	}

	return err
}
