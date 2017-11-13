package models

// Artist describes a song artist
// https://docs.mopidy.com/en/latest/api/models/#mopidy.models.Artist
type Artist struct {
	URI           string `json:"uri,omitempty"`
	Name          string `json:"name,omitempty"`
	SortName      string `json:"sortname,omitempty"`
	MusicBrainzID string `json:"musicbrainz_id,omitempty"`
}
