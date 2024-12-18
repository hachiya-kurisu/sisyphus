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
	var closing string
	var opening string
	switch html.State {
	case List:
		closing = "</ul>\n"
	case Quote:
		closing = "</blockquote>\n"
	case Pre:
		closing = "</pre>\n"
	}
	switch state {
	case List:
		opening = "<ul>\n"
	case Quote:
		opening = "<blockquote>\n<p>"
	case Pre:
		opening = "<pre>\n"
	case Text:
		opening = "<p>"
	}
	html.State = state
	return closing + opening
}
