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

	"github.com/spf13/cobra"
)

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Show all users",
	Long:  `This command returns all users that have registerred`,
	Run: func(cmd *cobra.Command, args []string) {
		key := getLocalKey()
		if key == "" {
			fmt.Println("Please login first.")
			return
		}
		res, err := http.Get(host + "/v1/users?key=" + key)
		panicErr(err)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		panicErr(err)
		var data []map[string]interface{}
		err = json.Unmarshal(body, &data)
		panicErr(err)
		fmt.Println(string(body))
	},
}

func init() {
	RootCmd.AddCommand(usersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
