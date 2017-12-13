// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"net/http" 
	"github.com/spf13/cobra"
)

// MeetingDeleteCmd represents the MeetingDelete command
var MeetingDeleteCmd = &cobra.Command{
	Use:   "deleteM -t [title]",
	Short: "delete meeting with the title [title]",
	Long:  `you can delete one meeting with the title [title]`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("Title")
		res, err := http.NewRequest("DELETE", host+"/v1/meeting/"+title, nil)
		panicErr(err)
		client := &http.Client{}
		response, _ := client.Do(res)
		if response.StatusCode != 200 {
			fmt.Println("Delete meeting failed.")
		} else {
			// Decode JSON
			fmt.Println("Delete meeting successfully.")
	}
	},
}

func init() {
	RootCmd.AddCommand(MeetingDeleteCmd)
	MeetingDeleteCmd.Flags().StringP("Title", "t", "", "meeting title")	
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MeetingDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MeetingDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
