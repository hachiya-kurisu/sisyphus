package sisyphus

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const Version = "0.0.2"

type Flavor interface {
	Link(url string, text string) string
	OpenList() string
	CloseList() string
	OpenQuote() string
	CloseQuote() string
	OpenPre() string
	ClosePre() string
	Pre(text string) string
	ListItem(text string) string
	Line(text string, nl bool) string
	Header(level int, text string) string
	Image(url string) string
	Paragraph(text string) string
}

func Gem(r io.Reader, w io.Writer, flavor Flavor) {
	var text, quote, list, pre bool

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		// close open lists and quotes if necessary
		if list && !strings.HasPrefix(line, "* ") {
			fmt.Fprintln(w, flavor.CloseList())
			list = false
		}
		if quote && !strings.HasPrefix(line, ">") {
			fmt.Fprintln(w, flavor.CloseQuote())
			quote = false
		}

		switch {
		case strings.HasPrefix(line, "```"):
			pre = !pre
			if pre {
				fmt.Fprintln(w, flavor.OpenPre())
			} else {
				fmt.Fprintln(w, flavor.ClosePre())
			}
		case pre:
			fmt.Fprintln(w, flavor.Pre(line))
		case strings.HasPrefix(line, "* "):
			if !list {
				fmt.Fprintf(w, flavor.OpenList())
				list = true
			}
			text := strings.TrimSpace(strings.TrimPrefix(line, "*"))
			fmt.Fprintf(w, flavor.ListItem(text))
		case strings.HasPrefix(line, ">"):
			if !quote {
				fmt.Fprintf(w, flavor.OpenQuote())
				quote = true
			} else {
				fmt.Fprintf(w, flavor.Line("", true))
			}
			quote = true
			text := strings.TrimSpace(strings.TrimPrefix(line, ">"))
			fmt.Fprintln(w, flavor.Line(text, false))
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
			if strings.HasSuffix(link, ".jpg") {
				fmt.Fprintln(w, flavor.Image(link))
			} else {
				parts := strings.SplitN(link, " ", 2)
				url := parts[0]
				if len(parts) == 1 {
					fmt.Fprintln(w, flavor.Link(url, url))
				} else {
					text := strings.TrimSpace(parts[1])
					fmt.Fprintln(w, flavor.Link(url, text))
				}
			}
		case strings.TrimSpace(line) == "":
			text = false
			fmt.Fprintln(w, "")
		default:
			if text {
				fmt.Fprintf(w, flavor.Line(line, true))
			} else {
				fmt.Fprintln(w, flavor.Paragraph(line))
				text = true
			}
		}
	}

	// close any remaining open tags
	if list {
		fmt.Fprintln(w, flavor.CloseList())
	}
	if quote {
		fmt.Fprintln(w, flavor.CloseQuote())
	}
	if pre {
		fmt.Fprintln(w, flavor.ClosePre())
	}
}
