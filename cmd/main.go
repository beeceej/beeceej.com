package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func contentPagePath(file string) string {
	return filepath.Join("content", file)
}

func commitHash() string {
	hash := os.Getenv("GIT_COMMIT_HASH")
	if hash != "" {
		return hash
	}
	panic("no commit hash")
}

type Post struct {
	Content template.HTML
	Path    string
	Title   string
}

var (
	helloWorldPost = Post{
		Path:  "notes/0-hello-world.html",
		Title: "Hello, World",
	}
	mountainGoatPost = Post{
		Path:  "notes/1-mountain-goats.html",
		Title: "Mountain Goats",
	}
	ackermanPost = Post{
		Path:  "notes/2-ackerman-function-expansions.html",
		Title: "Ackerman Function Expansions",
	}

	templates = map[string]RenderData{
		"about.html":   {PageToRender: "index.html", PageID: "about"},
		"error.html":   {PageToRender: "index.html", PageID: "error"},
		"contact.html": {PageToRender: "index.html", PageID: "contact"},
		"footer.html":  {PageToRender: "index.html", PageID: "footer"},
		"index.css":    {PageToRender: "index.css", PageID: "index-css"},
		"notes/l.html": {
			PageToRender: "index.html",
			PageID:       "notes-l",
			Other: struct{ Posts []Post }{
				Posts: []Post{ackermanPost, mountainGoatPost, helloWorldPost},
			},
		},
		"notes/0-hello-world.html": {
			ContentPagePath: contentPagePath("0-hello-world.html"),
			PageID:          "notes-note",
			PageToRender:    "index.html",
			Other:           helloWorldPost,
		},
		"notes/1-mountain-goats.html": {
			ContentPagePath: contentPagePath("1-mountain-goats.html"),
			PageID:          "notes-note",
			PageToRender:    "index.html",
			Other:           mountainGoatPost,
		},
		"notes/2-ackerman-function-expansions.html": {
			ContentPagePath: contentPagePath("2-ackerman-function-expansions.html"),
			PageID:          "notes-note",
			PageToRender:    "index.html",
			Other:           ackermanPost,
		},
	}
)

// RenderData acts as a manifest for making run time decisions on which files to render,
// and how to render a given file
type RenderData struct {
	ContentPagePath string
	CommitHash      string
	PageToRender    string
	PageID          string
	Other           interface{}
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
		postData := d.Other.(Post)
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
		_, err = t.Parse(string(templateBytes))
		return err
	})

	return root, err
}
