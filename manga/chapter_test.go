package manga_test

import (
	"reflect"
	"testing"

	"github.com/indiependente/mango/manga"
)

func TestNewChapter(t *testing.T) {
	tests := []struct {
		name string
		want *manga.Chapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manga.NewChapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChapterWithNAndPages(t *testing.T) {
	type args struct {
		n     int
		pages []*manga.Page
	}
	tests := []struct {
		name string
		args args
		want *manga.Chapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manga.NewChapterWithNAndPages(tt.args.n, tt.args.pages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChapterWithNAndPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_SetNumber(t *testing.T) {
	type fields struct {
		number int
		pages  []*manga.Page
	}
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := manga.NewChapterWithNAndPages(
				tt.fields.number,
				tt.fields.pages,
			)
			if err := c.SetNumber(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Chapter.SetNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
