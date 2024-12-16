package sisyphus

import (
	"blekksprut.net/aspeq"
	"fmt"
	"html"
	"net/url"
	"strings"
)

type MediaHandler func(*Html, string, string, string) string

type MediaRule struct {
	Suffix  string
	Ext     string
	Handler MediaHandler
}

var Keywords = [...]MediaRule{
	{"(image)", "", (*Html).image},
	{"(photo)", "", (*Html).image},
	{"(photograph)", "", (*Html).image},
	{"(picture)", "", (*Html).image},
	{"(video)", "", (*Html).video},
	{"(audio)", "", (*Html).audio},
	{"(music)", "", (*Html).audio},
	{"", ".jpg", (*Html).image},
	{"", ".png", (*Html).image},
	{"", ".gif", (*Html).image},
	{"", ".mp4", (*Html).video},
	{"", ".m4a", (*Html).audio},
	{"", ".mp3", (*Html).audio},
	{"", ".ogg", (*Html).audio},
}

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

func (html *Html) video(url string, text string, suffix string) string {
	text = strings.TrimSpace(strings.TrimSuffix(text, suffix))
	return fmt.Sprintf("<video controls src='%s' title='%s'></video>", url, text)
}

func (html *Html) audio(url string, text string, suffix string) string {
	text = strings.TrimSpace(strings.TrimSuffix(text, suffix))
	return fmt.Sprintf("<audio controls src='%s' title='%s'></audio>", url, text)
}

func (html *Html) image(uri string, text string, suffix string) string {
	text = strings.TrimSpace(strings.TrimSuffix(text, suffix))
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
	if html.Extended {
		for _, kw := range Keywords {
			switch {
			case kw.Suffix != "" && strings.HasSuffix(text, kw.Suffix):
				return kw.Handler(html, escape(url), escape(text), kw.Suffix)
			case kw.Ext != "" && strings.HasSuffix(url, kw.Ext):
				return kw.Handler(html, escape(url), escape(text), "")
			}
		}
	}
	if text == "" {
		text = url
	}
	if html.Current == url {
		return fmt.Sprintf("<a class=x href='%s'>%s</a>", escape(url), escape(text))
	} else {
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
