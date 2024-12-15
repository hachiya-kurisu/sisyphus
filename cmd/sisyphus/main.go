package main

import (
	"blekksprut.net/sisyphus"
	"flag"
	"os"
)

func main() {
	f := flag.String("f", "html", "choose flavor (html/markdown)")
	flag.Parse()

	var flavor sisyphus.Flavor
	switch *f {
	case "html":
		flavor = &sisyphus.Html{}
	case "markdown":
		flavor = &sisyphus.Markdown{}
	}
	sisyphus.Gem(os.Stdin, os.Stdout, flavor)
}
