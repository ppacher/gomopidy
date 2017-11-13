package models

// TlTrack represents the tltrack model from mopidy
// https://docs.mopidy.com/en/latest/api/models/#mopidy.models.TlTrack
type TlTrack struct {
	ID    int   `json:"tlid"`
	Track Track `json:"track"`
}
