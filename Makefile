
.PHONY: all
all: jcheck

.PHONY: jcheck
jcheck:
	go build jcheck.go

.PHONY: deps
deps:
	go get github.com/mattn/go-sqlite3
	go get github.com/pkg/errors

.PHONY: deptools
deptools:
	go get github.com/ChimeraCoder/gojson/gojson
