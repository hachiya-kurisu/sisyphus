package main

import (
	"blekksprut.net/sisyphus"
	"flag"
	"fmt"
	"os"
)

func main() {
	v := flag.Bool("v", false, "version")
	x := flag.Bool("x", false, "allow inline media")
	a := flag.String("a", "", "aspeq prefix")
	f := flag.String("f", "html", "flavor (html/markdown)")
	w := flag.String("w", "", "wrap output in tag")
	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], sisyphus.Version)
		os.Exit(0)
	}

	var flavor sisyphus.Flavor
	switch *f {
	case "html":
		flavor = &sisyphus.Html{Inline: *x, Aspeq: *a, Wrap: *w}
	case "markdown":
		flavor = &sisyphus.Markdown{Inline: *x}
	}
	sisyphus.Cook(os.Stdin, os.Stdout, flavor)
}
