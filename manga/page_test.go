package manga_test

import (
	"image"
	"reflect"
	"testing"

	"github.com/indiependente/mango/manga"
)

func TestNewPage(t *testing.T) {
	type args struct {
		n  int
		im image.Image
	}
	tests := []struct {
		name string
		args args
		want *manga.Page
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manga.NewPage(tt.args.n, tt.args.im); !reflect.DeepEqual(got, tt.want) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := manga.NewPage(
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := manga.NewPage(
				tt.fields.number,
				tt.fields.image,
			)

			if got := p.Image(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Page.Image() = %v, want %v", got, tt.want)
			}
		})
	}
}
