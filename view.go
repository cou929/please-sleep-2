package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/russross/blackfriday/v2"
)

// View manages views of the site
type View struct {
	c *Condition
}

// NewView initializes View
func NewView(c *Condition) *View {
	return &View{
		c: c,
	}
}

// Build builds and writes entire contents to distribute
func (v View) Build(posts Posts) error {
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
			if err := root.ExecuteTemplate(f, t.Name(), tv); err != nil {
				return fmt.Errorf("failed to execute template %s. err=%w", t.Name(), err)
			}
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
			if err := root.ExecuteTemplate(f, t.Name(), tv); err != nil {
				return fmt.Errorf("failed to execute template %s. err=%w", t.Name(), err)
			}
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
	SiteTitle     string
	PageTitle     string
	BuiltAt       time.Time
	Posts         Posts
	ArticleIndex  int
	SiteShortDesc string
}

func (v View) templateVariable(posts Posts) *templateVariable {
	return &templateVariable{
		SiteTitle:     v.c.SiteTitle,
		BuiltAt:       v.c.BuiltAt,
		Posts:         posts,
		ArticleIndex:  -1,
		SiteShortDesc: v.c.SiteShortDesc,
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
		"lastIndex": func(p Posts) int {
			return len(p) - 1
		},
		"convert": func(md string) string {
			return string(blackfriday.Run(([]byte)(md)))
		},
	}
}
