package sisyphus

import (
	"fmt"
	"html"
	"strings"
)

func escape(raw string) string {
	return html.EscapeString(raw)
}

type Html struct{
	Images bool
}

func (html Html) Header(level int, text string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, escape(text), level)
}

func (html Html) Image(url string) string {
	return fmt.Sprintf("<img src='%s' alt>", escape(url))
}

func (html Html) Link(url string, text string) string {
	if html.Images && strings.HasSuffix(url, ".jpg") {
		return fmt.Sprintf("<img src='%s' alt='%s'>", escape(url), escape(text))
	} else {
		return fmt.Sprintf("<a href='%s'>%s</a>", escape(url), escape(text))
	}
}

func (html Html) ListItem(text string) string {
	return fmt.Sprintf("<li>%s\n", escape(text))
}

func (html Html) Pre(text string) string {
	return escape(text)
}

func (html Html) Text(text string, ongoing bool) string {
	if ongoing {
		return fmt.Sprintf("<br>\n%s\n", escape(text))
	} else {
		return fmt.Sprintf("<p>%s", escape(text))
	}
}

func (html Html) ToggleList(open bool) string {
	if open {
		return "<ul>"
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
		return "<blockquote>"
	} else {
		return "</blockquote>"
	}
}
