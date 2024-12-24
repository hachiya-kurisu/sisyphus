package main

import (
	"blekksprut.net/sisyphus"
	"flag"
	"fmt"
	"os"
)

func main() {
	v := flag.Bool("v", false, "version")
	f := flag.String("f", "html", "flavor (html/markdown)")
	w := flag.String("w", "", "wrap output")
	s := flag.String("s", "", "url of current page")
	a := flag.String("a", "", "aspeq prefix (html only)")
	g := flag.Bool("g", false, "greentext mode (html only)")
	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], sisyphus.Version)
		os.Exit(0)
	}

	var flavor sisyphus.Flavor
	switch *f {
	case "html":
		flavor = &sisyphus.Html{Self: *s, Greentext: *g}
		flavor.Wrap(*w)
		if *a != "" {
			flavor.OnLink(".jpg", sisyphus.Aspeq(*a, false))
		}
	case "markdown":
		flavor = &sisyphus.Markdown{}
		flavor.Wrap(*w)
	}

	sisyphus.Cook(os.Stdin, os.Stdout, flavor)
}
