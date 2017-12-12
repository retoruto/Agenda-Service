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
	"os"

	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy your account and logout",
	Long:  `NOTICE: THIS COMMAND WILL PERMANTLY REMOVE YOUR ACCOUNT`,
	Run: func(cmd *cobra.Command, args []string) {
		key := getLocalKey()
		if key == "" {
			fmt.Println("Please login first.")
			return
		}
		client := http.Client{}
		req, err := http.NewRequest(http.MethodDelete,
			host+"/v1/users?key="+key, nil)
		panicErr(err)
		res, err := client.Do(req)
		panicErr(err)
		defer res.Body.Close()
		os.Truncate(keyPath, 0)
		fmt.Println("Your account has been removed")
	},
}

func init() {
	RootCmd.AddCommand(destroyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destroyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destroyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
