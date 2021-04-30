package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/beeceej/beeceej.com/posts"
)

func contentPagePath(file string) string {
	return filepath.Join("content", file)
}

func commitHash() string {
	hash := os.Getenv("GIT_COMMIT_HASH")
	if hash == "" {
		panic("no commit hash")
	}
	return hash
}

var (
	templates = map[string]RenderData{
		"robots.txt": {
			PageToRender: "robots.txt",
		},
		"about.html": {
			Description:  "about beeceej",
			Keywords:     []string{},
			PageID:       "about",
			PageToRender: "index.html",
		},
		"error.html": {
			Description:  "How'd you get here? this is an error page",
			Keywords:     []string{},
			PageID:       "error",
			PageToRender: "index.html",
		},
		"contact.html": {
			Description:  "Contact beeceej",
			Keywords:     []string{},
			PageID:       "contact",
			PageToRender: "index.html",
		},
		"index.css": {
			Keywords:     []string{},
			PageID:       "index-css",
			PageToRender: "index.css",
		},
		"notes/l.html": {
			Keywords:     []string{},
			PageID:       "notes-l",
			PageToRender: "index.html",
			Other: struct{ Posts []posts.Post }{
				Posts: posts.Posts,
			},
		},
		"notes/0-hello-world.html": {
			ContentPagePath: contentPagePath("0-hello-world.html"),
			Description:     "First blog post",
			Keywords: []string{
				"programming",
				"Hello world",
				"simple",
				"frontend",
				"blog",
				"go",
				"golang",
			},
			Other:        posts.HelloWorld,
			PageID:       "notes-note",
			PageToRender: "index.html",
		},
		"notes/1-mountain-goats.html": {
			ContentPagePath: contentPagePath("1-mountain-goats.html"),
			Description:     "Mountain Goat escapes enclosure, bound for greener greener grass, and taller mountains",
			Keywords: []string{
				"programming",
				"allegory",
				"goat",
				"cloud",
				"mountain",
			},
			Other:        posts.MountainGoat,
			PageID:       "notes-note",
			PageToRender: "index.html",
		},
		"notes/2-ackerman-function-expansions.html": {
			ContentPagePath: contentPagePath("2-ackerman-function-expansions.html"),
			Description:     "Ackermann function expansion in scheme",
			Keywords: []string{
				"programming",
				"lisp",
				"guile",
				"scheme",
				"sicp",
				"mit",
			},
			Other:        posts.Ackermann,
			PageID:       "notes-note",
			PageToRender: "index.html",
		},
		"notes/3-good-advice.html": {
			ContentPagePath: contentPagePath("3-good-advice.html"),
			Description:     "Guy Clark on good advice",
			Keywords: []string{
				"Guy Clark",
				"outlaw",
				"music",
				"random",
				"thoughts",
				"mit",
				"texas",
				"country",
			},
			Other:        goodAdvice,
			PageID:       "notes-note",
			PageToRender: "index.html",
		},
	}
)

// RenderData acts as a manifest for making run time decisions on which files to render,
// and how to render a given file
type RenderData struct {
	CommitHash      string
	ContentPagePath string
	Description     string
	Keywords        []string
	Other           interface{}
	PageID          string
	PageToRender    string
}

func main() {
	filename := filepath.Base(os.Args[1])
	filename = filepath.Join(strings.Split(filepath.FromSlash(os.Args[1]), string(os.PathSeparator))[1:]...)
	path := filepath.Dir(os.Args[1])
	if err := os.MkdirAll(path, 0755); err != nil && !os.IsExist(err) {
		panic(err.Error())
	}
	must(makeOutputFile(filename))
}

func augmentRenderData(d *RenderData) error {
	if d == nil {
		return fmt.Errorf("render data is nil")
	}
	if d.ContentPagePath != "" {
		f, err := os.Open(d.ContentPagePath)
		if err != nil {
			return err
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		postData := d.Other.(posts.Post)
		postData.Content = template.HTML(string(b))
		d.Other = postData
	}
	d.CommitHash = commitHash()
	return nil
}

func makeOutputFile(path string) error {
	var (
		f   *os.File
		err error
		tpl *template.Template
	)
	renderData, exists := templates[path]
	if !exists {
		return nil
	}
	if err = augmentRenderData(&renderData); err != nil {
		return err
	}

	if renderData.PageToRender != "" {
		if f, err = os.Create(filepath.Join("output", path)); err != nil {
			return err
		}
		defer func() {
			if err = f.Close(); err != nil {
				panic(err.Error())
			}
		}()
	}

	if renderData.PageToRender != "" {
		tpl = template.Must(findAndParseTemplates("templates", nil))
		return tpl.Lookup(renderData.PageToRender).Execute(f, renderData)
	}
	return nil
}

func must(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err.Error())
		}
	}
}

func findAndParseTemplates(rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	root := template.New("")
	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, walkErr error) (err error) {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".png" {
			return nil
		}

		var templateBytes []byte
		if templateBytes, err = ioutil.ReadFile(path); err != nil {
			return err
		}

		name := path[len(cleanRoot)+1:]
		t := root.New(name).Funcs(funcMap)
		t.Funcs(template.FuncMap{"JoinStr": strings.Join})
		_, err = t.Parse(string(templateBytes))
		return err
	})

	return root, err
}
