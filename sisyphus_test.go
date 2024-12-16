package sisyphus

import "testing"

var htmlcases = [][]string{
	{"text\nanother line\n\n", "<p>text\n<br>another line\n\n"},
	{"# header", "<h1>header</h1>\n"},
	{"## header", "<h2>header</h2>\n"},
	{"### header", "<h3>header</h3>\n"},
	{"* list", "<ul>\n<li>list\n</ul>\n"},
	{"* 1\n> 2", "<ul>\n<li>1\n</ul>\n<blockquote>\n<p>2\n</blockquote>\n"},
	{"=> link", "<p><a href='link'>link</a>\n"},
	{"=> src oh hay (image)", "<p><img src='src' alt='oh hay'>\n"},
	{"=> picto fff (photograph)", "<p><img src='picto' alt='fff'>\n"},
	{"=> ume.jpg", "<p><img src='ume.jpg' alt=''>\n"},
	{"=> ume.mp4", "<p><video controls src='ume.mp4' title=''></video>\n"},
	{"=> ume.m4a", "<p><audio controls src='ume.m4a' title=''></audio>\n"},
	{"=> vid (video)", "<p><video controls src='vid' title=''></video>\n"},
	{"=> meh uh (audio)", "<p><audio controls src='meh' title='uh'></audio>\n"},
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
		html := Convert(c[0], &Html{Extended: true})
		if html != c[1] {
			t.Errorf("%s should be %s", html, c[1])
		}
	}
}

func TestMarkdown(t *testing.T) {
	for _, c := range mdcases {
		md := Convert(c[0], &Markdown{Extended: true})
		if md != c[1] {
			t.Errorf("%s should be %s", md, c[1])
		}
	}
}

func TestAspeq(t *testing.T) {
	gmi := "=> ume.jpg 梅ちゃん (image)"
	expect := "<p><img src='ume.jpg' class=super16 alt='梅ちゃん'>\n"
	html := Convert(gmi, &Html{Extended: true, Aspeq: true})
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestCurrent(t *testing.T) {
	gmi := "=> /here a link"
	html := Convert(gmi, &Html{Current: "/here"})
	expect := "<p><a class=x href='/here'>a link</a>\n"
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}
