package controllers

import "github.com/ppacher/gomopidy/models"

type LibraryController interface {
	Browse(uri string) ([]models.Ref, error)
	Search(models.LibrarySearch) ([]models.SearchResult, error)
}
