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
	}

	Posts = []Post{MatchingParens, GoodAdvice, Ackermann, MountainGoat, HelloWorld}
)
