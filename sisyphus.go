package sisyphus

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const Version = "0.1.1"

type Flavor interface {
	Open() string
	Close() string
	Header(level int, text string) string
	Link(url string, text string) string
	ListItem(text string) string
	Pre(text string) string
	Quote(text string) string
	Text(text string) string
	SetState(state State) string
	GetState() State
}

type State int

const (
	None State = iota + 1
	Text
	List
	Pre
	Quote
)

func Convert(gmi string, flavor Flavor) string {
	rd := strings.NewReader(gmi)
	var b bytes.Buffer
	Cook(rd, &b, flavor)
	return b.String()
}

func Cook(r io.Reader, w io.Writer, flavor Flavor) {
	fmt.Fprintf(w, flavor.Open())
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
	fmt.Fprintf(w, flavor.Close())
}
