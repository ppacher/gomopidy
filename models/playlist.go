package models

// Playlist represents the playlist model from mopidy
// https://docs.mopidy.com/en/latest/api/models/#mopidy.models.Playlist
type Playlist struct {
	URI          string  `json:"uri,omitempty"`
	Name         string  `json:"name,omitempty"`
	Tracks       []Track `json:"tracks,omitempty"`
	LastModified int     `json:"last_modified,omitempty"`
}
