package main

import "time"

// Post represents a single article
type Post struct {
	Title    string
	Issued   time.Time
	Filename string
	Raw      []byte
	Content  string
}

// NewPost initializes Post
func NewPost(
	filename string,
	raw []byte,
) *Post {
	return &Post{
		Filename: filename,
		Raw:      raw,
	}
}
