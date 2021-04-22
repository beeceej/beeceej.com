TEMPLATE        := templates
OUTPUT          := output
TPL             := $(wildcard $(TEMPLATE)/*.css) \
                   $(shell find templates -name "*.html")
TPL_OUT         := $(patsubst $(TEMPLATE)%, $(OUTPUT)%, $(TPL))
STATIC          := $(wildcard $(TEMPLATE)/*.png) \
                   $(wildcard $(TEMPLATE)/*.txt)
STATIC_OUT      := $(patsubst $(TEMPLATE)%, $(OUTPUT)%, $(STATIC))
SRCS            := $(wildcard cmd/*.go *.go)
GIT_COMMIT_HASH := $(shell git rev-parse --short HEAD)

$(OUTPUT)/%.html: beeceej.com $(TPL)
	GIT_COMMIT_HASH=$(GIT_COMMIT_HASH) ./beeceej.com $@

$(OUTPUT)/%.css: beeceej.com $(TPL)
	GIT_COMMIT_HASH=$(GIT_COMMIT_HASH) ./beeceej.com $@

$(OUTPUT)/%.png: beeceej.com $(STATIC)
	cp $(TEMPLATE)/$(notdir $@) $@

$(OUTPUT)/%.txt: beeceej.com $(STATIC)
	cp $(TEMPLATE)/$(notdir $@) $@

all: beeceej.com $(STATIC_OUT) $(TPL_OUT)

beeceej.com: $(SRCS)
	go build -o beeceej.com cmd/main.go

clean:
	rm -rf output/*
