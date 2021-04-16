TEMPLATE := templates
OUTPUT := output
TEMPLATES_HTML := $(shell find templates -name "*.html")
TEMPLATES_CSS  := $(wildcard $(TEMPLATE)/*.css)
TEMPLATES_PNG  := $(wildcard $(TEMPLATE)/*.png)
TEMPLATES_TXT  := $(wildcard $(TEMPLATE)/*.txt)
OUTPUTS_HTML   := $(patsubst $(TEMPLATE)/%.html, $(OUTPUT)/%.html, $(TEMPLATES_HTML))
OUTPUTS_CSS    := $(patsubst $(TEMPLATE)/%.css, $(OUTPUT)/%.css, $(TEMPLATES_CSS))
OUTPUTS_PNG    := $(patsubst $(TEMPLATE)/%.png, $(OUTPUT)/%.png, $(TEMPLATES_PNG))
OUTPUTS_TXT    := $(patsubst $(TEMPLATE)/%.txt, $(OUTPUT)/%.txt, $(TEMPLATES_TXT))
GIT_COMMIT_HASH=$$(git rev-parse --short HEAD)

SRCS=$(wildcard cmd/*.go *.go)

all: beeceej.com $(OUTPUTS_HTML) $(OUTPUTS_CSS) $(OUTPUTS_PNG) $(OUTPUTS_TXT)

$(OUTPUT)/%.html: $(TEMPLATE)/%.html
	GIT_COMMIT_HASH=$(GIT_COMMIT_HASH) ./beeceej.com $@

$(OUTPUT)/%.css: $(TEMPLATE)/%.css
	GIT_COMMIT_HASH=$(GIT_COMMIT_HASH) ./beeceej.com $@

$(OUTPUT)/%.png: $(TEMPLATE)/%.png
	cp $(TEMPLATE)/$(notdir $@) $@

$(OUTPUT)/%.txt: $(TEMPLATE)/%.txt
	cp $(TEMPLATE)/$(notdir $@) $@


beeceej.com: $(SRCS)
	go build -o beeceej.com cmd/main.go

clean:
	rm -rf output/*
