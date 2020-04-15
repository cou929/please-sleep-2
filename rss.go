package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/gorilla/feeds"
	"github.com/russross/blackfriday/v2"
)

const feedItemNum = 10

// RSS generate rss feed xml
type RSS struct {
	c *Condition
}

// NewRSS initializes RSS
func NewRSS(c *Condition) *RSS {
	return &RSS{
		c: c,
	}
}

// Build generates rss feed xml
func (r RSS) Build(posts Posts) error {
	feed := &feeds.Feed{
		Title:       r.c.SiteTitle,
		Link:        &feeds.Link{Href: r.c.SiteURL},
		Description: r.c.SiteShortDesc,
		Author:      &feeds.Author{Name: r.c.AuthorName, Email: r.c.AuthorName},
		Created:     r.c.BuiltAt,
	}

	size := feedItemNum
	if size > len(posts) {
		size = len(posts)
	}
	items := make([]*feeds.Item, 0, size)
	for _, p := range posts[:size] {
		item := &feeds.Item{
			Title:       p.Title,
			Link:        &feeds.Link{Href: strings.Join([]string{r.c.SiteURL, p.Filename}, "")},
			Description: string(blackfriday.Run(([]byte)(p.Content))),
			Author:      &feeds.Author{Name: r.c.AuthorName, Email: r.c.AuthorName},
			Created:     p.Issued,
		}
		items = append(items, item)
	}
	feed.Items = items

	rss, err := feed.ToRss()
	if err != nil {
		return fmt.Errorf("failed to build rss. err=%w", err)
	}

	dest := filepath.Join(r.c.DestPath, r.c.RSSFileName)
	if err := ioutil.WriteFile(dest, ([]byte)(rss), 0644); err != nil {
		return fmt.Errorf("failed to write rss xml to %s. err=%w", dest, err)
	}

	return nil
}
