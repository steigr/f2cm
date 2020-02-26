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
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Download a set of configMaps into the given folders",
	Long:  `Download a set of configMaps into the given folders.`,
	Run: func(cmd *cobra.Command, args []string) {
		checkWorkloadMapping()
		for idx := 0; idx < len(configMapNames); idx++ {
			configmap.Download(namespace, configMapNames[idx], directoryNames[idx])
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}
