package manga

import (
	"reflect"
	"testing"
)

func NewChapters() map[int]*Chapter {
	return map[int]*Chapter{
		1: &Chapter{
			number: 0,
			pages:  nil,
		},
	}
}

func TestNewManga(t *testing.T) {
	tests := []struct {
		name string
		want *Manga
	}{
		{
			name: "new manga - empty",
			want: &Manga{
				chapters: make(map[int]*Chapter),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewManga(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManga() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMangaWithChapters(t *testing.T) {
	type args struct {
		cs map[int]*Chapter
	}
	tests := []struct {
		name string
		args args
		want *Manga
	}{
		{
			name: "new manga - with chapters",
			args: args{cs: NewChapters()},
			want: &Manga{
				chapters: NewChapters(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMangaWithChapters(tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMangaWithChapters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMangaWithChaptersAuthorTitle(t *testing.T) {
	type args struct {
		a  string
		t  string
		cs map[int]*Chapter
	}
	tests := []struct {
		name string
		args args
		want *Manga
	}{
		{
			name: "new manga - with autor, title and chapters",
			args: args{
				a:  "Hajime Isayama",
				t:  "Shingeki no kyojin",
				cs: NewChapters(),
			},
			want: &Manga{
				author:   "Hajime Isayama",
				title:    "Shingeki no kyojin",
				chapters: NewChapters(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMangaWithChaptersAuthorTitle(tt.args.a, tt.args.t, tt.args.cs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMangaWithChaptersAuthorTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_Author(t *testing.T) {
	type fields struct {
		author   string
		title    string
		chapters map[int]*Chapter
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
			m := &Manga{
				author:   tt.fields.author,
				title:    tt.fields.title,
				chapters: tt.fields.chapters,
			}
			if got := m.Author(); got != tt.want {
				t.Errorf("Author() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_Title(t *testing.T) {
	type fields struct {
		author   string
		title    string
		chapters map[int]*Chapter
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
			m := &Manga{
				author:   tt.fields.author,
				title:    tt.fields.title,
				chapters: tt.fields.chapters,
			}
			if got := m.Title(); got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_Chapter(t *testing.T) {
	type fields struct {
		chapters map[int]*Chapter
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Chapter
	}{
		{
			name: "chapter",
			fields: fields{
				chapters: NewChapters(),
			},
			args: args{
				i: 1,
			},
			want: &Chapter{},
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
			m := &Manga{
				chapters: tt.fields.chapters,
			}

			if got := m.Chapter(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManga_SetChapter(t *testing.T) {
	type fields struct {
		author   string
		title    string
		chapters map[int]*Chapter
	}
	type args struct {
		i int
		c *Chapter
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
				c: &Chapter{},
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
				c: &Chapter{},
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
			var m *Manga
			if tt.fields.chapters != nil {
				m = &Manga{
					author:   tt.fields.author,
					title:    tt.fields.title,
					chapters: tt.fields.chapters,
				}
			} else {
				m = &Manga{
					chapters: make(map[int]*Chapter),
				}
			}
			if err := m.SetChapter(tt.args.i, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SetChapter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
