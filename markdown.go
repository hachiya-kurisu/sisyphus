package sisyphus

import (
	"fmt"
	"strings"
)

type Markdown struct{}

func (md Markdown) Link(url string, text string) string {
	return fmt.Sprintf("[%s](%s)", text, url)
}

func (md Markdown) Image(url string) string {
	return fmt.Sprintf("![](%s)", url)
}

func (md Markdown) Header(level int, text string) string {
	return fmt.Sprintf("%s %s\n", strings.Repeat("#", level), text)
}

func (md Markdown) Text(text string, ongoing bool) string {
	return text
}

func (md Markdown) Pre(text string) string {
	return text
}

func (md Markdown) ListItem(text string) string {
	return fmt.Sprintf("* %s\n", text)
}

func (md Markdown) TogglePre(open bool) string {
	return "```"
}

func (md Markdown) ToggleList(open bool) string {
	return ""
}

func (md Markdown) ToggleQuote(open bool) string {
	return ""
}
