TEMPLATE        := templates
OUTPUT          := output
CONTENT         := content
TPL             := $(wildcard $(TEMPLATE)/*.css) \
                   $(shell find templates -name "*.html")
TPL_OUT         := $(patsubst $(TEMPLATE)%, $(OUTPUT)%, $(TPL))
STATIC          := $(wildcard $(TEMPLATE)/*.png) \
                   $(wildcard $(TEMPLATE)/*.webp) \
                   $(wildcard $(TEMPLATE)/*.jpg) \
                   $(wildcard $(TEMPLATE)/*.txt)
STATIC_OUT      := $(patsubst $(TEMPLATE)%, $(OUTPUT)%, $(STATIC))
SRCS            := $(wildcard cmd/*.go *.go posts/*.go)
GIT_COMMIT_HASH := $(shell git rev-parse --short HEAD)

TEMPLATIZE := GIT_COMMIT_HASH=$(GIT_COMMIT_HASH) ./beeceej.com

$(OUTPUT)/%.html: beeceej.com $(TPL)
	$(TEMPLATIZE) $@

$(OUTPUT)/%.css: beeceej.com $(TPL)
	$(TEMPLATIZE) $@

$(OUTPUT)/%.png: beeceej.com $(STATIC)
	cp $(TEMPLATE)/$(notdir $@) $@

$(OUTPUT)/%.jpg: beeceej.com $(STATIC)
	cp $(TEMPLATE)/$(notdir $@) $@

$(OUTPUT)/%.webp: beeceej.com $(STATIC)
	cp $(TEMPLATE)/$(notdir $@) $@

$(OUTPUT)/%.txt: beeceej.com $(STATIC)
	cp $(TEMPLATE)/$(notdir $@) $@

all: beeceej.com $(TPL_OUT) $(STATIC_OUT)

beeceej.com: $(SRCS)
	go build -o beeceej.com cmd/main.go

clean:
	rm -rf output/*
