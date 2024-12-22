package sisyphus

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Markdown struct {
	State State

	LinkHooks map[string]LinkHook
	QuoteHook QuoteHook
	OpenHook  Hook
	CloseHook Hook
}

func (md *Markdown) OnLink(rule string, hook LinkHook) {
	if md.LinkHooks == nil {
		md.LinkHooks = make(map[string]LinkHook)
	}
	md.LinkHooks[rule] = hook
}

func (md *Markdown) OnQuote(hook QuoteHook) {
	md.QuoteHook = hook
}

func (md *Markdown) OnOpen(hook Hook) {
	md.OpenHook = hook
}

func (md *Markdown) OnClose(hook Hook) {
	md.CloseHook = hook
}

func (md *Markdown) Wrap(s string) {
	if s == "" {
		return
	}
	md.OnOpen(func() string {
		return s
	})
	md.OnClose(func() string {
		return s
	})
}

func (md *Markdown) Open() string {
	if md.OpenHook != nil {
		return md.OpenHook()
	}
	return ""
}

func (md *Markdown) Close() string {
	if md.CloseHook != nil {
		return md.CloseHook()
	}
	return ""
}

func (md *Markdown) Header(level int, text string) string {
	return fmt.Sprintf("%s %s\n", strings.Repeat("#", level), text)
}

func (md *Markdown) Link(url string, text string) string {
	ext := filepath.Ext(url)
	hook, ok := md.LinkHooks[ext]
	if ok {
		return hook(Safe(url), Safe(text), ext)
	}
	if text == "" {
		text = url
	}
	return fmt.Sprintf("[%s](%s)", text, url)
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
	if md.QuoteHook != nil {
		return md.QuoteHook(text)
	}
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
