package condition

import "time"

// Condition holds configuration of entire site
type Condition struct {
	PostPath        string
	ViewPath        string
	ViewSuffix      string
	DestPath        string
	SiteTitle       string
	BuiltAt         time.Time
	SiteURL         string
	SiteShortDesc   string
	AuthorName      string
	AuthorMail      string
	RSSFileName     string
	TwitterAccount  string
	SitemapFileName string
}

// NewCondition initializes Condition
func NewCondition() *Condition {
	return &Condition{
		PostPath:        "post",
		ViewPath:        "view",
		ViewSuffix:      ".html",
		DestPath:        "dist",
		SiteTitle:       "Please Sleep",
		BuiltAt:         time.Now(),
		SiteURL:         "https://please-sleep.cou929.nu/",
		SiteShortDesc:   "From notes on my laptop",
		AuthorName:      "Kosei Moriyama",
		AuthorMail:      "cou929@gmail.com",
		RSSFileName:     "rss.xml",
		TwitterAccount:  "@cou929",
		SitemapFileName: "sitemap.xml",
	}
}
