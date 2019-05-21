package adapters

import "github.com/indiependente/mango/manga"

// Adapter represents an abstraction able to return a Manga related to a chapter
type Adapter interface {
	Manga(string) (*manga.Manga, error)
}
