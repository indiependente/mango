package adapters

// Adapter represents an abstraction able to return a Manga related to a chapter
type Adapter interface {
	Manga(string) (*Manga, error)
}
