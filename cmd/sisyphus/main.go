package main

import (
	"blekksprut.net/sisyphus"
	"flag"
	"fmt"
	"os"
)

func main() {
	v := flag.Bool("v", false, "version")
	a := flag.String("a", "", "aspeq prefix (empty string disables aspeq)")
	f := flag.String("f", "html", "flavor (html/markdown)")
	w := flag.String("w", "", "wrap output in tag")
	s := flag.String("s", "", "url of current page")
	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], sisyphus.Version)
		os.Exit(0)
	}

	var flavor sisyphus.Flavor
	switch *f {
	case "html":
		flavor = &sisyphus.Html{Self: *s}
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
