package controllers

import "github.com/homebot/gomopidy/client"

type CoreController interface {
	// Playback returns a new PlaybackController
	Playback() PlaybackController

	// Tracklist returns a new TracklistController
	Tracklist() TrackListController

	// GetURISchemes returns a list of supported URI schemes
	GetURISchemes() ([]string, error)

	// GetVersion returns the version of the Mopidy server
	GetVersion() (string, error)
}

type core struct {
	cli *client.MopidyClient
}

// New returns a new mopidy controller
func New(cli *client.MopidyClient) CoreController {
	return &core{cli}
}

func (c *core) Playback() PlaybackController {
	return NewPlaybackController(c.cli)
}

func (c *core) Tracklist() TrackListController {
	return NewTracklistController(c.cli)
}

func (c *core) GetURISchemes() ([]string, error) {
	var res []string

	if err := c.cli.Call("core.get_uri_schemes", nil, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *core) GetVersion() (string, error) {
	var res string

	err := c.cli.Call("core.get_version", nil, &res)
	return res, err
}
