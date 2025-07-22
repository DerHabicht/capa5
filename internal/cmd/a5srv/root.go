package a5srv

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/capa5/internal/a5srv"
	"github.com/derhabicht/capa5/internal/config"
	"github.com/derhabicht/capa5/internal/logging"
)

var logLevel string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: viper.GetString("version"),
	Use:     "a5srv",
	Short:   "Run the CAP/A5 server",
	Long:    ``,
	Run:     runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	r, err := a5srv.Setup()
	if err != nil {
		logging.Fatal().Err(err).Stack().Msg("failed to configure router")
	}

	err = r.Run()
	if err != nil {
		logging.Fatal().Err(err).Stack().Msg("failed to start server")
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	config.Set(config.Version, version)
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		logging.Fatal().Err(err).Stack().Msg("Error encountered while running a5srv")
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logLevel, "loglevel", "info", "")

	logging.InitLogging(logLevel, true)
}
