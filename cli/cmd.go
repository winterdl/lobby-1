package cli

import (
	"log"
	"os"
	"path"

	"github.com/BurntSushi/toml"
	"github.com/asdine/lobby/cli/app"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// New returns the lobby CLI application.
func New() *cobra.Command {
	app := app.NewApp()
	cmd := newRootCmd(app)
	cmd.AddCommand(newRunCmd(app))
	return cmd.Command
}

func newRootCmd(app *app.App) *rootCmd {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}

	defaultConfigDir := path.Join(home, ".config/lobby")

	var cfgMeta toml.MetaData

	cmd := cobra.Command{
		Use:          "lobby",
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			configPath := path.Join(app.Config.Paths.ConfigDir, "lobby.toml")
			f, err := os.Open(configPath)
			if err == nil {
				cfgMeta, err = toml.DecodeReader(f, &app.Config)
				_ = f.Close()
				if err != nil {
					return err
				}
			}

			if app.Config.Paths.SocketDir == "" {
				app.Config.Paths.SocketDir = path.Join(app.Config.Paths.ConfigDir, "sockets")
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&app.Config.Paths.ConfigDir, "config-dir", "c", defaultConfigDir, "Location of Lobby configuration files")

	return &rootCmd{
		Command: &cmd,
		cfgMeta: &cfgMeta,
	}
}

type rootCmd struct {
	*cobra.Command
	cfgMeta *toml.MetaData
}
