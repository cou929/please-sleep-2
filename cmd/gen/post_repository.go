package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/cou929/please-sleep-2/internal/condition"
	"github.com/cou929/please-sleep-2/internal/post"
)

const postSuffix = ".md"

// PostRepository loads and parses blog posts
type PostRepository struct {
	c      *condition.Condition
	posts  post.Posts
	reader reader
}

type fileInfo interface {
	Name() string
	IsDir() bool
}

type reader interface {
	ReadDir(dirname string) ([]fileInfo, error)
	ReadFile(filename string) ([]byte, error)
}

type ioUtil struct{}

func (i ioUtil) ReadDir(dirname string) ([]fileInfo, error) {
	info, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, fmt.Errorf("failed to ioutil.ReadDir. dirname=%s, err=%w", dirname, err)
	}

	res := make([]fileInfo, 0, len(info))
	for _, f := range info {
		res = append(res, fileInfo(f))
	}

	return res, nil
}

func (i ioUtil) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// NewPostRepository initializes PostRepository
func NewPostRepository(c *condition.Condition) *PostRepository {
	return &PostRepository{
		c:      c,
		reader: ioUtil{},
	}
}

// List retrieves all posts
func (r *PostRepository) List() (post.Posts, error) {
	if len(r.posts) > 0 {
		return r.posts, nil
	}

	posts, err := r.load()
	if err != nil {
		return nil, fmt.Errorf("failed to load posts. err=%w", err)
	}
	sort.Sort(posts)
	r.posts = posts

	return r.posts, nil
}

func (r *PostRepository) load() (post.Posts, error) {
	files, err := r.reader.ReadDir(r.c.PostPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir. path=%s, err=%w", r.c.PostPath, err)
	}

	res := make(post.Posts, 0, len(files))

	for _, file := range files {
		if !r.isTarget(file) {
			continue
		}
		f := filepath.Join(r.c.PostPath, file.Name())
		content, err := r.reader.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("failed to read file. file=%s, err=%w", f, err)
		}
		post, err := post.NewPost(file.Name(), content, r.c)
		if err != nil {
			return nil, fmt.Errorf("failed to NewPost. filename=%s, err=%w", file.Name(), err)
		}
		res = append(res, post)
	}

	return res, nil
}

func (r *PostRepository) isTarget(f fileInfo) bool {
	if f.IsDir() {
		return false
	}
	if filepath.Ext(f.Name()) != postSuffix {
		return false
	}

	return true
}
