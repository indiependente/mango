package manga_test

import (
	"reflect"
	"testing"

	"github.com/indiependente/mango/manga"
)

func NewChapters() map[int]*manga.Chapter {
	return map[int]*manga.Chapter{
		1: manga.NewChapter(),
	}
}

func TestNewManga(t *testing.T) {
	tests := []struct {
		name string
		want *manga.Manga
	}{
		{
			name: "new manga - empty",
			want: manga.NewManga(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manga.NewManga(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManga() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMangaWithChapters(t *testing.T) {
	type args struct {
		cs map[int]*manga.Chapter
	}
	tests := []struct {
		name string
		args args
		want *manga.Manga
	}{
		{
			name: "new manga - with chapters",
			args: args{cs: NewChapters()},
			want: manga.NewMangaWithChapters(NewChapters()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manga.NewMangaWithChapters(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMangaWithChapters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMangaWithChaptersAuthorTitle(t *testing.T) {
	type args struct {
		a  string
		t  string
		cs map[int]*manga.Chapter
	}
	tests := []struct {
		name string
		args args
		want *manga.Manga
	}{
		{
			name: "new manga - with autor, title and chapters",
			args: args{
				a:  "Hajime Isayama",
				t:  "Shingeki no kyojin",
				cs: NewChapters(),
			},
			want: manga.NewMangaWithChaptersAuthorTitle("Hajime Isayama", "Shingeki no kyojin", NewChapters()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := manga.NewMangaWithChaptersAuthorTitle(tt.args.a, tt.args.t, tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMangaWithChaptersAuthorTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_Author(t *testing.T) {
	type fields struct {
		author   string
		title    string
		chapters map[int]*manga.Chapter
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "author",
			fields: fields{
				author:   "Hajime Isayama",
				title:    "Shingeki no kyojin",
				chapters: NewChapters(),
			},
			want: "Hajime Isayama",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := manga.NewMangaWithChaptersAuthorTitle(tt.fields.author, tt.fields.title, tt.fields.chapters)
			if got := m.Author(); got != tt.want {
				t.Errorf("Manga.Author() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_Title(t *testing.T) {
	type fields struct {
		author   string
		title    string
		chapters map[int]*manga.Chapter
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "author",
			fields: fields{
				author:   "Hajime Isayama",
				title:    "Shingeki no kyojin",
				chapters: NewChapters(),
			},
			want: "Shingeki no kyojin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := manga.NewMangaWithChaptersAuthorTitle(tt.fields.author, tt.fields.title, tt.fields.chapters)
			if got := m.Title(); got != tt.want {
				t.Errorf("Manga.Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_Chapter(t *testing.T) {
	type fields struct {
		author   string
		title    string
		chapters map[int]*manga.Chapter
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *manga.Chapter
	}{
		{
			name: "chapter",
			fields: fields{
				chapters: NewChapters(),
			},
			args: args{
				i: 1,
			},
			want: manga.NewChapter(),
		},
		{
			name: "chapter not found",
			fields: fields{
				chapters: NewChapters(),
			},
			args: args{
				i: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := manga.NewMangaWithChapters(tt.fields.chapters)

			if got := m.Chapter(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Manga.Chapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_SetChapter(t *testing.T) {
	type fields struct {
		author   string
		title    string
		chapters map[int]*manga.Chapter
	}
	type args struct {
		i int
		c *manga.Chapter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "set chapter",
			fields: fields{},
			args: args{
				i: 1,
				c: manga.NewChapter(),
			},
			wantErr: false,
		},
		{
			name: "set chapter - ErrChapterOverride",
			fields: fields{
				chapters: NewChapters(),
			},
			args: args{
				i: 1,
				c: manga.NewChapter(),
			},
			wantErr: true,
		},
		{
			name:   "set chapter - ErrNilChapter",
			fields: fields{},
			args: args{
				i: 1,
				c: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var m *manga.Manga
			if tt.fields.chapters != nil {
				m = manga.NewMangaWithChaptersAuthorTitle(tt.fields.author, tt.fields.title, tt.fields.chapters)
			} else {
				m = manga.NewManga()
			}
			if err := m.SetChapter(tt.args.i, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Manga.SetChapter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
