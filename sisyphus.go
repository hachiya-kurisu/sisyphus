package sisyphus

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const Version = "0.0.6"

type Flavor interface {
	Header(level int, text string) string
	Link(url string, text string) string
	ListItem(text string) string
	Pre(text string) string
	Quote(text string) string
	Text(text string) string
	SetState(state State) string
	GetState() State
}

// states
type State int

const (
	None State = iota + 1
	Text
	List
	Pre
	Quote
)

func Gem(r io.Reader, w io.Writer, flavor Flavor) {
	flavor.SetState(None)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "```"):
			if flavor.GetState() == Pre {
				fmt.Fprintf(w, flavor.SetState(None))
			} else {
				fmt.Fprintf(w, flavor.SetState(Pre))
			}
		case flavor.GetState() == Pre:
			fmt.Fprintln(w, flavor.Pre(line))
		case strings.HasPrefix(line, "* "):
			fmt.Fprintf(w, flavor.SetState(List))
			raw := strings.TrimSpace(strings.TrimPrefix(line, "*"))
			fmt.Fprintf(w, flavor.ListItem(raw))
		case strings.HasPrefix(line, ">"):
			fmt.Fprintf(w, flavor.SetState(Quote))
			raw := strings.TrimSpace(strings.TrimPrefix(line, ">"))
			fmt.Fprintln(w, flavor.Quote(raw))
		case strings.HasPrefix(line, "###"):
			raw := strings.TrimSpace(strings.TrimPrefix(line, "###"))
			fmt.Fprintf(w, flavor.Header(3, raw))
		case strings.HasPrefix(line, "##"):
			raw := strings.TrimSpace(strings.TrimPrefix(line, "##"))
			fmt.Fprintf(w, flavor.Header(2, raw))
		case strings.HasPrefix(line, "#"):
			raw := strings.TrimSpace(strings.TrimPrefix(line, "#"))
			fmt.Fprintf(w, flavor.Header(1, raw))
		case strings.HasPrefix(line, "=>"):
			link := strings.TrimSpace(strings.TrimPrefix(line, "=>"))
			url, text, _ := strings.Cut(link, " ")
			fmt.Fprintf(w, flavor.SetState(Text))
			fmt.Fprintln(w, flavor.Link(url, strings.TrimSpace(text)))
		case strings.TrimSpace(line) == "":
			fmt.Fprintf(w, flavor.SetState(None))
			fmt.Fprintln(w, "")
		default:
			fmt.Fprintf(w, flavor.SetState(Text))
			fmt.Fprintln(w, flavor.Text(line))
		}
	}
	fmt.Fprintf(w, flavor.SetState(None))
}
