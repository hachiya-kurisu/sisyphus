package sisyphus

import (
	"blekksprut.net/aspeq"
	"fmt"
	"html"
	"net/url"
	"strings"
)

func escape(raw string) string {
	return html.EscapeString(raw)
}

type Html struct {
	Extended bool
	Aspeq    bool
	Current  string
	State    State
}

func (html *Html) Header(level int, text string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, escape(text), level)
}

func (html *Html) video(url string, text string) string {
	return fmt.Sprintf("<video controls src='%s' title='%s'></video>", url, text)
}

func (html *Html) image(uri string, text string) string {
	parsed, err := url.Parse(uri)
	if err == nil && html.Aspeq && !parsed.IsAbs() {
		ar, err := aspeq.FromImage(uri)
		if err == nil {
			return fmt.Sprintf("<img src='%s' class=%s alt='%s'>", uri, ar.Name, text)
		}
	}
	return fmt.Sprintf("<img src='%s' alt='%s'>", uri, text)
}

func (html *Html) Link(url string, text string) string {
	if text == "" {
		text = url
	}
	switch {
	case html.Extended && strings.HasSuffix(text, "(image)"):
		text = strings.TrimSpace(strings.TrimSuffix(text, "(image)"))
		return html.image(escape(url), escape(text))
	case html.Extended && strings.HasSuffix(text, "(video)"):
		text = strings.TrimSpace(strings.TrimSuffix(text, "(video)"))
		return html.video(escape(url), escape(text))
	default:
		return fmt.Sprintf("<a href='%s'>%s</a>", escape(url), escape(text))
	}
}

func (html *Html) ListItem(text string) string {
	return fmt.Sprintf("<li>%s\n", escape(text))
}

func (html *Html) Pre(text string) string {
	return escape(text)
}

func (html *Html) Quote(text string) string {
	return html.Text(text)
}

func (html *Html) Text(text string) string {
	return escape(text)
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
