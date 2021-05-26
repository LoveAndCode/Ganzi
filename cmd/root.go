package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ganzi",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Hello command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Echo Argument parameter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Ganzi! args: => [" + strings.Join(args, " , ") + "]")
	},
}

// Message Command
var message = &cobra.Command{
	Use:   "message",
	Short: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		text := strings.Join(args, "")
		homePath, err := homedir.Dir()

		if err != nil {
			panic(err)
		}

		textFileName := "banner.txt"

		fo, err := os.Create(homePath + "/" + textFileName)

		if err != nil {
			panic(err)
		}
		defer fo.Close()

		_, err = fo.WriteString(text)

		if err != nil {
			panic(err)
		}

		// open profile
		profile, err := os.OpenFile(homePath+"/.profile", os.O_APPEND|os.O_WRONLY, 0644)

		defer profile.Close()

		if err != nil {
			panic(err)
		}
		_, err = profile.WriteString("cat " + homePath + "/" + textFileName)
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ganzi.yaml)")

	rootCmd.AddCommand(message)
	rootCmd.AddCommand(echoCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".ganzi" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ganzi")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
