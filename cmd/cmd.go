package cmd

import (
	"net"

	"github.com/spf13/cobra"

	"github.com/drakejin/fiber-aws-serverless/config"
	_const "github.com/drakejin/fiber-aws-serverless/const"
	"github.com/drakejin/fiber-aws-serverless/db"
	"github.com/drakejin/fiber-aws-serverless/internal/app"
	"github.com/drakejin/fiber-aws-serverless/internal/container"
)

func Start(cfg *config.Config) error {
	c := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	c.AddCommand(HTTPCommand(cfg))
	c.AddCommand(GORMCommand(cfg))
	return c.Execute()
}

func HTTPCommand(cfg *config.Config) *cobra.Command {
	c := &cobra.Command{
		Use:     "http",
		Aliases: []string{"h"},
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(&cobra.Command{
		Use:     "start",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			cont, err := container.New(cfg)
			if err != nil {
				return err
			}

			appHttp := app.NewHTTP(cont)
			if err := appHttp.Listen(net.JoinHostPort("0.0.0.0", cont.Config.HTTPServer.Port)); err != nil {
				return err
			}
			return nil
		},
	})

	return c
}

func GORMCommand(cfg *config.Config) *cobra.Command {
	c := &cobra.Command{
		Use:     "gorm",
		Aliases: []string{"g"},
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	c.AddCommand(&cobra.Command{
		Use:     "init",
		Aliases: []string{"i"},
		RunE: func(cmd *cobra.Command, _ []string) error {

			if err := db.Clean(cfg.Env, cfg.ServiceDB); err != nil {
				if cfg.Env != _const.EnvAlpha {
					return err
				}
			}
			if err := db.Initialize(cfg.ServiceDB); err != nil {
				return err
			}
			if cont, err := container.New(cfg); err != nil {
				return err
			} else {
				return cont.ServiceDB.InitMigrator()
			}
		},
	})
	return c
}