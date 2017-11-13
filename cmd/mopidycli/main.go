package main

import (
	"fmt"

	"github.com/homebot/gomopidy/client"
	"github.com/homebot/gomopidy/controllers"
)

func main() {
	endpoint := "http://music:6680/mopidy/rpc"

	cli := client.NewMopidyClient(endpoint)

	if cli == nil {
		panic("failed to create mopidy client")
	}

	core := controllers.New(cli)
	playback := core.Playback()

	version, _ := core.GetVersion()
	fmt.Printf("Version: %s\n", version)

	tracklist := core.Tracklist()

	consume, _ := tracklist.GetConsume()
	single, _ := tracklist.GetSingle()
	random, _ := tracklist.GetRandom()
	repeat, _ := tracklist.GetRepeat()

	fmt.Printf("Consume: %v\n", consume)
	fmt.Printf("Single: %v\n", single)
	fmt.Printf("Random: %v\n", random)
	fmt.Printf("Repeat: %v\n", repeat)

	tracklist.SetConsume(!consume)
	tracklist.SetSingle(!single)
	tracklist.SetRepeat(!repeat)
	tracklist.SetRandom(!random)

	length, _ := tracklist.GetLength()
	tracks, _ := tracklist.GetTlTracks()

	if length != len(tracks) {
		panic("invalid tracklist length")
	}
	fmt.Printf("Number of tracks: %d\n", length)

	state, err := playback.GetState()
	if err != nil {
		panic(err)
	}
	fmt.Printf("State: %s\n", state)

	switch state {
	case controllers.PlaybackPaused:
		playback.Resume()
	case controllers.PlaybackPlaying:
		playback.Pause()
	}

	currentTrack, err := playback.GetCurrentTlTrack()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", currentTrack)
}
