package main

import (
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
			Other:        posts.GoodAdvice,
			PageID:       "notes-note",
			PageToRender: "index.html",
		},
		"notes/4-matching-parens.html": {
			ContentPagePath: contentPagePath("4-matching-parens.html"),
			Description:     "javascript demo of parentheses matching",
			Keywords:        []string{},
			Other:           posts.MatchingParens,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/5-peppers-intro.html": {
			ContentPagePath: contentPagePath("5-peppers-intro.html"),
			Description:     "Log of growing peppers entry number 1",
			Keywords:        []string{"Hot pepper", "capsaicin", "Trinidad Moruga Scorpion", "Aerogarden"},
			Other:           posts.PeppersIntro,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/6-mutable-call-args-in-python.html": {
			ContentPagePath: contentPagePath("6-mutable-call-args-in-python.html"),
			Description:     "How Mutable arguments in python caused test assertions to behave unexpectedly",
			Keywords:        []string{"python", "mutability"},
			Other:           posts.MutableCallArgsInPython,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/7-peppers-part-2.html": {
			ContentPagePath: contentPagePath("7-peppers-part-2.html"),
			Description:     "Log of growing peppers entry number 2",
			Keywords:        []string{"Hot pepper", "capsaicin", "Trinidad Moruga Scorpion", "Aerogarden"},
			Other:           posts.PeppersPartTwo,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/8-the-problem-with-modern-web-development.html": {
			ContentPagePath: contentPagePath("8-the-problem-with-modern-web-development.html"),
			Description:     "Maintainability of large javascript based web projects over time is called into question",
			Keywords:        []string{"React", "Javascript", "modern", "frontend", "web", "development"},
			Other:           posts.ModernFrontendProblems,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/9-maple-leaf-coffee-nicaraguan-el-finca.html": {
			ContentPagePath: contentPagePath("9-maple-leaf-coffee-nicaraguan-el-finca.html"),
			Description:     "Maple leaf Coffee Roasters, Nicaraguan El Finca Review",
			Keywords:        []string{"coffee", "light roast", "aeropress", "light roast", "brew", "maple leaf coffee roasters"},
			Other:           posts.CoffeeFromAnOldCoworker,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/10-peppers-part-3-fruits.html": {
			ContentPagePath: contentPagePath("10-peppers-part-3-fruits.html"),
			Description:     "Log of growing peppers entry number 3",
			Keywords:        []string{"Hot pepper", "capsaicin", "Trinidad Moruga Scorpion", "Aerogarden"},
			Other:           posts.PeppersPartThree,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/11-alacritty-spawn-new-window.html": {
			ContentPagePath: contentPagePath("11-alacritty-spawn-new-window.html"),
			Description:     "Spawning New Windows on Fedora 35 with Alacritty 0.9.0",
			Keywords:        []string{"Terminal Emulator", "alacritty", "linux"},
			Other:           posts.AlacrittySpawnNewWindow,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/12-tic-tac-toe.html": {
			ContentPagePath: contentPagePath("12-tic-tac-toe.html"),
			Description:     "tic-tac-toe game",
			Keywords:        []string{"game", "tic-tac-toe", "javsacript"},
			Other:           posts.TicTacToe,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/13-peppers-part-four-end-of-a-season.html": {
			ContentPagePath: contentPagePath("13-peppers-part-four-end-of-a-season.html"),
			Description:     "Peppers End of 2021",
			Keywords:        []string{"Hot pepper", "capsaicin", "Trinidad Moruga Scorpion", "Aerogarden"},
			Other:           posts.PeppersPartFour,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/14-khao-piak-sen.html": {
			ContentPagePath: contentPagePath("14-khao-piak-sen.html"),
			Description:     "Lao Chicken Noodle Soup",
			Keywords:        []string{"Lao", "Noodle", "Soup", "chicken"},
			Other:           posts.KhaoPiakSen,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
		"notes/15-lgtm.html": {
			ContentPagePath: contentPagePath("15-lgtm.html"),
			Description:     "LGTM a Silly OCAML github action",
			Keywords:        []string{"ocaml", "oss", "github", "github action"},
			Other:           posts.LGTM,
			PageID:          "notes-note",
			PageToRender:    "index.html",
		},
	}
)

// RenderData acts as a manifest for making run time decisions on which files to render,
// and how to render a given file
type RenderData struct {
	// CommitHash is the commit hash of the repository, used for
	// exposing the version of the statically generated site.
	CommitHash string

	// ContentPagePath is the path of the content page
	// This page will be loaded in if requested, and used as input into a template file
	ContentPagePath string

	// Description is the description of a page, inserted into the description meta tag
	Description string

	// Keywords is a set of keywords inserted into the keywords meta tag
	Keywords []string

	// Other is an arbitrary structure used for rendering dynamic pages
	Other interface{}

	// PageID is used for the templating engine to
	// make runtime decisions on which template to render
	PageID string

	// PageToRender is the name of template to use when rendering.
	PageToRender string
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
	if d == nil { // No RenderData, so nothing to do
		return nil
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

// findAndParseTemplates walks the file system recursively parsing templates as it goes.
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
		if filepath.Ext(path) == ".png" || filepath.Ext(path) == ".webp" || filepath.Ext(path) == ".jpg" {
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
