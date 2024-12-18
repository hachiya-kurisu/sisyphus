package sisyphus

import (
	"fmt"
	"html"
	"strings"
)

func Safe(raw string) string {
	return html.EscapeString(raw)
}

type Html struct {
	Current string
	State   State
	Hooks   []*Hook
}

var tags = map[State][2]string{
	None:  {"", ""},
	Text:  {"<p>", ""},
	List:  {"<ul>\n", "</ul>\n"},
	Pre:   {"<pre>", "</pre>\n"},
	Quote: {"<blockquote>\n<p>", "</blockquote>\n"},
}

func (html *Html) On(state State, rule string, cb Callback) {
	html.Hooks = append(html.Hooks, &Hook{rule, cb})
}

func (html *Html) Open() string {
	return ""
}

func (html *Html) Close() string {
	return ""
}

func (html *Html) Header(level int, text string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, Safe(text), level)
}

func (html *Html) Link(url string, text string) string {
	for _, h := range html.Hooks {
		if strings.HasSuffix(text, h.Rule) || strings.HasSuffix(url, h.Rule) {
			return h.Callback(Safe(url), Safe(text), h.Rule)
		}
	}
	if text == "" {
		text = url
	}
	if html.Current == url {
		return fmt.Sprintf("<a class=x href='%s'>%s</a>", Safe(url), Safe(text))
	} else {
		return fmt.Sprintf("<a href='%s'>%s</a>", Safe(url), Safe(text))
	}
}

func (html *Html) ListItem(text string) string {
	return fmt.Sprintf("<li>%s\n", Safe(text))
}

func (html *Html) Pre(text string) string {
	return Safe(text)
}

func (html *Html) Quote(text string) string {
	return html.Text(text)
}

func (html *Html) Text(text string) string {
	return Safe(text)
}

func (html *Html) GetState() State {
	return html.State
}

func (html *Html) SetState(state State) string {
	if state == html.State {
		if state == Text || state == Quote {
			return "<br>"
		}
		return ""
	}
	closing := tags[html.State][1]
	opening := tags[state][0]
	html.State = state
	return closing + opening
}
