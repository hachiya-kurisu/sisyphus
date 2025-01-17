all: sisyphus

again: clean all

sisyphus: html.go markdown.go sisyphus.go cmd/sisyphus/main.go
	go build -o sisyphus cmd/sisyphus/main.go

clean:
	rm -f sisyphus

test:
	go test -cover

push:
	got send
	git push github

fmt:
	gofmt -s -w *.go cmd/*/*.go

cover:
	go test -coverprofile=cover.out
	go tool cover -html cover.out

doc: README.md

README.md: README.gmi INSTALL.gmi
	cat README.gmi INSTALL.gmi | sisyphus -f markdown > README.md

release: push
	git push github --tags

