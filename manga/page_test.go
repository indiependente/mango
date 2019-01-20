package manga

import (
	"image"
	"reflect"
	"testing"
)

func NewRectangle() image.Image {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	return img
}

func TestNewPage(t *testing.T) {
	type args struct {
		n  int
		im image.Image
	}
	tests := []struct {
		name string
		args args
		want *Page
	}{
		{
			name: "new page",
			args: args{
				n:  1,
				im: nil,
			},
			want: &Page{
				number: 1,
				image:  nil,
			},
		},
		{
			name: "page - with image",
			args: args{
				n:  1,
				im: NewRectangle(),
			},
			want: &Page{
				number: 1,
				image:  NewRectangle(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPage(tt.args.n, tt.args.im); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_Number(t *testing.T) {
	type fields struct {
		number int
		image  image.Image
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "page - number",
			fields: fields{
				number: 1,
				image:  nil,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPage(
				tt.fields.number,
				tt.fields.image,
			)
			if got := p.Number(); got != tt.want {
				t.Errorf("Page.Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_Image(t *testing.T) {
	type fields struct {
		number int
		image  image.Image
	}
	tests := []struct {
		name   string
		fields fields
		want   image.Image
	}{
		{
			name: "page - image",
			fields: fields{
				number: 1,
				image:  NewRectangle(),
			},
			want: NewRectangle(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPage(
				tt.fields.number,
				tt.fields.image,
			)

			if got := p.Image(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Page.Image() = %v, want %v", got, tt.want)
			}
		})
	}
}
