package models

// Album is the album model from Mopidy
// https://docs.mopidy.com/en/latest/api/models/#mopidy.models.Album
type Album struct {
	URI           string   `json:"uri,omitempty"`
	Name          string   `json:"name,omitempty"`
	Artists       []Artist `json:"artists,omitempty"`
	NumTracks     int      `json:"num_tracks,omitempty"`
	NumDiscs      int      `json:"num_discs,omitempty"`
	Date          string   `json:"date,omitempty"`
	MusicBrainzID string   `json:"musicbrainz_id,omitempty"`
	Images        []string `json:"images,omitempty"`
}
