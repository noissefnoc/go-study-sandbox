package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"os"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cui = rwi.New() // CUI instance
	cfgFile string
)

func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cui = ui

	rootCmd := &cobra.Command{
		Use:   "cobra",
		Short: "short comment",
		Long: "long comment",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}

	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	rootCmd.AddCommand(newShowCmd())

	return rootCmd
}

func Execute(cui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		// panic handling
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)

				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	// execution
	exit = exitcode.Normal

	if err := newRootCmd(cui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}

	return
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
