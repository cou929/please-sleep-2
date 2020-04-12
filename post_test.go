package main

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_parseState_transit(t *testing.T) {
	type fields struct {
		cur parseStatus
	}
	type args struct {
		next parseStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    parseStatus
		wantErr bool
	}{
		{
			name:    "normal transition",
			fields:  fields{cur: headerParsed},
			args:    args{next: bodyParsing},
			want:    bodyParsing,
			wantErr: false,
		},
		{
			name:    "invalid transition",
			fields:  fields{cur: bodyParsing},
			args:    args{next: initialized},
			want:    bodyParsing,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &parseState{
				cur: tt.fields.cur,
			}
			if err := s.transit(tt.args.next); (err != nil) != tt.wantErr {
				t.Errorf("parseState.transit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if s.currentStatus() != tt.want {
				t.Errorf("parseState.transit(%v) = %v, want %v", tt.args.next, s.currentStatus(), tt.want)
			}
		})
	}
}

func TestPost_parseHeader(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    *postHeader
		wantErr bool
	}{
		{
			name: "valid json string",
			args: args{line: `{"title":"test post","date":"2014-09-21T12:58:19+09:00","tags":["golang"]}`},
			want: &postHeader{
				Title:  "test post",
				Issued: time.Date(2014, 9, 21, 12, 58, 19, 0, time.FixedZone("JST", 9*60*60)),
			},
			wantErr: false,
		},
		{
			name:    "invalid json string",
			args:    args{line: `{"title":"test post","date":"2014-09-21T12:58:19+09:00","tags":["golang"]`}, // no trailing bracket
			want:    (*postHeader)(nil),
			wantErr: true,
		},
		{
			name:    "required date field",
			args:    args{line: `{"title":"test post","tags":["golang"]}`},
			want:    (*postHeader)(nil),
			wantErr: true,
		},
		{
			name:    "required title field",
			args:    args{line: `{"date":"2014-09-21T12:58:19+09:00","tags":["golang"]}`},
			want:    (*postHeader)(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{}
			got, err := p.parseHeader(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post.parseHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Post.parseHeader() diff (-got +want)\n%s", diff)
			}
		})
	}
}

func TestPost_parseRaw(t *testing.T) {
	type args struct {
		raw []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *postContent
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				raw: ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00","tags":["golang"]}


body here
blah blah`),
			},
			want: &postContent{
				Title:   "test post",
				Issued:  time.Date(2014, 9, 21, 12, 58, 19, 0, time.FixedZone("JST", 9*60*60)),
				Content: "body here\nblah blah\n",
			},
			wantErr: false,
		},
		{
			name: "no content",
			args: args{
				raw: ([]byte)(``),
			},
			want:    (*postContent)(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Post{}
			got, err := p.parseRaw(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post.parseRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Post.parseRaw() diff (-got +want)\n%s", diff)
			}
		})
	}
}

func TestNewPost(t *testing.T) {
	type args struct {
		filename string
		raw      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Post
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				filename: "test.md",
				raw: ([]byte)(`{"title":"test post","date":"2014-09-21T12:58:19+09:00","tags":["golang"]}


body here
blah blah`),
			},
			want: &Post{
				Title:    "test post",
				Issued:   time.Date(2014, 9, 21, 12, 58, 19, 0, time.FixedZone("JST", 9*60*60)),
				Filename: "test.md",
				Content:  "body here\nblah blah\n",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPost(tt.args.filename, tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmpopts.IgnoreFields(Post{}, "Raw")
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewPost() diff (-got +want)\n%s", diff)
			}
		})
	}
}
