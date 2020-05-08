package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/cou929/please-sleep-2/internal/condition"
	"github.com/cou929/please-sleep-2/internal/post"
)

func main() {
	if len(os.Args) != 2 {
		log.Panicf("only one target is required")
	}
	file := os.Args[1]
	if file == "" {
		log.Panicf("target is required")
	}
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicf("failed to read file %s. err=%+v", file, err)
	}

	c := condition.NewCondition()

	post, err := post.NewPost(file, body, c)
	if err != nil {
		log.Panicf("failed to load post %s. err=%+v", file, err)
	}

	log.Println(post)

	log.Println("finished")
}
