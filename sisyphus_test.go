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
	{"> hello\n> hm", "> hello\n> hm\n"},
	{"> hello\nhm", "> hello\nhm\n"},
	{"```\npre", "```\npre\n```\n"},
}

func TestHtml(t *testing.T) {
	for _, c := range htmlcases {
		html := Convert(c[0], &Html{})
		if html != c[1] {
			t.Errorf("%s should be %s", html, c[1])
		}
	}
}

func TestMarkdown(t *testing.T) {
	for _, c := range mdcases {
		md := Convert(c[0], &Markdown{})
		if md != c[1] {
			t.Errorf("%s should be %s", md, c[1])
		}
	}
}

func TestCallback(t *testing.T) {
	flavor := &Markdown{}
	flavor.On(Link, "", ".jpg", func(url, what, lol string) string {
		return "hijacked!"
	})
	gmi := "=> test.jpg"
	expect := "hijacked!\n"
	html := Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestAspeq(t *testing.T) {
	flavor := &Html{}
	flavor.On(Link, "", ".jpg", Aspeq("."))
	gmi := "=> ume.jpg 梅ちゃん"
	expect := "<p><img src='ume.jpg' class=super16 alt='梅ちゃん'>\n"
	html := Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestAspeqMissing(t *testing.T) {
	flavor := &Html{}
	flavor.On(Link, "", ".jpg", Aspeq("."))
	gmi := "=> notfound.jpg"
	expect := "<p><img src='notfound.jpg' class=unknown alt=''>\n"
	html := Convert(gmi, flavor)
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

func TestWrap(t *testing.T) {
	gmi := "wrap me up"
	html := Convert(gmi, &Html{Wrap: "article"})
	expect := "<article>\n<p>wrap me up\n</article>"
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func callback(uri, text, suffix string) string {
	return "-.-"
}

