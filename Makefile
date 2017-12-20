

.PHONY: deps
deps:
	go get github.com/mattn/go-sqlite3 

.PHONY: deptools
deptools:
	go get github.com/ChimeraCoder/gojson/gojson
