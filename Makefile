
PACKAGE=shootora-voting
SRC=main.go
GOPATH=$(PWD)

.PHONY: all init compile run clean clean-all

all: init compile

init: main.db src/github.com

compile:
	GOPATH=$(GOPATH) go build $(PACKAGE)

run: compile
	./$(PACKAGE)

main.db: init.sql
	sqlite3 -init init.sql main.db ""

src/github.com:
	GOPATH=$(GOPATH) go get ...

clean:
	rm -rf $(PACKAGE)

clean-all: clean
	rm -rf main.db
