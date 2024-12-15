package sisyphus

import (
	"fmt"
	"html"
)

func escape(raw string) string {
	return html.EscapeString(raw)
}

type Html struct{}

func (html Html) Link(url string, text string) string {
	return fmt.Sprintf("<a href='%s'>%s</a>", escape(url), escape(text))
}

func (html Html) Image(url string) string {
	return fmt.Sprintf("<img src='%s' alt>", escape(url))
}

func (html Html) Header(level int, text string) string {
	return fmt.Sprintf("<h%d>%s</h%d>\n", level, escape(text), level)
}

func (html Html) Paragraph(text string) string {
	return fmt.Sprintf("<p>%s", escape(text))
}

func (html Html) Line(text string, nl bool) string {
	if nl {
		return fmt.Sprintf("<br>\n%s\n", escape(text))
	} else {
		return fmt.Sprintf("%s\n", escape(text))
	}
}

func (html Html) Pre(text string) string {
	return escape(text)
}

func (html Html) ListItem(text string) string {
	return fmt.Sprintf("<li>%s\n", escape(text))
}

func (html Html) OpenList() string {
	return "<ul>"
}

func (html Html) CloseList() string {
	return "</ul>"
}

func (html Html) OpenQuote() string {
	return "<blockquote>"
}

func (html Html) CloseQuote() string {
	return "</blockquote>"
}

func (html Html) OpenPre() string {
	return "<pre>"
}

func (html Html) ClosePre() string {
	return "</pre>"
}
