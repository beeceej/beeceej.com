package main

import (
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	filename := filepath.Base(os.Args[1])
	filename = filepath.Join(
		strings.Split(filepath.FromSlash(os.Args[1]), string(os.PathSeparator))[1:]...)
	path := filepath.Dir(os.Args[1])
	if err := os.MkdirAll(path, 0755); err != nil && !os.IsExist(err) {
		panic(err.Error())
	}
	must(makeOutputFile(filename))
}

// RenderData acts as a manifest for making run time decisions on which files to render,
// and how to render a given file
type RenderData struct {
	PageToRender string
	PageID       string
}

var templates = map[string]RenderData{
	"about.html":   {PageToRender: "index.html", PageID: "about"},
	"error.html":   {PageToRender: "index.html", PageID: "error"},
	"contact.html": {PageToRender: "index.html", PageID: "contact"},
	"footer.html":  {PageToRender: "index.html"},
	"index.css":    {PageToRender: "index.css"},
	"nav.html":     {PageToRender: "", PageID: "index.html"},
	"notes/l.html": {PageToRender: "index.html", PageID: "notes-l"},
}

func makeOutputFile(path string) error {
	var (
		f   *os.File
		err error
		tpl *template.Template
	)
	renderData, exists := templates[path]
	if exists && renderData.PageToRender != "" {
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
