package sisyphus_test

import (
	"blekksprut.net/sisyphus"
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
	{"> hello\n> hm", "<blockquote>\n<p>hello\n<br>hm\n</blockquote>\n"},
	{"> hello\nhm", "<blockquote>\n<p>hello\n</blockquote>\n<p>hm\n"},
	{"```\npre", "<pre>pre\n</pre>\n"},
	{"```\npre\n```", "<pre>pre\n</pre>\n"},
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
		html := sisyphus.Convert(c[0], &sisyphus.Html{})
		if html != c[1] {
			t.Errorf("%s should be %s", html, c[1])
		}
	}
}

func TestMarkdown(t *testing.T) {
	for _, c := range mdcases {
		md := sisyphus.Convert(c[0], &sisyphus.Markdown{})
		if md != c[1] {
			t.Errorf("%s should be %s", md, c[1])
		}
	}
}

func TestCallback(t *testing.T) {
	flavor := &sisyphus.Markdown{}
	flavor.OnLink(".jpg", func(_, _, _ string) string {
		return "hijacked!"
	})
	gmi := "=> test.jpg"
	expect := "hijacked!\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestAspeq(t *testing.T) {
	flavor := &sisyphus.Html{}
	flavor.OnLink(".jpg", flavor.Aspeq(".", false))
	gmi := "=> ume.jpg 梅ちゃん"
	expect := "<p><img src='ume.jpg' class=super16 alt='梅ちゃん'>\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestAspeqMissing(t *testing.T) {
	flavor := &sisyphus.Html{}
	flavor.OnLink(".jpg", flavor.Aspeq(".", true))
	gmi := "=> notfound.jpg"
	expect := "<p><img src='notfound.jpg' class=unknown alt=''>\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestMarkdownAspeq(t *testing.T) {
	flavor := &sisyphus.Markdown{}
	flavor.OnLink(".jpg", flavor.Aspeq(".", true))
	gmi := "=> ume.jpg 梅ちゃん"
	expect := "![ume.jpg](梅ちゃん)\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestSelf(t *testing.T) {
	gmi := "=> /here a link"
	html := sisyphus.Convert(gmi, &sisyphus.Html{Self: "/here"})
	expect := "<p><a data-friendly class=self href='/here'>a link</a>\n"
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestHtmlWrap(t *testing.T) {
	gmi := "wrap me up"
	flavor := sisyphus.Html{}
	flavor.Wrap("article")
	html := sisyphus.Convert(gmi, &flavor)
	expect := "<article>\n<p>wrap me up\n</article>\n"
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestHtmlEmptyWrap(t *testing.T) {
	gmi := "do not wrap me up"
	flavor := sisyphus.Html{}
	flavor.Wrap("")
	html := sisyphus.Convert(gmi, &flavor)
	expect := "<p>do not wrap me up\n"
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestMarkdownWrap(t *testing.T) {
	gmi := "wrap me up"
	flavor := sisyphus.Markdown{}
	flavor.Wrap("```")
	html := sisyphus.Convert(gmi, &flavor)
	expect := "```wrap me up\n```"
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestMarkdownEmptyWrap(t *testing.T) {
	gmi := "do not wrap me up"
	flavor := sisyphus.Markdown{}
	flavor.Wrap("")
	html := sisyphus.Convert(gmi, &flavor)
	expect := "do not wrap me up\n"
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestHtmlLinkTag(t *testing.T) {
	flavor := &sisyphus.Html{}
	flavor.OnLink("photo", func(_, _, _ string) string {
		return "!"
	})
	gmi := "=> test.jpg (photo)"
	expect := "<p>!\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestHtmlQuote(t *testing.T) {
	flavor := &sisyphus.Html{}
	flavor.OnQuote(func(_ string) string {
		return "!"
	})
	gmi := ">>12435"
	expect := "<blockquote>\n<p>!\n</blockquote>\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestMdWrap(t *testing.T) {
	gmi := "wrap me up"
	flavor := sisyphus.Markdown{}
	flavor.Wrap("*")
	md := sisyphus.Convert(gmi, &flavor)
	expect := "*wrap me up\n*"
	if md != expect {
		t.Errorf("%s should be %s", md, expect)
	}
}

func TestMdQuote(t *testing.T) {
	flavor := &sisyphus.Markdown{}
	flavor.OnQuote(func(_ string) string {
		return "!"
	})
	gmi := ">>12435"
	expect := "!\n"
	md := sisyphus.Convert(gmi, flavor)
	if md != expect {
		t.Errorf("%s should be %s", md, expect)
	}
}

func TestGreentext(t *testing.T) {
	flavor := &sisyphus.Html{Greentext: true}
	gmi := "> 12435"
	expect := "<blockquote>\n<p>&gt; 12435\n</blockquote>\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}

func TestGreentextLink(t *testing.T) {
	flavor := &sisyphus.Html{Greentext: true}
	gmi := "=> / a link"
	expect := "<p><a data-friendly href='/'>=&gt; a link</a>\n"
	html := sisyphus.Convert(gmi, flavor)
	if html != expect {
		t.Errorf("%s should be %s", html, expect)
	}
}
