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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
		res, err := http.Get(host + "/v1/key?username=" +
			username + "&password=" + password)
		panicErr(err)
		defer res.Body.Close()
		if res.StatusCode != 200 {
			fmt.Println("Login failed. Either Username or password isn't correct")
		} else {
			// Decode JSON.
			body, err := ioutil.ReadAll(res.Body)
			panicErr(err)
			var data map[string]interface{}
			err = json.Unmarshal(body, &data)
			panicErr(err)

			fmt.Printf("Login successfully. Your api key is\n%v\n",
				data["key"].(string))

			// Write the api key to file.
			keyFile, err := os.OpenFile(keyPath,
				os.O_CREATE|os.O_RDWR, os.ModePerm)
			defer keyFile.Close()
			panicErr(err)
			_, err = keyFile.Write([]byte(data["key"].(string)))
			panicErr(err)
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
