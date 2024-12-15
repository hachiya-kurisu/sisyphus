package sisyphus

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const Version = "0.0.3"

type Flavor interface {
	Header(level int, text string) string
	Image(url string) string
	Link(url string, text string) string
	ListItem(text string) string
	Pre(text string) string
	Text(text string, nl bool) string
	ToggleList(open bool) string
	TogglePre(open bool) string
	ToggleQuote(open bool) string
}

func Gem(r io.Reader, w io.Writer, flavor Flavor) {
	var text, quote, list, pre bool

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		// close open lists and quotes if necessary
		if list && !strings.HasPrefix(line, "* ") {
			fmt.Fprintln(w, flavor.ToggleList(false))
			list = false
		}
		if quote && !strings.HasPrefix(line, ">") {
			fmt.Fprintln(w, flavor.ToggleQuote(false))
			quote = false
		}

		switch {
		case strings.HasPrefix(line, "```"):
			pre = !pre
			fmt.Fprintln(w, flavor.TogglePre(pre))
		case pre:
			fmt.Fprintln(w, flavor.Pre(line))
		case strings.HasPrefix(line, "* "):
			if !list {
				fmt.Fprintf(w, flavor.ToggleList(true))
				list = true
			}
			text := strings.TrimSpace(strings.TrimPrefix(line, "*"))
			fmt.Fprintf(w, flavor.ListItem(text))
		case strings.HasPrefix(line, ">"):
			if !quote {
				fmt.Fprintf(w, flavor.ToggleQuote(true))
			}
			text := strings.TrimSpace(strings.TrimPrefix(line, ">"))
			fmt.Fprintln(w, flavor.Text(text, quote))
			quote = true
		case strings.HasPrefix(line, "###"):
			text := strings.TrimSpace(strings.TrimPrefix(line, "###"))
			fmt.Fprintf(w, flavor.Header(3, text))
		case strings.HasPrefix(line, "##"):
			text := strings.TrimSpace(strings.TrimPrefix(line, "##"))
			fmt.Fprintf(w, flavor.Header(2, text))
		case strings.HasPrefix(line, "#"):
			text := strings.TrimSpace(strings.TrimPrefix(line, "#"))
			fmt.Fprintf(w, flavor.Header(1, text))
		case strings.HasPrefix(line, "=>"):
			link := strings.TrimSpace(strings.TrimPrefix(line, "=>"))
			parts := strings.SplitN(link, " ", 2)
			if len(parts) == 1 {
				fmt.Fprintln(w, flavor.Link(parts[0], parts[0]))
			} else {
				fmt.Fprintln(w, flavor.Link(parts[0], strings.TrimSpace(parts[1])))
			}
		case strings.TrimSpace(line) == "":
			text = false
			fmt.Fprintln(w, "")
		default:
			fmt.Fprintln(w, flavor.Text(line, text))
			text = true
		}
	}

	// close any remaining open tags
	if list {
		fmt.Fprintln(w, flavor.ToggleList(false))
	} else if quote {
		fmt.Fprintln(w, flavor.ToggleQuote(false))
	} else if pre {
		fmt.Fprintln(w, flavor.TogglePre(false))
	}
}
