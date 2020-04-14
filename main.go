package main

import (
	"log"
	"time"
)

// Condition holds configuration of entire site
type Condition struct {
	PostPath   string
	ViewPath   string
	ViewSuffix string
	DestPath   string
	SiteTitle  string
	BuiltAt    time.Time
}

// NewCondition initializes Condition
func NewCondition() *Condition {
	return &Condition{
		PostPath:   "post",
		ViewPath:   "view",
		ViewSuffix: ".html",
		DestPath:   "dist",
		SiteTitle:  "Please Sleep",
		BuiltAt:    time.Now(),
	}
}

func main() {
	c := NewCondition()

	repo := NewPostRepository(c)
	posts, err := repo.List()
	if err != nil {
		log.Panicf("failed to list posts. err=%+v", err)
	}

	view, err := NewView(c)
	if err != nil {
		log.Panicf("failed to NewView. err=%+v", err)
	}

	if err := view.Build(posts); err != nil {
		log.Panicf("failed to build view. err=%+v", err)
	}

	log.Println("finished")
}
