package main

// View manages views of the site
type View struct{}

// NewView initializes View
func NewView(c *Condition) (*View, error) {
	return nil, nil
}

// Build builds and writes entire contents to distribute
func (v View) Build(posts []*Post) error {
	return nil
}
