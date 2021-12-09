package posts

import (
	"fmt"
	"html/template"
	"strings"
)

type Post struct {
	Content template.HTML
	Path    string
	Title   string
	Posted  string
	Updated string
	Tags    []string
}

func (p Post) RenderTags() string {
	if p.Tags == nil {
		return ""
	}

	return fmt.Sprintf("[%s]", strings.Join(p.Tags, ", "))
}

var (
	HelloWorld = Post{
		Path:   "notes/0-hello-world.html",
		Title:  "Hello, World",
		Posted: "2021-04-11",
	}
	MountainGoat = Post{
		Path:   "notes/1-mountain-goats.html",
		Title:  "Mountain Goats",
		Posted: "2021-04-14",
		Tags:   []string{"allegory"},
	}
	Ackermann = Post{
		Path:    "notes/2-ackerman-function-expansions.html",
		Title:   "Ackermann Function Expansions",
		Posted:  "2021-04-16",
		Updated: "2021-04-16",
		Tags:    []string{"programming", "lisp"},
	}
	GoodAdvice = Post{
		Path:    "notes/3-good-advice.html",
		Title:   "Guy Clark On Good Advice",
		Posted:  "2021-04-19",
		Updated: "2021-04-20",
		Tags:    []string{"music"},
	}
	MatchingParens = Post{
		Path:    "notes/4-matching-parens.html",
		Title:   "Matching Parens",
		Posted:  "2021-04-30",
		Updated: "2021-04-30",
		Tags:    []string{"programming", "lisp", "javascript"},
	}
	PeppersIntro = Post{
		Path:    "notes/5-peppers-intro.html",
		Title:   "Peppers Intro",
		Posted:  "2021-08-01",
		Updated: "2021-10-15",
		Tags:    []string{"peppers", "gardening"},
	}
	MutableCallArgsInPython = Post{
		Path:    "notes/6-mutable-call-args-in-python.html",
		Title:   "Mutable Call Arguments in Python",
		Posted:  "2021-08-25",
		Updated: "2021-12-05",
		Tags:    []string{"programming"},
	}
	PeppersPartTwo = Post{
		Path:    "notes/7-peppers-part-2.html",
		Title:   "Peppers Part 2: The Flowering",
		Posted:  "2021-09-22",
		Updated: "2021-10-15",
		Tags:    []string{"peppers", "gardening"},
	}
	ModernFrontendProblems = Post{
		Path:    "notes/8-the-problem-with-modern-web-development.html",
		Title:   "The Problem With Modern Web Development",
		Posted:  "2021-09-22",
		Updated: "2021-10-09",
		Tags:    []string{"programming"},
	}
	CoffeeFromAnOldCoworker = Post{
		Path:    "notes/9-maple-leaf-coffee-nicaraguan-el-finca.html",
		Title:   "Coffee From an Old Coworker",
		Posted:  "2021-10-01",
		Updated: "2021-10-09",
		Tags:    []string{"coffee"},
	}
	PeppersPartThree = Post{
		Path:    "notes/10-peppers-part-3-fruits.html",
		Title:   "Peppers Part 3: Fruits",
		Posted:  "2021-10-10",
		Updated: "2021-10-15",
		Tags:    []string{"peppers", "gardening"},
	}
	AlacrittySpawnNewWindow = Post{
		Path:    "notes/11-alacritty-spawn-new-window.html",
		Title:   "Spawning New Windows in Alacritty (On Fedora 35)",
		Posted:  "2021-12-04",
		Updated: "2021-12-07",
		Tags:    []string{"programming"},
	}
	TicTacToe = Post{
		Path:    "notes/12-tic-tac-toe.html",
		Title:   "Tic Tac Toe",
		Posted:  "2021-12-08",
		Updated: "2021-12-08",
		Tags:    []string{"game", "javascript"},
	}

	Posts = []Post{
		TicTacToe,
		AlacrittySpawnNewWindow,
		PeppersPartThree,
		CoffeeFromAnOldCoworker,
		ModernFrontendProblems,
		PeppersPartTwo,
		MutableCallArgsInPython,
		PeppersIntro,
		MatchingParens,
		GoodAdvice,
		Ackermann,
		MountainGoat,
		HelloWorld}
)
