SRC = $(shell find . -name "*.go")

# Credit to https://github.com/commissure/go-git-build-vars for giving me a starting point for this.
BUILD_TIME = `date +%Y%m%d%H%M%S`
GIT_REVISION = `git rev-parse --short HEAD`
GIT_BRANCH = `git rev-parse --symbolic-full-name --abbrev-ref HEAD | sed 's/\//-/g'`
GIT_DIRTY = `git diff-index --quiet HEAD -- || echo 'x-'`

LDFLAGS = -ldflags "-s -X main.BuildTime=${BUILD_TIME} -X main.GitRevision=${GIT_DIRTY}${GIT_REVISION} -X main.GitBranch=${GIT_BRANCH}"

.PHONY: all
all: bin/a5srv

bin/a5srv: $(foreach f, $(SRC), $(f))
	go build ${LDFLAGS} -o bin/a5srv cmd/a5srv/main.go

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: clean
clean:
	rm -rf bin/