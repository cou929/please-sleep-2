package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type parseStatus int

const (
	_ parseStatus = iota
	initialized
	headerParsed
	bodyParsing
)

// ErrInvalidParseStatusTransition is error for post content parsing
var ErrInvalidParseStatusTransition = fmt.Errorf("invalid parsing status transition")

type parseState struct {
	cur parseStatus
}

func newParseState() *parseState {
	return &parseState{cur: initialized}
}

func (s *parseState) currentStatus() parseStatus {
	return s.cur
}

func (s *parseState) transit(next parseStatus) error {
	switch s.cur {
	case initialized:
		if next != headerParsed {
			return fmt.Errorf("%w cur=%v next=%v", ErrInvalidParseStatusTransition, s.cur, next)
		}
	case headerParsed:
		if next != bodyParsing {
			return fmt.Errorf("%w cur=%v next=%v", ErrInvalidParseStatusTransition, s.cur, next)
		}
	default:
		return fmt.Errorf("%w cur=%v next=%v", ErrInvalidParseStatusTransition, s.cur, next)
	}
	s.cur = next
	return nil
}

// Post represents a single article
type Post struct {
	Title    string
	Issued   time.Time
	Filename string
	Raw      []byte
	Content  string
}

type postHeader struct {
	Title  string    `json:"title"`
	Issued time.Time `json:"date"`
}

type postContent struct {
	Title   string
	Issued  time.Time
	Content string
}

// NewPost initializes Post
func NewPost(
	filename string,
	raw []byte,
) (*Post, error) {
	p := &Post{
		Filename: filename,
		Raw:      raw,
	}

	content, err := p.parseRaw(p.Raw)
	if err != nil {
		return nil, fmt.Errorf("failed to parse post. file=%s, err=%w", filename, err)
	}

	p.Title = content.Title
	p.Issued = content.Issued
	p.Content = content.Content

	return p, err
}

func (p *Post) parseRaw(raw []byte) (*postContent, error) {
	scanner := bufio.NewScanner(bytes.NewReader(raw))
	state := newParseState()
	header := &postHeader{}
	content := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		switch state.currentStatus() {
		case initialized:
			var err error
			header, err = p.parseHeader(line)
			if err != nil {
				return nil, fmt.Errorf("invalid header. err=%w", err)
			}
			if err := state.transit(headerParsed); err != nil {
				return nil, fmt.Errorf("invalid header. err=%w", err)
			}
		case headerParsed:
			if line == "" {
				continue
			}
			content = append(content, line)
			if err := state.transit(bodyParsing); err != nil {
				return nil, fmt.Errorf("invalid in header following lines. err=%w", err)
			}
		case bodyParsing:
			content = append(content, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan contents. err=%w", err)
	}

	if header.Title == "" || header.Issued.IsZero() || len(content) == 0 {
		return nil, fmt.Errorf("invalid content. Title=%v, Issued=%v, Content=%v", header.Title, header.Issued, content)
	}

	return &postContent{
		Title:   header.Title,
		Issued:  header.Issued,
		Content: strings.Join(content, "\n") + "\n",
	}, nil
}

func (p *Post) parseHeader(line string) (*postHeader, error) {
	res := &postHeader{}
	if err := json.Unmarshal(([]byte)(line), res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal header. err=%w", err)
	}

	if res.Title == "" {
		return nil, fmt.Errorf("title field is required. header=%s", line)
	}

	if res.Issued.IsZero() {
		return nil, fmt.Errorf("date field is required. header=%s", line)
	}

	return res, nil
}
