package models

// SearchResult represents the search-result model from mopidy
// https://docs.mopidy.com/en/latest/api/models/#mopidy.models.SearchResult
type SearchResult struct {
	URI     string   `json:"uri,omitempty"`
	Tracks  []Track  `json:"tracks,omitempty"`
	Artists []Artist `json:"artist,omitempty"`
	Albums  []Album  `json:"albums,omitempty"`
}
