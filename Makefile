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
	gofmt -s -w *.go
	gofmt -s -w cmd/sisyphus/main.go

doc: README.md

README.md: README.gmi
	./sisyphus -f markdown <README.gmi >README.md

release: push
	git push github --tags

