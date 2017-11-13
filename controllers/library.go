package controllers

import "github.com/homebot/gomopidy/models"

type LibraryController interface {
	Browse(uri string) ([]models.Ref, error)
	Search(models.LibrarySearch) ([]models.SearchResult, error)
}
