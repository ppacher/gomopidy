package controllers

import (
	"fmt"
	"time"

	"github.com/homebot/gomopidy/client"
	"github.com/homebot/gomopidy/models"
)

// PlayOptions for PlaybackController.Play()
type PlayOptions struct {
	TlTrack *models.TlTrack `json:"tl_track,omitempty"`
	TlID    string          `json:"tlid,omitempty"`
}

// PlaybackState is the current state of playback in Mopidy
type PlaybackState string

const (
	// PlaybackStopped indicates that playback is currently stopped
	PlaybackStopped = PlaybackState("stopped")

	// PlaybackPlaying indicates that mopidy is currently playing music
	PlaybackPlaying = PlaybackState("playing")

	// PlaybackPaused indicates that playback is currently paused
	PlaybackPaused = PlaybackState("paused")
)

// PlaybackController interfaces with Mopidy's Playback controller (core.playback)
type PlaybackController interface {
	// Play the given track of, if PlayOptions is nil, play the currently active track
	Play(*PlayOptions) error

	// Next changes to the next track in the tracklist
	Next() error

	// Previous changes to the previous track in the tracklist
	Previous() error

	// Stop playback
	Stop() error

	// Pause playback
	Pause() error

	// Resume playback if paused
	Resume() error

	// Seek to the given position
	Seek(time.Duration) error

	// GetCurrentTlTrack returns the current track in the tracklist
	GetCurrentTlTrack() (*models.TlTrack, error)

	// GetCurrentTrack returns the current track in the tracklist (but only Track instead of TlTrack)
	GetCurrentTrack() (*models.Track, error)

	// GetStreamTitle returns the title of the current stream
	GetStreamTitle() (string, error)

	// GetTimePosition returns the position of the current track
	GetTimePosition() (time.Duration, error)

	// GetState returns the current playback state
	GetState() (PlaybackState, error)

	// SetState sets the plyback state
	SetState(state PlaybackState) error
}

type playback struct {
	cli *client.MopidyClient
}

func NewPlaybackController(cli *client.MopidyClient) PlaybackController {
	return &playback{cli}
}

func (p *playback) Play(opt *PlayOptions) error {
	return p.cli.Call("core.playback.play", opt, nil)
}

func (p *playback) Next() error {
	return p.cli.Call("core.playback.next", nil, nil)
}

func (p *playback) Stop() error {
	return p.cli.Call("core.playback.stop", nil, nil)
}

func (p *playback) Previous() error {
	return p.cli.Call("core.playback.previous", nil, nil)
}

func (p *playback) Pause() error {
	return p.cli.Call("core.playback.pause", nil, nil)
}

func (p *playback) Resume() error {
	return p.cli.Call("core.playback.resume", nil, nil)
}

func (p *playback) Seek(duration time.Duration) error {
	return p.cli.Call("core.playback.seek", map[string]int{
		"time_position": int(duration.Nanoseconds() / 1000),
	}, nil)
}

func (p *playback) GetCurrentTlTrack() (*models.TlTrack, error) {
	var tl models.TlTrack

	if err := p.cli.Call("core.playback.get_current_tl_track", nil, &tl); err != nil {
		return nil, err
	}

	return &tl, nil
}

func (p *playback) GetCurrentTrack() (*models.Track, error) {
	var tl models.Track

	if err := p.cli.Call("core.playback.get_current_track", nil, &tl); err != nil {
		return nil, err
	}

	return &tl, nil
}

func (p *playback) GetStreamTitle() (string, error) {
	var s string

	if err := p.cli.Call("core.playback.get_stream_title", nil, &s); err != nil {
		return "", err
	}

	return s, nil
}

func (p *playback) GetTimePosition() (time.Duration, error) {
	var ms int

	if err := p.cli.Call("core.playback.get_time_position", nil, &ms); err != nil {
		return time.Duration(0), err
	}

	return time.Duration(ms * 1000), nil
}

func (p *playback) GetState() (PlaybackState, error) {
	var state string

	if err := p.cli.Call("core.playback.get_state", nil, &state); err != nil {
		return "", err
	}

	switch state {
	case string(PlaybackPaused):
		return PlaybackPaused, nil
	case string(PlaybackPlaying):
		return PlaybackPlaying, nil
	case string(PlaybackStopped):
		return PlaybackStopped, nil
	default:
		return "", fmt.Errorf("invalid playback state: %s", state)
	}
}

func (p *playback) SetState(state PlaybackState) error {
	return p.cli.Call("core.playback.set_state", map[string]string{
		"new_state": string(state),
	}, nil)
}
