// Copyright © 2017 HinanawiTenshi <dr.paper@live.com>
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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the agenda server",
	Long: `Login to the agenda server, you will receive your own api key.
Keep it secret`,
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
	RootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	loginCmd.Flags().StringP("username", "u",
		"", "Your username.")
	loginCmd.Flags().StringP("password", "p",
		"", "Your password.")
}
