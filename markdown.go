package sisyphus

import (
	"fmt"
	"strings"
)

type Markdown struct {
	Inline  bool
	Current string
	State   State
}

func (md *Markdown) Open() string {
	return ""
}

func (md *Markdown) Close() string {
	return ""
}

func (md *Markdown) Header(level int, text string) string {
	return fmt.Sprintf("%s %s\n", strings.Repeat("#", level), text)
}

func (md *Markdown) Link(url string, text string) string {
	if text == "" {
		text = url
	}
	switch {
	case md.Inline && strings.HasSuffix(text, "(image)"):
		text = strings.TrimSpace(strings.TrimSuffix(text, "(image)"))
		return fmt.Sprintf("![%s](%s)", text, url)
	default:
		return fmt.Sprintf("[%s](%s)", text, url)
	}
}

func (md *Markdown) ListItem(text string) string {
	return fmt.Sprintf("* %s\n", text)
}

func (md *Markdown) Pre(text string) string {
	return text
}

func (md *Markdown) Text(text string) string {
	return text
}

func (md *Markdown) Quote(text string) string {
	return "> " + text
}

func (md *Markdown) GetState() State {
	return md.State
}

func (md *Markdown) SetState(state State) string {
	if state == md.State {
		return ""
	}
	var closing, opening string
	switch md.State {
	case Pre:
		closing = "```\n"
	}
	switch state {
	case Pre:
		opening = "```\n"
	}
	md.State = state
	return closing + opening
}
