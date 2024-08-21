package commands

import (
	"context"

	"github.com/bep/simplecobra"
	"github.com/spf13/cobra"
)

type comando struct {
	use   string
	name  string
	short string
	long  string
	run   func(ctx context.Context, cmder *simplecobra.Commandeer, rootCmd *rootComando, args []string) error
	withc func(cmd *cobra.Command, r *rootComando)
	initc func(cmder *simplecobra.Commandeer) error

	comandos []simplecobra.Commander
	rootCmd  *rootComando
}

func (c *comando) Commands() []simplecobra.Commander {
	return c.comandos
}

func (c *comando) Name() string {
	return c.name
}

func (c *comando) PreRun(cmd, runner *simplecobra.Commandeer) error {
	if c.initc != nil {
		return c.initc(cmd)
	}

	return nil
}

func (c *comando) Run(ctx context.Context, cmder *simplecobra.Commandeer, args []string) error {
	if c.run == nil {
		return nil
	}

	return c.run(ctx, cmder, c.rootCmd, args)
}

func (c *comando) Init(cmder *simplecobra.Commandeer) error {
	c.rootCmd = cmder.Root.Command.(*rootComando)

	cmd := cmder.CobraCommand
	cmd.Short = c.short
	cmd.Long = c.long
	if c.use != "" {
		cmd.Use = c.use
	}

	if c.withc != nil {
		c.withc(cmd, c.rootCmd)
	}

	return nil
}
