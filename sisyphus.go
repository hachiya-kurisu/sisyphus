package sisyphus

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"
)

const Version = "0.4.5"

type LinkHook func(uri string, text string, match string) string
type QuoteHook func(text string) string
type Hook func() string

type Flavor interface {
	Open() string
	Close() string
	ListItem(string) string
	Header(int, string) string
	Link(string, string) string
	Pre(string) string
	Quote(string) string
	Text(string) string

	SetState(State) string
	GetState() State

	OnLink(string, LinkHook)
	OnQuote(QuoteHook)
	OnOpen(Hook)
	OnClose(Hook)

	Aspeq(string, bool) LinkHook
	Wrap(string)
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
	fmt.Fprintf(w, "%s", flavor.Open())
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, "```"):
			if flavor.GetState() == Pre {
				fmt.Fprintf(w, "%s", flavor.SetState(None))
			} else {
				fmt.Fprintf(w, "%s", flavor.SetState(Pre))
			}
		case flavor.GetState() == Pre:
			fmt.Fprintln(w, flavor.Pre(line))
		case strings.HasPrefix(line, "* "):
			fmt.Fprintf(w, "%s", flavor.SetState(List))
			raw := strings.TrimSpace(strings.TrimPrefix(line, "*"))
			fmt.Fprintf(w, "%s", flavor.ListItem(raw))
		case strings.HasPrefix(line, ">"):
			fmt.Fprintf(w, "%s", flavor.SetState(Quote))
			raw := strings.TrimSpace(strings.TrimPrefix(line, ">"))
			fmt.Fprintln(w, flavor.Quote(raw))
		case strings.HasPrefix(line, "###"):
			fmt.Fprintf(w, "%s", flavor.SetState(None))
			raw := strings.TrimSpace(strings.TrimPrefix(line, "###"))
			fmt.Fprintf(w, "%s", flavor.Header(3, raw))
		case strings.HasPrefix(line, "##"):
			fmt.Fprintf(w, "%s", flavor.SetState(None))
			raw := strings.TrimSpace(strings.TrimPrefix(line, "##"))
			fmt.Fprintf(w, "%s", flavor.Header(2, raw))
		case strings.HasPrefix(line, "#"):
			fmt.Fprintf(w, "%s", flavor.SetState(None))
			raw := strings.TrimSpace(strings.TrimPrefix(line, "#"))
			fmt.Fprintf(w, "%s", flavor.Header(1, raw))
		case strings.HasPrefix(line, "=>"):
			link := strings.TrimSpace(strings.TrimPrefix(line, "=>"))
			s := regexp.MustCompile("[[:space:]]+").Split(link, 2)
			var ln string
			switch len(s) {
			case 1:
				ln = flavor.Link(s[0], "")
			case 2:
				ln = flavor.Link(s[0], strings.TrimSpace(s[1]))
			}
			if ln != "" {
				fmt.Fprintf(w, "%s", flavor.SetState(Text))
				fmt.Fprintln(w, ln)
			}
		case strings.TrimSpace(line) == "":
			fmt.Fprintf(w, "%s", flavor.SetState(None))
			fmt.Fprintln(w, "")
		default:
			fmt.Fprintf(w, "%s", flavor.SetState(Text))
			fmt.Fprintln(w, flavor.Text(line))
		}
	}
	fmt.Fprintf(w, "%s", flavor.SetState(None))
	fmt.Fprintf(w, "%s", flavor.Close())
}
