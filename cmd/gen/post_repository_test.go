package main

import (
	"testing"

	"github.com/cou929/please-sleep-2/internal/condition"
	"github.com/cou929/please-sleep-2/internal/post"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type fileInfoMock struct {
	name  string
	isDir bool
}

func (f fileInfoMock) Name() string {
	return f.name
}

func (f fileInfoMock) IsDir() bool {
	return f.isDir
}

type readerMock struct {
	filesInDir    []fileInfo
	contentByName map[string]([]byte)
}

func (r readerMock) ReadDir(dirname string) ([]fileInfo, error) {
	return r.filesInDir, nil
}

func (r readerMock) ReadFile(filename string) ([]byte, error) {
	return r.contentByName[filename], nil
}

func TestPostRepository_isTarget(t *testing.T) {
	type args struct {
		f fileInfo
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "target",
			args: args{
				f: &fileInfoMock{name: "test.md", isDir: false},
			},
			want: true,
		},
		{
			name: "directory is not target",
			args: args{
				f: &fileInfoMock{name: "images", isDir: true},
			},
			want: false,
		},
		{
			name: "only .md file is target",
			args: args{
				f: &fileInfoMock{name: ".DS_Store", isDir: false},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostRepository{}
			if got := r.isTarget(tt.args.f); got != tt.want {
				t.Errorf("PostRepository.isTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostRepository_load(t *testing.T) {
	type fields struct {
		reader reader
		c      *condition.Condition
	}
	tests := []struct {
		name    string
		fields  fields
		want    post.Posts
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				reader: &readerMock{
					filesInDir: []fileInfo{
						&fileInfoMock{
							name:  ".DS_Store",
							isDir: false,
						},
						&fileInfoMock{
							name:  "file001.md",
							isDir: false,
						},
						&fileInfoMock{
							name:  "file002.md",
							isDir: false,
						},
					},
					contentByName: map[string]([]byte){
						"file001.md": ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00"}
file001 content`),
						"file002.md": ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00","tags":["golang"]}
file002 content`),
					},
				},
				c: &condition.Condition{PostPath: ""},
			},
			want: post.Posts{
				&post.Post{
					Filename: "file001.md",
					Raw: ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00"}
file001 content`),
				},
				&post.Post{
					Filename: "file002.md",
					Raw: ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00","tags":["golang"]}
file002 content`),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostRepository{
				reader: tt.fields.reader,
				c:      tt.fields.c,
			}
			got, err := r.load()
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRepository.load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmpopts.IgnoreFields(post.Post{}, "Title", "Issued", "Content", "C")
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("PostRepository.load() diff (-got +want)\n%s", diff)
			}
		})
	}
}

func TestPostRepository_List(t *testing.T) {
	type fields struct {
		posts  post.Posts
		reader reader
		c      *condition.Condition
	}
	tests := []struct {
		name    string
		fields  fields
		want    post.Posts
		wantErr bool
	}{
		{
			name: "cache loaded posts",
			fields: fields{
				posts: post.Posts{
					&post.Post{Filename: "loaded.md"},
				},
				reader: &readerMock{
					filesInDir: []fileInfo{
						&fileInfoMock{
							name:  "file001.md",
							isDir: false,
						},
					},
					contentByName: map[string]([]byte){
						"file001.md": ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00"}
file001 content`)},
				},
				c: &condition.Condition{PostPath: ""},
			},
			want: post.Posts{
				&post.Post{Filename: "loaded.md"},
			},
			wantErr: false,
		},
		{
			name: "load posts if no cache",
			fields: fields{
				posts: (post.Posts)(nil),
				reader: &readerMock{
					filesInDir: []fileInfo{
						&fileInfoMock{
							name:  "file001.md",
							isDir: false,
						},
					},
					contentByName: map[string]([]byte){
						"file001.md": ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00"}
file001 content`)},
				},
				c: &condition.Condition{PostPath: ""},
			},
			want: post.Posts{
				&post.Post{Filename: "file001.md"},
			},
			wantErr: false,
		},
		{
			name: "posts are sorted",
			fields: fields{
				posts: (post.Posts)(nil),
				reader: &readerMock{
					filesInDir: []fileInfo{
						&fileInfoMock{
							name:  "file001.md",
							isDir: false,
						},
						&fileInfoMock{
							name:  "file002.md",
							isDir: false,
						},
					},
					contentByName: map[string]([]byte){
						"file001.md": ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00"}
file001 content`),
						"file002.md": ([]byte)(`{"title":"test post","date":"2015-09-21T12:58:19+09:00"}
file001 content`),
					},
				},
				c: &condition.Condition{PostPath: ""},
			},
			want: post.Posts{
				&post.Post{Filename: "file002.md"},
				&post.Post{Filename: "file001.md"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PostRepository{
				posts:  tt.fields.posts,
				reader: tt.fields.reader,
				c:      tt.fields.c,
			}
			got, err := r.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("PostRepository.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmpopts.IgnoreFields(post.Post{}, "Title", "Issued", "Content", "Raw", "C")
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("PostRepository.List() diff (-got +want)\n%s", diff)
			}
		})
	}
}
