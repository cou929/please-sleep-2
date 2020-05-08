package main

import (
	"log"

	"github.com/cou929/please-sleep-2/internal/condition"
)

func main() {
	c := condition.NewCondition()

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

	sitemap := NewSitemap(c)
	if err := sitemap.Build(posts); err != nil {
		log.Panicf("failed to build sitemap. err=%+v", err)
	}

	log.Println("finished")
}
