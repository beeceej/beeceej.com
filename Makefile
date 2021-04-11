TEMPLATE := templates
OUTPUT := output
TEMPLATES_HTML := $(shell find templates -name "*.html")
TEMPLATES_CSS  := $(wildcard $(TEMPLATE)/*.css)
TEMPLATES_PNG  := $(wildcard $(TEMPLATE)/*.png)
OUTPUTS_HTML   := $(patsubst $(TEMPLATE)/%.html, $(OUTPUT)/%.html, $(TEMPLATES_HTML))
OUTPUTS_CSS    := $(patsubst $(TEMPLATE)/%.css, $(OUTPUT)/%.css, $(TEMPLATES_CSS))
OUTPUTS_PNG    := $(patsubst $(TEMPLATE)/%.png, $(OUTPUT)/%.png, $(TEMPLATES_PNG))

SRCS=$(wildcard cmd/*.go *.go)

all: beeceej.com $(OUTPUTS_HTML) $(OUTPUTS_CSS) $(OUTPUTS_PNG)

$(OUTPUT)/%.html: $(TEMPLATE)/%.html
	./beeceej.com $@

$(OUTPUT)/%.css: $(TEMPLATE)/%.css
	./beeceej.com $@

$(OUTPUT)/%.png: $(TEMPLATE)/%.png
	cp $(TEMPLATE)/$(notdir $@) $@


beeceej.com: $(SRCS)
	go build -o beeceej.com cmd/main.go

clean:
	rm -rf output/*
