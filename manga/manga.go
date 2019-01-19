package manga

import (
	"github.com/pkg/errors"
)

var (
	// ErrOutOfBounds is an out of bounds error.
	ErrOutOfBounds = errors.New("Out of bounds")
	// ErrNilPage is a nil page error.
	ErrNilPage = errors.New("Nil page")
	// ErrPageOverride is a page override error.
	ErrPageOverride = errors.New("Page override")
	// ErrChapterNumberOverride is a chapter number override error.
	ErrChapterNumberOverride = errors.New("Chapter number override")
	// ErrChapterOverride is a chapter override error.
	ErrChapterOverride = errors.New("Chapter override")
	// ErrNilChapter is a nil chapter error.
	ErrNilChapter = errors.New("Nil chapter")
)

// Manga is a collection of chapters.
type Manga struct {
	author   string
	title    string
	chapters map[int]*Chapter
}

// NewManga returns a new pointer to a chapter.
func NewManga() *Manga {
	return &Manga{
		chapters: make(map[int]*Chapter),
	}
}

// NewMangaWithChapters returns a new pointer to a chapter.
func NewMangaWithChapters(cs map[int]*Chapter) *Manga {
	return &Manga{
		chapters: cs,
	}
}

// NewMangaWithChaptersAuthorTitle returns a new pointer to a chapter.
func NewMangaWithChaptersAuthorTitle(a string, t string, cs map[int]*Chapter) *Manga {
	return &Manga{
		author:   a,
		title:    t,
		chapters: cs,
	}
}

// Author returns the author's name.
func (m *Manga) Author() string {
	return m.author
}

// Title returns the title of the manga.
func (m *Manga) Title() string {
	return m.title
}

// Chapter returns a pointer to chapter at index i.
func (m *Manga) Chapter(i int) *Chapter {
	c, ok := m.chapters[i]
	if !ok {
		return nil
	}
	return c
}

// SetChapter sets a chapter at index i.
func (m *Manga) SetChapter(i int, chap *Chapter) error {
	_, ok := m.chapters[i]
	if ok {
		return ErrChapterOverride
	}
	if chap == nil {
		return ErrNilChapter
	}
	m.chapters[i] = chap
	return nil
}
