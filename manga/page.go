package manga

import "image"

// Page is a pair number, Image.
type Page struct {
	number int
	image  image.Image
}

// NewPage returns a new pointer to Page.
func NewPage(n int, im image.Image) *Page {
	return &Page{
		number: n,
		image:  im,
	}
}

// Number returns the page number.
func (p *Page) Number() int {
	return p.number
}

// Image returns the page image.
func (p *Page) Image() image.Image {
	return p.image
}
