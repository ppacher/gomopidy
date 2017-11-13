package controllers

import (
	"github.com/homebot/gomopidy/client"
	"github.com/homebot/gomopidy/models"
)

type TrackListController interface {
	Add(models.TracklistAddOptions) ([]models.TlTrack, error)
	Remove(models.SearchQuery) ([]models.TlTrack, error)
	Clear() error
	Move(start, end, position int) error
	Shuffel(*models.TracklistShuffelOptions) error
	GetTlTracks() ([]models.TlTrack, error)
	Index(models.TracklistIndex) (int, error)
	GetVersion() (int, error)
	GetLength() (int, error)
	GetTracks() ([]models.Track, error)
	Slice(start, end int) ([]models.TlTrack, error)
	Filter(models.SearchQuery) ([]models.TlTrack, error)

	//
	// Future state
	//

	GetEOTTrackListID() (int, error)
	GetNextTrackListID() (int, error)
	GetPreviousTrackListID() (int, error)
	GetEotTrack(*models.TlTrack) (*models.TlTrack, error)
	GetNextTrack(*models.TlTrack) (*models.TlTrack, error)
	GetPreviousTrack(*models.TlTrack) (*models.TlTrack, error)

	//
	// Options
	//
	GetConsume() (bool, error)
	SetConsume(bool) error
	GetRandom() (bool, error)
	SetRandom(bool) error
	GetRepeat() (bool, error)
	SetRepeat(bool) error
	GetSingle() (bool, error)
	SetSingle(bool) error
}

type tracklist struct {
	cli *client.MopidyClient
}

// NewTracklistController returns a new tracklist controller for the given mopidy client
func NewTracklistController(cli *client.MopidyClient) TrackListController {
	return &tracklist{cli}
}

func (tl *tracklist) Add(what models.TracklistAddOptions) ([]models.TlTrack, error) {
	var res []models.TlTrack

	err := tl.cli.Call("core.tracklist.add", what, &res)
	return res, err
}

func (tl *tracklist) Remove(criteria models.SearchQuery) ([]models.TlTrack, error) {
	var res []models.TlTrack

	err := tl.cli.Call("core.tracklist.remove", criteria, &res)
	return res, err
}

func (tl *tracklist) Clear() error {
	return tl.cli.Call("core.tracklist.clear", nil, nil)
}

func (tl *tracklist) Move(start, end, position int) error {
	return tl.cli.Call("core.tracklist.move", models.TracklistMove{
		Start:    start,
		End:      end,
		Position: position,
	}, nil)
}

func (tl *tracklist) Shuffel(opts *models.TracklistShuffelOptions) error {
	return tl.cli.Call("core.tracklist.shuffel", opts, nil)
}

func (tl *tracklist) GetTlTracks() ([]models.TlTrack, error) {
	var res []models.TlTrack
	err := tl.cli.Call("core.tracklist.get_tl_tracks", nil, &res)
	return res, err
}

func (tl *tracklist) Index(opts models.TracklistIndex) (int, error) {
	var res int

	err := tl.cli.Call("core.tracklist.index", opts, &res)
	return res, err
}

func (tl *tracklist) GetVersion() (int, error) {
	var res int
	err := tl.cli.Call("core.tracklist.get_version", nil, &res)
	return res, err
}

func (tl *tracklist) GetLength() (int, error) {
	var res int
	err := tl.cli.Call("core.tracklist.get_length", nil, &res)
	return res, err
}

func (tl *tracklist) GetTracks() ([]models.Track, error) {
	var res []models.Track
	err := tl.cli.Call("core.tracklist.get_tracks", nil, &res)
	return res, err
}

func (tl *tracklist) Slice(start, end int) ([]models.TlTrack, error) {
	var res []models.TlTrack
	err := tl.cli.Call("core.tracklist.slice", models.TracklistSlice{
		Start: start,
		End:   end,
	}, &res)
	return res, err
}

func (tl *tracklist) Filter(criteria models.SearchQuery) ([]models.TlTrack, error) {
	var res []models.TlTrack
	err := tl.cli.Call("core.tracklist.filter", criteria, &res)
	return res, err
}

func (tl *tracklist) GetEOTTrackListID() (int, error) {
	var res int
	err := tl.cli.Call("core.tracklist.get_eot_tlid", nil, &res)
	return res, err
}

func (tl *tracklist) GetNextTrackListID() (int, error) {
	var res int
	err := tl.cli.Call("core.tracklist.get_next_tlid", nil, &res)
	return res, err
}

func (tl *tracklist) GetPreviousTrackListID() (int, error) {
	var res int
	err := tl.cli.Call("core.tracklist.get_previous_tlid", nil, &res)
	return res, err
}

func (tl *tracklist) GetEotTrack(ref *models.TlTrack) (*models.TlTrack, error) {
	var res models.TlTrack
	err := tl.cli.Call("core.tracklist.eot_track", ref, &res)
	return &res, err
}

func (tl *tracklist) GetNextTrack(ref *models.TlTrack) (*models.TlTrack, error) {
	var res models.TlTrack
	err := tl.cli.Call("core.tracklist.next_track", ref, &res)
	return &res, err
}

func (tl *tracklist) GetPreviousTrack(ref *models.TlTrack) (*models.TlTrack, error) {
	var res models.TlTrack
	err := tl.cli.Call("core.tracklist.previous_track", ref, &res)
	return &res, err
}

func (tl *tracklist) GetConsume() (bool, error) {
	var res bool
	err := tl.cli.Call("core.tracklist.get_consume", nil, &res)
	return res, err
}

func (tl *tracklist) SetConsume(b bool) error {
	return tl.cli.Call("core.tracklist.set_consume", b, nil)
}

func (tl *tracklist) GetRandom() (bool, error) {
	var res bool
	err := tl.cli.Call("core.tracklist.get_random", nil, &res)
	return res, err
}

func (tl *tracklist) SetRandom(b bool) error {
	return tl.cli.Call("core.tracklist.set_random", b, nil)
}

func (tl *tracklist) GetRepeat() (bool, error) {
	var res bool
	err := tl.cli.Call("core.tracklist.get_repeat", nil, &res)
	return res, err
}

func (tl *tracklist) SetRepeat(b bool) error {
	return tl.cli.Call("core.tracklist.set_repeat", b, nil)
}

func (tl *tracklist) GetSingle() (bool, error) {
	var res bool
	err := tl.cli.Call("core.tracklist.get_single", nil, &res)
	return res, err
}

func (tl *tracklist) SetSingle(b bool) error {
	return tl.cli.Call("core.tracklist.set_single", b, nil)
}
