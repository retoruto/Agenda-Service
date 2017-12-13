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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
  "io/ioutil"
  "Agenda-Service/entity"
	"github.com/spf13/cobra"
)

// MeetingCreateCmd represents the MeetingCreate command
var MeetingCreateCmd = &cobra.Command{
	Use:   "create -t [Title] -p [Participator] -s [StartTime] -e [EndTime]",
	Short: "To create a new meeting",
	Long: `To create a new meeting with:

[Title] the Title of the meeting
[Participator] the Participator of the meeting,the Participator can only attend one meeting during one meeting time
[StartTime] the StartTime of the meeting
[EndTime] the EndTime of the meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		
		if cmd.Flags().NFlag() != 4 {
			cmd.Help()
			return
		}
		title, _ := cmd.Flags().GetString("Title")
		members, _ := cmd.Flags().GetStringSlice("Participator")
		starttime, _ := cmd.Flags().GetString("StartTime")
		endtime, _ := cmd.Flags().GetString("EndTime")
		/*
		data := struct {
			Name string `json:"username"`
			Password string `json:"password"`
			Email    string `json:"email"`
			Phone    string `json:"phone"`
		}{username, password, email, phone}
		*/
		
		buf, err := json.Marshal(entity.Meeting{"", title, starttime, endtime, members})
		fmt.Println(entity.CurrentUser.Name)
		fmt.Println(title)
		fmt.Println(starttime)
		fmt.Println(endtime)
		fmt.Println(members)

		
		panicErr(err)
		res, err := http.Post(host+"/v1/meeting","application/json", bytes.NewBuffer(buf))
		panicErr(err)
		defer res.Body.Close()
		fmt.Println(res.StatusCode)
		if res.StatusCode != http.StatusCreated {
			fmt.Println("CreateMeeting failed. Meeting already exists")
		} else {
			// Decode JSON
			body, err := ioutil.ReadAll(res.Body)
			panicErr(err)
			var data entity.Meeting
			err = json.Unmarshal(body, &data)
			panicErr(err)
			
			fmt.Printf("CreateMeeting successfully. \n%v\n", string(body))
	}
	},
}

func init() {
	RootCmd.AddCommand(MeetingCreateCmd)
	MeetingCreateCmd.Flags().StringP("Title", "t", "", "meeting title")
	MeetingCreateCmd.Flags().StringSliceP("Participator", "p", []string{}, "meeting's participator")
	MeetingCreateCmd.Flags().StringP("StartTime", "s", "", "meeting's startTime")
	MeetingCreateCmd.Flags().StringP("EndTime", "e", "", "meeting's endTime")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// MeetingCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// MeetingCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
