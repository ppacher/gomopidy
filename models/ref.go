package models

// RefType defines one of the possible values for `Ref.Type`
type RefType string

// Possible values for `Ref.Type`
const (
	AlbumRef     = RefType("album")
	ArtistRef    = RefType("artist")
	TrackRef     = RefType("track")
	DirectoryRef = RefType("directory")
	PlaylistRef  = RefType("playlist")
)

// Ref is a reference to a Mopidy core model. It contains a type field describing the
// model referenced as well as the a common name field and the URI to inspect/use the
// referenced entry.
type Ref struct {
	URI  string  `json:"uri"`
	Name string  `json:"name"`
	Type RefType `json:"type"`
}

// IsAlbum returns true if `r` is a reference to an album
func (r Ref) IsAlbum() bool {
	return r.Type == AlbumRef
}

// IsTrack returns true if `r` is a reference to an track
func (r Ref) IsTrack() bool {
	return r.Type == TrackRef
}

// IsArtist returns true if `r` is a reference to an artist
func (r Ref) IsArtist() bool {
	return r.Type == ArtistRef
}

// IsDirectory returns true if `r` is a reference to an directory
func (r Ref) IsDirectory() bool {
	return r.Type == DirectoryRef
}

// IsPlaylist returns true if `r` is a reference to an playlist
func (r Ref) IsPlaylist() bool {
	return r.Type == PlaylistRef
}
