package posts

import "html/template"

type Post struct {
	Content template.HTML
	Path    string
	Title   string
	Posted  string
	Updated string
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
	}
	Ackermann = Post{
		Path:    "notes/2-ackerman-function-expansions.html",
		Title:   "Ackermann Function Expansions",
		Posted:  "2021-04-16",
		Updated: "2021-04-16",
	}
	GoodAdvice = Post{
		Path:    "notes/3-good-advice.html",
		Title:   "Guy Clark On Good Advice",
		Posted:  "2021-04-19",
		Updated: "2021-04-20",
	}
	MatchingParens = Post{
		Path:    "notes/4-matching-parens.html",
		Title:   "Matching Parens",
		Posted:  "2021-04-30",
		Updated: "2021-04-30",
	}
	PeppersIntro = Post{
		Path:    "notes/5-peppers-intro.html",
		Title:   "Peppers Intro",
		Posted:  "2021-08-01",
		Updated: "2021-08-01",
	}
	MutableCallArgsInPython = Post{
		Path:    "notes/6-mutable-call-args-in-python.html",
		Title:   "Mutable Call Arguments in Python",
		Posted:  "2021-08-25",
		Updated: "2021-08-25",
	}
	PeppersPartTwo = Post{
		Path:    "notes/7-peppers-part-2.html",
		Title:   "Peppers Part 2: The Flowering",
		Posted:  "2021-09-22",
		Updated:  "2021-09-24",
	}
	ModernFrontendProblems = Post{
		Path:    "notes/8-the-problem-with-modern-web-development.html",
		Title:   "The Problem With Modern Web Development",
		Posted:  "2021-09-22",
		Updated:  "2021-09-22",
	}
	CoffeeFromAnOldCoworker = Post{
		Path:    "notes/9-maple-leaf-coffee-nicaraguan-el-finca.html",
		Title:   "Coffee From an Old Coworker",
		Posted:  "2021-10-01",
		Updated:  "2021-10-01",
	}

	Posts = []Post{
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
