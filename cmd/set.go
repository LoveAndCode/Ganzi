/*
Copyright © 2021 NAME HERE <practice1356@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/mitchellh/go-homedir"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Typing text message which displayed shell activate message",
	Example: `ganzi set "Ganzi Production Server!"`,
	Run: func(cmd *cobra.Command, args []string) {
		// concat all text in passing as arguments
		text := strings.Join(args, "")

		// get user home directory path
		dir, err := homedir.Dir()

		if err != nil {
			panic(err)
		}

		// create text banner file on user home directory
		createBannerTextFile(text, dir)

		// get current user shell program type
		shell := getCurrentShellType()

		// Search User Shell Condition
		if strings.Contains(shell, "bash") && exists(dir, ".bashrc") {
			shellConfiguration(dir, ".bashrc")
		} else if strings.Contains(shell, "zsh") && exists(dir, ".zshrc") {
			shellConfiguration(dir, ".zshrc")
		} else if exists(dir, ".profile") {
			shellConfiguration(dir, ".profile")
		}
	},
}

func getCurrentShellType() string {
	shell := os.Getenv("SHELL")
	log.Printf("-- current shell %s\n", shell)
	return shell
}

func shellConfiguration(dir string, targetFile string) {
	filePath := path.Join(dir, targetFile)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	log.Printf("-- %s open and confguration begin\n", filePath)

	// shell command for read banner file
	configuration := "\n# Ganzi Configuration \n" +
		"if [[ -r \"$HOME/.banner.txt\" ]]; then\n" +
		"	cat \"$HOME/.banner.txt\";\n" +
		"fi"

	defer file.Close()

	if err != nil {
		panic(err)
	}

	// write shell command in current shell configuration file
	_, err = file.WriteString(configuration)

	if err != nil {
		panic(err)
	}
}

func co() {

}

func exists(dir string, filename string) bool {
	if _, err := os.Stat(path.Join(dir, filename)); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createBannerTextFile(text string, homeDir string) {
	textFileName := ".banner.txt"
	filePath := path.Join(homeDir, textFileName)
	log.Printf("-- create banner file in %s", filePath)

	fo, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	defer fo.Close()

	if err != nil {
		panic(err)
	}

	_, err = fo.WriteString("\n" + text + "\n")

	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(setCmd)
}
