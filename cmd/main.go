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

type Data struct {
	PageToRender string
	PageID       string
}

var templates = map[string]Data{
	"about.html":   Data{PageToRender: "index.html", PageID: "about"},
	"notes/l.html": Data{PageToRender: "index.html", PageID: "notes-l"},
	"contact.html": Data{PageToRender: "index.html", PageID: "contact"},
	"nav.html":     Data{PageToRender: "", PageID: "index.html"},
	"index.css":    Data{PageToRender: "index.css"},
	"footer.html":  Data{PageToRender: "index.html"},
}

func makeOutputFile(path string) error {
	var (
		f   *os.File
		err error
		tpl *template.Template
	)
	if f, err = os.Create(filepath.Join("output", path)); err != nil {
		return err
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err.Error())
		}
	}()

	renderData := templates[path]
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
