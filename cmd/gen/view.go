package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/cou929/please-sleep-2/internal/condition"
	"github.com/cou929/please-sleep-2/internal/post"
	"github.com/russross/blackfriday/v2"
)

const ogDescLen = 300

// View manages views of the site
type View struct {
	c *condition.Condition
}

// NewView initializes View
func NewView(c *condition.Condition) *View {
	return &View{
		c: c,
	}
}

// Build builds and writes entire contents to distribute
func (v View) Build(posts post.Posts) error {
	root, err := v.prepareTemplates(v.c.ViewPath, v.c.ViewSuffix, v.viewFunc())
	if err != nil {
		return fmt.Errorf("failed to prepare templates. err=%w", err)
	}
	if err := v.ensureDestDir(v.c.DestPath); err != nil {
		return fmt.Errorf("failed to ensure dest dir %s. err=%w", v.c.DestPath, err)
	}

	for _, t := range root.Templates() {
		if !v.isDriver(t.Name()) {
			continue
		}

		if !v.isPostTemplate(t.Name()) {
			dest := filepath.Join(v.c.DestPath, t.Name())
			f, err := os.Create(dest)
			if err != nil {
				return fmt.Errorf("failed to create %s. err=%w", dest, err)
			}
			defer f.Close()
			tv := v.templateVariable(posts)
			tv.OgType = "website"
			urlStr := v.c.SiteURL + t.Name()
			pageURL, err := url.Parse(urlStr)
			if err != nil {
				return fmt.Errorf("failed to parse url %s. err=%w", urlStr, err)
			}
			tv.URL = pageURL
			if err := root.ExecuteTemplate(f, t.Name(), tv); err != nil {
				return fmt.Errorf("failed to execute template %s. err=%w", t.Name(), err)
			}
			f.Close()
			continue
		}

		for i, p := range posts {
			dest := filepath.Join(v.c.DestPath, p.DestFileName())
			f, err := os.Create(dest)
			if err != nil {
				return fmt.Errorf("failed to create %s. err=%w", dest, err)
			}
			defer f.Close()

			tv := v.templateVariable(posts)
			tv.ArticleIndex = i
			tv.PageTitle = p.Title
			tv.OgType = "article"
			urlStr := v.c.SiteURL + p.DestFileName()
			pageURL, err := url.Parse(urlStr)
			if err != nil {
				return fmt.Errorf("failed to parse url %s. err=%w", urlStr, err)
			}
			tv.URL = pageURL
			if p.OgImage != "" {
				tv.OgImage = p.OgImage
			}
			if err := root.ExecuteTemplate(f, t.Name(), tv); err != nil {
				return fmt.Errorf("failed to execute template %s. err=%w", t.Name(), err)
			}
			f.Close()
		}
	}

	return err
}

func (v View) prepareTemplates(rootDir string, suffix string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, suffix) {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			t := root.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				return e2
			}
		}

		return nil
	})

	return root, err
}

func (v View) isDriver(templateName string) bool {
	return filepath.Base(templateName) == templateName
}

func (v View) isPostTemplate(templateName string) bool {
	return templateName == "post.html"
}

func (v View) ensureDestDir(dirName string) error {
	if err := os.RemoveAll(dirName); err != nil {
		return fmt.Errorf("failed to remove dest dir %s. err=%w", dirName, err)
	}
	err := os.Mkdir(dirName, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to make dest dir %s. err=%w", dirName, err)
	}
	return nil
}

type templateVariable struct {
	SiteTitle      string
	PageTitle      string
	BuiltAt        time.Time
	Posts          post.Posts
	ArticleIndex   int
	SiteShortDesc  string
	OgType         string
	URL            *url.URL
	TwitterAccount string
	OgImage        string
}

func (v View) templateVariable(posts post.Posts) *templateVariable {
	return &templateVariable{
		SiteTitle:      v.c.SiteTitle,
		BuiltAt:        v.c.BuiltAt,
		Posts:          posts,
		ArticleIndex:   -1,
		SiteShortDesc:  v.c.SiteShortDesc,
		TwitterAccount: v.c.TwitterAccount,
		OgImage:        v.c.DefaultOgImage,
	}
}

func (v View) viewFunc() template.FuncMap {
	return template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"dec": func(i int) int {
			return i - 1
		},
		"lastIndex": func(p post.Posts) int {
			return len(p) - 1
		},
		"convert": func(md string) string {
			return string(blackfriday.Run(([]byte)(md)))
		},
		"shorten": func(content string) string {
			parser := blackfriday.New()
			ast := parser.Parse([]byte(content))
			texts := ""
			ast.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
				if node.Type == blackfriday.Text || node.Type == blackfriday.Code {
					texts = strings.Join([]string{texts, string(node.Literal)}, "")
				}
				r := []rune(texts)
				if len(r) > ogDescLen {
					return blackfriday.Terminate
				}
				return blackfriday.GoToNext
			})

			r := []rune(texts)
			l := ogDescLen
			suf := "…"
			if len(r) < ogDescLen {
				l = len(r)
				suf = ""
			}
			return strings.Replace(strings.Replace(string(r[0:l]), "\n", " ", -1), `"`, `&quot;`, -1) + suf
		},
	}
}
