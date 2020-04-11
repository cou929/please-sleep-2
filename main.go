package main

import (
	"log"
)

// Condition holds configuration of entire site
type Condition struct{}

// NewCondition initializes Condition
func NewCondition() *Condition {
	return &Condition{}
}

func main() {
	c := NewCondition()

	repo, err := NewPostRepository(c)
	if err != nil {
		log.Panicf("failed to NewPostRepository. err=%+v", err)
	}

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
