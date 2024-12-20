package sisyphus

import (
	"fmt"
	"html"
	"path/filepath"
	"regexp"
)

func Safe(raw string) string {
	return html.EscapeString(raw)
}

type Html struct {
	Current   string
	State     State
	LinkHooks map[string]LinkHook
	QuoteHook QuoteHook
	OpenHook  Hook
	CloseHook Hook
}

var tags = map[State][2]string{
	None:  {"", ""},
	Text:  {"<p>", ""},
	List:  {"<ul>\n", "</ul>\n"},
	Pre:   {"<pre>", "</pre>\n"},
	Quote: {"<blockquote>\n<p>", "</blockquote>\n"},
}

func (html *Html) OnLink(rule string, hook LinkHook) {
	if html.LinkHooks == nil {
		html.LinkHooks = make(map[string]LinkHook)
	}
	html.LinkHooks[rule] = hook
}

func (html *Html) OnQuote(hook QuoteHook) {
	html.QuoteHook = hook
}

func (html *Html) OnOpen(hook Hook) {
	html.OpenHook = hook
}

func (html *Html) OnClose(hook Hook) {
	html.CloseHook = hook
}

func (html *Html) Wrap(tag string) {
	html.OnOpen(func() string {
		return fmt.Sprintf("<%s>", tag)
	})
	html.OnClose(func() string {
		return fmt.Sprintf("</%s>", tag)
	})
}

func (html *Html) Open() string {
	if html.OpenHook != nil {
		return html.OpenHook()
	}
	return ""
}

func (html *Html) Close() string {
	if html.CloseHook != nil {
		return html.CloseHook()
	}
	return ""
}

func (html *Html) Header(level int, text string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, Safe(text), level)
}

func (html *Html) Link(url string, text string) string {
	// todo - sjekk suffix for (image)
	ext := filepath.Ext(url)
	hook, ok := html.LinkHooks[ext]
	if ok {
		return hook(Safe(url), Safe(text), ext)
	}

	re := regexp.MustCompile(`\((.+)\)$`)
	tag := re.FindStringSubmatch(text)
	if len(tag) > 1 {
		hook, ok := html.LinkHooks[tag[1]]
		if ok {
			return hook(Safe(url), Safe(text), ext)
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
	if html.QuoteHook != nil {
		return html.QuoteHook(Safe(text))
	}
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
