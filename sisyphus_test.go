package sisyphus

import (
	"strings"
	"testing"
)

var htmlcases = [][]string{
	{"text\nanother line\n\n", "<p>text\n<br>another line\n\n"},
	{"# header", "<h1>header</h1>\n"},
	{"## header", "<h2>header</h2>\n"},
	{"### header", "<h3>header</h3>\n"},
	{"* list", "<ul>\n<li>list\n</ul>\n"},
	{"* 1\n> 2", "<ul>\n<li>1\n</ul>\n<blockquote>\n<p>2\n</blockquote>\n"},
	{"=> link", "<p><a href='link'>link</a>\n"},
	{"=> src oh hay (image)", "<p><img src='src' alt='oh hay'>\n"},
	{"=> vid (video)", "<p><video controls src='vid' title=''></video>\n"},
	{"> hello\n> hm", "<blockquote>\n<p>hello\n<br>hm\n</blockquote>\n"},
	{"> hello\nhm", "<blockquote>\n<p>hello\n</blockquote>\n<p>hm\n"},
	{"```\npre", "<pre>\npre\n</pre>\n"},
	{"```\npre\n```", "<pre>\npre\n</pre>\n"},
}

var mdcases = [][]string{
	{"text\nanother line\n", "text\nanother line\n"},
	{"# header", "# header\n"},
	{"## header", "## header\n"},
	{"### header", "### header\n"},
	{"* list", "* list\n"},
	{"* 1\n> 2", "* 1\n> 2\n"},
	{"=> link", "[link](link)\n"},
	{"=> src oh hay (image)", "![oh hay](src)\n"},
	{"> hello\n> hm", "> hello\n> hm\n"},
	{"> hello\nhm", "> hello\nhm\n"},
	{"```\npre", "```\npre\n```\n"},
}

func TestHtml(t *testing.T) {
	for _, c := range htmlcases {
		var out strings.Builder
		Gem(
			strings.NewReader(c[0]),
			&out,
			&Html{Extended: true},
		)
		if out.String() != c[1] {
			t.Errorf("[%s] should be [%s]", out.String(), c[1])
		}
	}
}

func TestMarkdown(t *testing.T) {
	for _, c := range mdcases {
		var out strings.Builder
		Gem(
			strings.NewReader(c[0]),
			&out,
			&Markdown{Extended: true},
		)
		if out.String() != c[1] {
			t.Errorf("[%s] should be [%s]", out.String(), c[1])
		}
	}
}

func TestAspeq(t *testing.T) {
	var out strings.Builder
	Gem(
		strings.NewReader("=> ume.jpg 梅ちゃん (image)"),
		&out,
		&Html{Extended: true, Aspeq: true},
	)
	if out.String() != "<p><img src='ume.jpg' class=super16 alt='梅ちゃん'>\n" {
		t.Errorf("[%s] should be [%s]", out.String(), "lol")
	}
}
