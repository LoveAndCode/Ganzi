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
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"path"
	"strings"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Showing current welcome message",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := homedir.Dir()
		if err != nil {
			panic(err)
		}
		path := path.Join(homeDir, ".banner.txt")
		log.Printf("read [%s]\n", path)
		fileByteData, err := ioutil.ReadFile(path)
		data := string(fileByteData)

		if err != nil {
			panic(err)
		}

		log.Println(data)
		data = strings.ReplaceAll(data, "\n", "")
		fontData := figure.NewFigure(data, "", true).String()
		fmt.Println(fontData)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
