// Copyright © 2017 Patrick Pacher <patrick.pacher@gmail.com>
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
	"log"

	"github.com/spf13/cobra"
)

var statusOptions struct {
	Title  bool
	Artist bool
	Album  bool
	Config bool
	All    bool
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current status",
	Run: func(cmd *cobra.Command, args []string) {
		all := statusOptions.All
		checkErr := func(err error) {
			if err != nil {
				log.Fatal(err)
			}
		}

		if !all && !(statusOptions.Title && statusOptions.Artist && statusOptions.Config) {
			all = true
		}

		current, err := cli.Playback().GetCurrentTlTrack()
		checkErr(err)

		if all || statusOptions.Title {
			fmt.Println(current.Track.Name)
		}
		if all || statusOptions.Album {
			fmt.Printf("%s (%s)\n", current.Track.Album.Name, current.Track.Album.URI)
		}
		if all || statusOptions.Artist && len(current.Track.Artists) > 0 {
			fmt.Printf("%s (%s)\n", current.Track.Artists[0].Name, current.Track.Artists[0].URI)
		}

		if all || statusOptions.Config {
			if all || (statusOptions.Title || statusOptions.Album || statusOptions.Artist) {
				fmt.Println("===========================================================")
			}
			playback := cli.Tracklist()
			repeat, err := playback.GetRepeat()
			checkErr(err)
			consume, err := playback.GetConsume()
			checkErr(err)
			single, err := playback.GetSingle()
			checkErr(err)
			random, err := playback.GetRandom()
			checkErr(err)

			tracks, err := playback.GetLength()
			checkErr(err)

			var status []string

			if repeat {
				status = append(status, "Repeat")
			}
			if single {
				status = append(status, "Single")
			}
			if consume {
				status = append(status, "Consume")
			}
			if random {
				status = append(status, "Random")
			}

			fmt.Printf("Tracks: #%d\n", tracks)
		}
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)

	statusCmd.Flags().BoolVarP(&statusOptions.Title, "title", "t", false, "Show current title")
	statusCmd.Flags().BoolVarP(&statusOptions.Artist, "artist", "a", false, "Show current artist")
	statusCmd.Flags().BoolVarP(&statusOptions.Album, "album", "l", false, "Show current album")
	statusCmd.Flags().BoolVarP(&statusOptions.Config, "config", "c", false, "Show current control")
	statusCmd.Flags().BoolVarP(&statusOptions.All, "all", "A", false, "Show everything")
}
