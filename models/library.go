package models

type SearchQuery map[string]interface{}

type LibrarySearch struct {
	Query SearchQuery `json:"query"`
	URIs  []string    `json:"uris"`
	Exact bool        `json:"exact"`
}
