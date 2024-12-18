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
	Wrap    string
	State   State
	Hooks   []*Hook
}

var Tags = map[State][2]string{
	None:  {"", ""},
	Text:  {"<p>", ""},
	List:  {"<ul>\n", "</ul>\n"},
	Pre:   {"<pre>", "</pre>\n"},
	Quote: {"<blockquote>\n<p>", "</blockquote>\n"},
}

func (html *Html) On(state State, suffix, ext string, cb Callback) {
	html.Hooks = append(html.Hooks, &Hook{suffix, ext, cb})
}

func (html *Html) Open() string {
	if html.Wrap != "" {
		return fmt.Sprintf("<%s>\n", html.Wrap)
	}
	return ""
}

func (html *Html) Close() string {
	if html.Wrap != "" {
		return fmt.Sprintf("</%s>", html.Wrap)
	}
	return ""
}

func (html *Html) Header(level int, text string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, Safe(text), level)
}

func (html *Html) Link(url string, text string) string {
	for _, h := range html.Hooks {
		if strings.HasSuffix(text, h.Suffix) && strings.HasSuffix(url, h.Ext) {
			return h.Callback(Safe(url), Safe(text), h.Suffix)
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
	closing := Tags[html.State][1]
	opening := Tags[state][0]
	html.State = state
	return closing + opening
}
