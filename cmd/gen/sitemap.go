package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ikeikeikeike/go-sitemap-generator/v2/stm"
)

// Sitemap generates sitemap xml
type Sitemap struct {
	c *Condition
}

// NewSitemap initializes Sitemap
func NewSitemap(c *Condition) *Sitemap {
	return &Sitemap{
		c: c,
	}
}

// Build generates sitemap xml
func (s Sitemap) Build(posts Posts) error {
	sm := stm.NewSitemap(1)
	sm.Create()
	sm.SetDefaultHost(s.c.SiteURL)
	sm.Add(stm.URL{{"loc", "index.html"}})
	sm.Add(stm.URL{{"loc", "about.html"}})
	sm.Add(stm.URL{{"loc", "rss.xml"}})

	for _, p := range posts {
		sm.Add(stm.URL{{"loc", p.DestFileName()}, {"lastmod", p.Issued}})
	}

	dest := filepath.Join(s.c.DestPath, s.c.SitemapFileName)
	if err := ioutil.WriteFile(dest, ([]byte)(sm.XMLContent()), 0644); err != nil {
		return fmt.Errorf("failed to write sitemap xml to %s. err=%w", dest, err)
	}

	return nil
}
