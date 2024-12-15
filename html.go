package sisyphus

import (
	"fmt"
	"html"
	"strings"
)

func escape(raw string) string {
	return html.EscapeString(raw)
}

type Html struct {
	Extended bool
}

func (html Html) Header(level int, text string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, escape(text), level)
}

func (html Html) video(url string, text string) string {
	return fmt.Sprintf("<video controls src='%s' title='%s'></video>", url, text)
}

func (html Html) Link(url string, text string, ongoing bool) string {
	switch {
	case html.Extended && strings.HasSuffix(text, "(image)"):
		text = strings.TrimSpace(strings.TrimSuffix(text, "(image)"))
		return fmt.Sprintf("<p><img src='%s' alt='%s'>", escape(url), escape(text))
	case html.Extended && strings.HasSuffix(text, "(video)"):
		text = strings.TrimSpace(strings.TrimSuffix(text, "(video)"))
		return html.video(escape(url), escape(text))
	default:
		return fmt.Sprintf("<p><a href='%s'>%s</a>", escape(url), escape(text))
	}
}

func (html Html) ListItem(text string) string {
	return fmt.Sprintf("<li>%s\n", escape(text))
}

func (html Html) Pre(text string) string {
	return escape(text)
}

func (html Html) Quote(text string, ongoing bool) string {
	return html.Text(text, ongoing)
}

func (html Html) Text(text string, ongoing bool) string {
	if ongoing {
		return fmt.Sprintf("<br>\n%s", escape(text))
	} else {
		return fmt.Sprintf("<p>%s", escape(text))
	}
}

func (html Html) ToggleList(open bool) string {
	if open {
		return "<ul>\n"
	} else {
		return "</ul>"
	}
}

func (html Html) TogglePre(open bool) string {
	if open {
		return "<pre>"
	} else {
		return "</pre>"
	}
}

func (html Html) ToggleQuote(open bool) string {
	if open {
		return "<blockquote>\n"
	} else {
		return "</blockquote>"
	}
}
