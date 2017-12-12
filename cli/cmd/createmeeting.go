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
	"net/http"

	"github.com/spf13/cobra"
)

// createmeetingCmd represents the createmeeting command
var createmeetingCmd = &cobra.Command{
	Use:   "createmeeting",
	Short: "Create a meeting",
	Long: `This command will create a meeting hosted by you.
Please type in the information correctly.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() != 4 {
			cmd.Help()
			return
		}
		key := getLocalKey()
		if key == "" {
			fmt.Println("Please login first.")
			return
		}
		title, _ := cmd.Flags().GetString("title")
		members, _ := cmd.Flags().GetStringSlice("members")
		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtime")
		data := struct {
			Title     string   `json:"title"`
			Members   []string `json:"members"`
			Starttime string   `json:"starttime"`
			Endtime   string   `json:"endtime"`
		}{title, members, starttime, endtime}
		buf, err := json.Marshal(data)
		panicErr(err)
		res, err := http.Post(host+"/v1/meetings?key="+key,
			"application/json", bytes.NewBuffer(buf))
		panicErr(err)
		defer res.Body.Close()
		if res.StatusCode == http.StatusCreated {
			fmt.Printf("Meeting '%v' created\n", title)
		} else {
			fmt.Printf("Fail to create meeting '%v'\n", title)
		}
	},
}

func init() {
	RootCmd.AddCommand(createmeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createmeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createmeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createmeetingCmd.Flags().StringP("title", "t", "",
		"Title of the meeting")
	createmeetingCmd.Flags().StringSliceP("members", "m", nil,
		"Members of the meeting")
	createmeetingCmd.Flags().StringP("starttime", "s", "",
		"The start time of the meeting. In form YYYY/MM/DD/HH:MM")
	createmeetingCmd.Flags().StringP("endtime", "e", "",
		"The end time of the meeting. In form YYYY/MM/DD/HH:MM")
}
