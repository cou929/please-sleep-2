package main

import (
	"log"
	"time"
)

// Condition holds configuration of entire site
type Condition struct {
	PostPath      string
	ViewPath      string
	ViewSuffix    string
	DestPath      string
	SiteTitle     string
	BuiltAt       time.Time
	SiteURL       string
	SiteShortDesc string
	AuthorName    string
	AuthorMail    string
	RSSFileName   string
}

// NewCondition initializes Condition
func NewCondition() *Condition {
	return &Condition{
		PostPath:      "post",
		ViewPath:      "view",
		ViewSuffix:    ".html",
		DestPath:      "dist",
		SiteTitle:     "Please Sleep",
		BuiltAt:       time.Now(),
		SiteURL:       "https://please-sleep.cou929.nu/",
		SiteShortDesc: "From notes on my laptop",
		AuthorName:    "Kosei Moriyama",
		AuthorMail:    "cou929@gmail.com",
		RSSFileName:   "rss.xml",
	}
}

func main() {
	c := NewCondition()

	repo := NewPostRepository(c)
	posts, err := repo.List()
	if err != nil {
		log.Panicf("failed to list posts. err=%+v", err)
	}

	view := NewView(c)
	if err := view.Build(posts); err != nil {
		log.Panicf("failed to build view. err=%+v", err)
	}

	rss := NewRSS(c)
	if err := rss.Build(posts); err != nil {
		log.Panicf("failed to build rss. err=%+v", err)
	}

	log.Println("finished")
}
