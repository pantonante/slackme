// Copyright © 2019 NAME HERE p.antonante@gmail.com
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
	"errors"
	"fmt"
	"strings"

	"github.com/pantonante/slackme/slackhooks"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Message struct {
	Text string `json:"text"`
}

var channel string

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Send a message to your Slack channel",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		url := viper.GetString(channel)
		if url == "" {
			return errors.New("Channel name is invalid")
		}
		message := &Message{Text: strings.Join(args, " ")}
		payload, err := json.Marshal(message)
		if err != nil {
			return errors.New("Error while converting message to JSON")
		}
		retcode := slackhooks.SendMessage(url, payload)
		fmt.Println(retcode)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
	messageCmd.Flags().StringVarP(&channel, "channel", "c", "default", "Channel name where to send the message (match your .slackme.yaml file)")
}
