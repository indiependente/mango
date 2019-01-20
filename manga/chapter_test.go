package manga

import (
	"reflect"
	"testing"
)

func TestNewChapter(t *testing.T) {
	tests := []struct {
		name string
		want *Chapter
	}{
		{
			name: "new chapter",
			want: &Chapter{
				number: 0,
				pages:  make([]*Page, 0, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewChapterWithNAndPages(t *testing.T) {
	type args struct {
		n     int
		pages []*Page
	}
	tests := []struct {
		name string
		args args
		want *Chapter
	}{
		{
			name: "new chapter with n and pages",
			args: args{
				n:     1,
				pages: make([]*Page, 0, 0),
			},
			want: &Chapter{
				number: 1,
				pages:  make([]*Page, 0, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChapterWithNAndPages(tt.args.n, tt.args.pages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChapterWithNAndPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_SetNumber(t *testing.T) {
	type fields struct {
		number int
		pages  []*Page
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
		{
			name: "chapter - set number",
			fields: fields{
				number: 0,
				pages:  make([]*Page, 0, 0),
			},
			args: args{
				n: 1,
			},
			wantErr: false,
		},
		{
			name: "chapter - set number - number override",
			fields: fields{
				number: 2,
				pages:  make([]*Page, 0, 0),
			},
			args: args{
				n: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				number: tt.fields.number,
				pages:  tt.fields.pages,
			}

			if err := c.SetNumber(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Chapter.SetNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChapter_Number(t *testing.T) {
	type fields struct {
		number int
		pages  []*Page
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "chapter - number",
			fields: fields{
				number: 1,
				pages:  make([]*Page, 0, 0),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				number: tt.fields.number,
				pages:  tt.fields.pages,
			}
			if got := c.Number(); got != tt.want {
				t.Errorf("Chapter.Number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_Len(t *testing.T) {
	type fields struct {
		number int
		pages  []*Page
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "chapter - len",
			fields: fields{
				number: 1,
				pages:  make([]*Page, 0, 0),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				number: tt.fields.number,
				pages:  tt.fields.pages,
			}
			if got := c.Len(); got != tt.want {
				t.Errorf("Chapter.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_Pages(t *testing.T) {
	type fields struct {
		number int
		pages  []*Page
	}
	tests := []struct {
		name   string
		fields fields
		want   []*Page
	}{
		{
			name: "chapter - pages",
			fields: fields{
				number: 1,
				pages:  make([]*Page, 0, 0),
			},
			want: make([]*Page, 0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				number: tt.fields.number,
				pages:  tt.fields.pages,
			}
			if got := c.Pages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chapter.Pages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_SetPage(t *testing.T) {
	type fields struct {
		number int
		pages  []*Page
	}
	type args struct {
		index int
		p     *Page
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "chapter - set page",
			fields: fields{
				number: 1,
				pages:  make([]*Page, 1, 1),
			},
			args: args{
				index: 0,
				p:     &Page{},
			},
			wantErr: false,
		},
		{
			name: "chapter - set page - out of bounds",
			fields: fields{
				number: 1,
				pages:  make([]*Page, 1, 1),
			},
			args: args{
				index: 10,
				p:     &Page{},
			},
			wantErr: true,
		},
		{
			name: "chapter - set page - nil page",
			fields: fields{
				number: 1,
				pages:  make([]*Page, 1, 1),
			},
			args: args{
				index: 0,
				p:     nil,
			},
			wantErr: true,
		},
		{
			name: "chapter - set page - out of bounds",
			fields: fields{
				number: 1,
				pages: []*Page{
					{
						number: 1,
						image:  nil,
					},
				},
			},
			args: args{
				index: 0,
				p:     &Page{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Chapter{
				number: tt.fields.number,
				pages:  tt.fields.pages,
			}
			if err := c.SetPage(tt.args.index, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Chapter.SetPage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
