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

// UserLoginCmd represents the UserLogin command
var UserLoginCmd = &cobra.Command{
	Use:   "login -u [UserName] -p [PassWord]",
	Short: "Using UserName with PassWord to login Agenda.",
	Long: `Using UserName and PassWord to login Agenda:

attention:If the PassWord is right,you can login Agenda and use it
If forget the PassWord,you must register another one User`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() != 2 {
			cmd.Help()
			return
		}
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		data := struct {
			Name string `json:"username"`
			Password string `json:"password"`

		} {username, password}
		fmt.Println(data)
		res, err := http.Get(host + "/v1/login?username=" + username + "&password=" + password)
		CheckErr(err)
		defer res.Body.Close()
		if res.StatusCode != 200 {
			fmt.Println("Login failed.")
		} else {
			fmt.Print("Login successfully!")

		}
	  },
}

func init() {
	RootCmd.AddCommand(UserLoginCmd)
	UserLoginCmd.Flags().StringP("username", "u", "", "new user's username")
	UserLoginCmd.Flags().StringP("password", "p", "", "new user's password")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UserLoginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UserLoginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
