package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
	"strings"
)

var setCommand = &cobra.Command{
	Use:     "set",
	Example: "set \"Welcome! Ganzi Production Server!\"",
	Short:   "set shell display text",
	Run: func(cmd *cobra.Command, args []string) {
		// concat all text in passing as arguments
		text := strings.Join(args, "")

		// get user home directory path
		dir, err := homedir.Dir()

		if err != nil {
			panic(err)
		}

		// create text banner file on user home directory
		bannerFileName := createBannerTextFile(text, dir)

		// get current user shell program type
		shell := getCurrentShellType()

		// Search User Shell Condition
		if strings.Contains(shell, "bash") && exists(dir, ".bashrc") {
			setBannerFileDisplay(dir, ".bashrc", bannerFileName)
		} else if strings.Contains(shell, "zsh") && exists(dir, ".zshrc") {
			setBannerFileDisplay(dir, ".zshrc", bannerFileName)
		} else if exists(dir, ".profile") {
			setBannerFileDisplay(dir, ".profile", bannerFileName)
		}
	},
}

func getCurrentShellType() string {
	shell := os.Getenv("SHELL")
	fmt.Println(shell)
	return shell
}

func setBannerFileDisplay(dir string, targetFile string, bannerFileName string) {
	file, err := os.OpenFile(path.Join(dir, targetFile), os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Println("open " + path.Join(dir, targetFile) + " and write user text")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString("cat " + path.Join(dir, bannerFileName))

	if err != nil {
		panic(err)
	}
}

func exists(dir string, filename string) bool {
	if _, err := os.Stat(path.Join(dir, filename)); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createBannerTextFile(text string, homeDir string) string {
	textFileName := "banner.txt"
	filePath := path.Join(homeDir, textFileName)
	log.Print("Create file: " + filePath)

	fo, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	defer fo.Close()

	if err != nil {
		panic(err)
	}

	_, err = fo.WriteString(text + "\n")

	if err != nil {
		panic(err)
	}

	return textFileName
}

func init() {
	rootCmd.AddCommand(setCommand)
}
