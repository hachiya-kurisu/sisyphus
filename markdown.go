package sisyphus

import (
	"fmt"
	"strings"
)

type Markdown struct{}

func (md Markdown) Link(url string, text string) string {
	return fmt.Sprintf("[%s](%s)\n", text, url)
}

func (md Markdown) Image(url string) string {
	return fmt.Sprintf("![](%s)\n", url)
}

func (md Markdown) Header(level int, text string) string {
	return fmt.Sprintf("%s %s\n", strings.Repeat("#", level), text)
}

func (md Markdown) Paragraph(text string) string {
	return text + "\n"
}

func (md Markdown) Line(text string, nl bool) string {
	return text + "\n"
}

func (md Markdown) Pre(text string) string {
	return text + "\n"
}

func (md Markdown) ListItem(text string) string {
	return fmt.Sprintf("* %s\n", text)
}

func (md Markdown) OpenList() string {
	return ""
}

func (md Markdown) CloseList() string {
	return ""
}

func (md Markdown) OpenQuote() string {
	return ""
}

func (md Markdown) CloseQuote() string {
	return ""
}

func (md Markdown) OpenPre() string {
	return "```\n"
}

func (md Markdown) ClosePre() string {
	return "```\n"
}
