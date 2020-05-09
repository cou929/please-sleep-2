package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/cou929/please-sleep-2/internal/condition"
	"github.com/cou929/please-sleep-2/internal/post"
	"github.com/russross/blackfriday/v2"
	"golang.org/x/net/html"
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

	doc, err := html.Parse(strings.NewReader(post.Content))
	if err != nil {
		log.Panicf("failed to parse html. err=%+v", err)
	}

	var buf bytes.Buffer
	if err := html.Render(&buf, doc); err != nil {
		log.Panicf("failed to render. err=%+v", err)
	}
	// ad-hoc cleaning
	cleaned := strings.Replace(strings.Replace(buf.String(), "<html><head></head><body>", "", -1), "</body></html>", "", -1)
	fmt.Println(cleaned)

	// debug
	fmt.Println("---")
	cvt := string(blackfriday.Run(([]byte)(post.Content)))
	fmt.Println(cvt)
}
