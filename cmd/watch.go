/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"github.com/steigr/f2cm/pkg/configmap"
	"github.com/steigr/f2cm/pkg/watch"
	"log"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch a file or directory and sync it into a configMap",
	Long:  `Watch checks the contents of a given directory or file and updates a configmap.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkWorkloadMapping()

		defer watch.Close()

		for idx := 0; idx < len(configMapNames); idx++ {
			configmap.Upload(namespace,directoryNames[idx], configMapNames[idx])
			events, errors, err := watch.Init(directoryNames[idx])
			if err != nil {
				log.Panicln(err)
			}
			configmap.WatchAndUpload(namespace, directoryNames[idx], configMapNames[idx], events, errors)
		}

		waitForTermination()
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
