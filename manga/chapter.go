package manga

import "github.com/pkg/errors"

// Chapter is a collection of related pages.
type Chapter struct {
	number int
	pages  []*Page
}

// NewChapter returns a new pointer to a chapter.
func NewChapter() *Chapter {
	return &Chapter{
		number: 0,
		pages:  make([]*Page, 0, 0),
	}
}

// NewChapterWithNAndPages returns a new pointer to a chapter.
func NewChapterWithNAndPages(n int, pages []*Page) *Chapter {
	return &Chapter{
		number: n,
		pages:  pages,
	}
}

// Number returns the chapter number.
func (c *Chapter) Number() int {
	return c.number
}

// SetNumber sets the chapter's number.
func (c *Chapter) SetNumber(n int) error {
	if c.number != 0 {
		return ErrChapterNumberOverride
	}
	c.number = n
	return nil
}

// Len returns the length of the chapter, given by the number of pages it contains.
func (c *Chapter) Len() int {
	return len(c.pages)
}

// Pages returns the chapter pages.
func (c *Chapter) Pages() []*Page {
	return c.pages
}

// SetPage sets page p in pages at position index.
func (c *Chapter) SetPage(index int, p *Page) error {
	if index < 0 || index >= len(c.pages) {
		return errors.Wrap(ErrOutOfBounds, "Could not set page")
	}

	if p == nil {
		return errors.Wrap(ErrNilPage, "Could not set page")
	}

	if c.pages[index] != nil {
		return errors.Wrap(ErrPageOverride, "Could not set page")
	}

	c.pages[index] = p
	return nil
}
