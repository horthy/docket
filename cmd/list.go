// Copyright © Copyright 2016 Dexter Horthy
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of all Allocations",
	Long:  "TODO",
	RunE: func(cmd *cobra.Command, args []string) error {
		return NewCli(cmd, args).List()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().String("host", "http://localhost:3000", "The host to use")
}
