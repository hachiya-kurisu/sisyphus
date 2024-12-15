package main

import (
	"blekksprut.net/sisyphus"
	"flag"
	"fmt"
	"os"
)

func main() {
	v := flag.Bool("v", false, "version")
	x := flag.Bool("x", false, "use multimedia tags for links")
	f := flag.String("f", "html", "choose flavor (html/markdown)")
	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], sisyphus.Version)
		os.Exit(0)
	}

	var flavor sisyphus.Flavor
	switch *f {
	case "html":
		flavor = &sisyphus.Html{
			Extended: *x,
		}
	case "markdown":
		flavor = &sisyphus.Markdown{
			Extended: *x,
		}
	}
	sisyphus.Gem(os.Stdin, os.Stdout, flavor)
}
