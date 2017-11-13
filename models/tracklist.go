package models

type TracklistAddOptions struct {
	tracks     []Track  `json:"tracks,omitempty"`
	URI        string   `json:"uri,omitempty"`
	URIs       []string `json:"uris,omitempty"`
	AtPosition *int     `json:"at_position,omitempty"`
}

type TracklistShuffelOptions struct {
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`
}

type TracklistIndex struct {
	Track *TlTrack `json:"tl_track,omitempty"`
	ID    *int     `json:"tlid,omitempty"`
}

type TracklistMove struct {
	Start    int `json:"start"`
	End      int `json:"end"`
	Position int `json:"at_position"`
}

type TracklistSlice struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
