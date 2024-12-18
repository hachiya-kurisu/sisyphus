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

var Keywords = []MediaRule{
	{"(image)", "", (*Html).image},
	{"(photo)", "", (*Html).image},
	{"(photograph)", "", (*Html).image},
	{"(illustration)", "", (*Html).image},
	{"(picture)", "", (*Html).image},
	{"(foto)", "", (*Html).image},
	{"(bilde)", "", (*Html).image},
	{"(写真)", "", (*Html).image},
	{"(映像)", "", (*Html).image},
	{"(video)", "", (*Html).video},
	{"(動画)", "", (*Html).video},
	{"(audio)", "", (*Html).audio},
	{"(music)", "", (*Html).audio},
	{"(lyd)", "", (*Html).audio},
	{"(音)", "", (*Html).audio},
	{"(音声)", "", (*Html).audio},
	{"(音楽)", "", (*Html).audio},
	{"", ".jpg", (*Html).image},
	{"", ".png", (*Html).image},
	{"", ".gif", (*Html).image},
	{"", ".mp4", (*Html).video},
	{"", ".m4a", (*Html).audio},
	{"", ".mp3", (*Html).audio},
	{"", ".ogg", (*Html).audio},
}

func Safe(raw string) string {
	return html.EscapeString(raw)
}

type OnImage func(string, string, string) string

type Html struct {
	Inline  bool
	Aspeq   string
	Current string
	Wrap    string
	OnImage OnImage
	State   State
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

func (html *Html) video(url string, text string, suffix string) string {
	text = strings.TrimSpace(strings.TrimSuffix(text, suffix))
	return fmt.Sprintf("<video controls src='%s' title='%s'></video>", url, text)
}

func (html *Html) audio(url string, text string, suffix string) string {
	text = strings.TrimSpace(strings.TrimSuffix(text, suffix))
	return fmt.Sprintf("<audio controls src='%s' title='%s'></audio>", url, text)
}

func (html *Html) image(uri string, text string, suffix string) string {
	if html.OnImage != nil {
		return html.OnImage(uri, text, suffix)
	}
	text = strings.TrimSpace(strings.TrimSuffix(text, suffix))
	parsed, err := url.Parse(uri)
	if err == nil && html.Aspeq != "" && !parsed.IsAbs() {
		path := fmt.Sprintf("%s/%s", html.Aspeq, uri)
		ar, err := aspeq.FromImage(path)
		if err == nil {
			return fmt.Sprintf("<img src='%s' class=%s alt='%s'>", uri, ar.Name, text)
		}
	}
	return fmt.Sprintf("<img src='%s' alt='%s'>", uri, text)
}

func (html *Html) Link(url string, text string) string {
	if html.Inline {
		for _, kw := range Keywords {
			switch {
			case kw.Suffix != "" && strings.HasSuffix(text, kw.Suffix):
				return kw.Handler(html, Safe(url), Safe(text), kw.Suffix)
			case kw.Ext != "" && strings.HasSuffix(url, kw.Ext):
				return kw.Handler(html, Safe(url), Safe(text), "")
			}
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
