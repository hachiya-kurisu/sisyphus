package main

import (
	"blekksprut.net/sisyphus"
	"flag"
	"fmt"
	"os"
)

func main() {
	v := flag.Bool("v", false, "version")
	a := flag.String("a", "", "aspeq prefix")
	f := flag.String("f", "html", "flavor (html/markdown)")
	_ = flag.String("w", "", "wrap output in tag")
	flag.Parse()

	if *v {
		fmt.Printf("%s %s\n", os.Args[0], sisyphus.Version)
		os.Exit(0)
	}

	var flavor sisyphus.Flavor
	switch *f {
	case "html":
		flavor = &sisyphus.Html{} // Wrap: *w}
		if *a != "" {
			flavor.OnLink(".jpg", sisyphus.Aspeq(*a, false))
		}
	case "markdown":
		flavor = &sisyphus.Markdown{}
	}

	sisyphus.Cook(os.Stdin, os.Stdout, flavor)
}
