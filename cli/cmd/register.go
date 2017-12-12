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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new account",
	Long:  `Register a new account, and an api key will be generated`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() != 4 {
			cmd.Help()
			return
		}
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		data := struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
		}{username, password, email, phone}
		buf, err := json.Marshal(data)
		panicErr(err)
		res, err := http.Post(host+"/v1/key",
			"application/json", bytes.NewBuffer(buf))
		panicErr(err)
		defer res.Body.Close()
		if res.StatusCode != 201 {
			fmt.Println("Register failed. Username already exists")
		} else {
			// Decode JSON
			body, err := ioutil.ReadAll(res.Body)
			panicErr(err)
			var data map[string]interface{}
			err = json.Unmarshal(body, &data)
			panicErr(err)

			fmt.Printf("Register successfully. You are logged in as\n%v\n", string(body))

			// Write the api key to file
			keyFile, err := os.OpenFile(keyPath,
				os.O_CREATE|os.O_RDWR, os.ModePerm)
			panicErr(err)
			defer keyFile.Close()
			_, err = keyFile.Write([]byte(data["key"].(string)))
			panicErr(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	registerCmd.Flags().StringP("username", "u",
		"", "Your username")
	registerCmd.Flags().StringP("password", "p",
		"", "Your password")
	registerCmd.Flags().StringP("email", "e",
		"", "Your email address")
	registerCmd.Flags().StringP("phone", "o",
		"", "Your phone number")
}
