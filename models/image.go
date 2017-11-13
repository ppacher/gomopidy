package models

// Image represents the image model from mopidy
// https://docs.mopidy.com/en/latest/api/models/#mopidy.models.Image
type Image struct {
	URI    string `json:"uri,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}
