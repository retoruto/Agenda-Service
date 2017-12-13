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

// UserDeleteCmd represents the UserDelete command
var UserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "To delete your account in Agenda",
	Long: `you can delete your account in the database of Agenda:

attention:After deleting,you will need to register a new User to login Agenda.`,
Run: func(cmd *cobra.Command, args []string) {

		res, err := http.NewRequest("DELETE", host+"/v1/user", nil)
		panicErr(err)
		
		client := &http.Client{}
		response, _ := client.Do(res)
		
		if response.StatusCode != 200 {
			fmt.Println("Delete user failed.")
		} else {
			// Decode JSON
			fmt.Printf("Delete successfully.")
		}
},
}

func init() {
	RootCmd.AddCommand(UserDeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UserDeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UserDeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
