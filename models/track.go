package models

// Track represents the track model from mopidy
// https://docs.mopidy.com/en/latest/api/models/#mopidy.models.Track
type Track struct {
	URI           string   `json:"uri,omitempty"`
	Name          string   `json:"name,omitempty"`
	Artists       []Artist `json:"artists,omitempty"`
	Album         Album    `json:"album"`
	Composers     []Artist `json:"composers,omitempty"`
	Performers    []Artist `json:"performers,omitempty"`
	Genre         string   `json:"genre,omitempty"`
	TrackNo       int      `json:"track_no,omitempty"`
	DiscNo        int      `json:"disc_no,omitempty"`
	Date          string   `json:"data,omitempty"`
	Length        int      `json:"length,omitempty"`
	Bitrate       int      `json:"bitrate,omitempty"`
	Comment       string   `json:"comment,omitempty"`
	MusicBrainzID string   `json:"musicbrainz_id,omitempty"`
	LastModified  int      `json:"last_modified,omitempty"`
}
