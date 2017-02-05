package urlutil

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestResolve(t *testing.T) {
	resolveTestCases := []struct {
		base, path, must string
	}{
		{"https://example.com", "/foo/bar", "https://example.com/foo/bar"},
		{"https://example.com/", "/foo/bar", "https://example.com/foo/bar"},
		{"https://example.com/baz", "/foo/bar", "https://example.com/foo/bar"},
		{"https://example.com/baz/", "/foo/bar", "https://example.com/foo/bar"},
		{"https://example.com/foo/bar/", "baz", "https://example.com/foo/bar/baz"},
		{"https://example.com/foo/bar", "baz", "https://example.com/foo/baz"},
		{"https://example.com/foo/bar", "http://example.org/fred", "http://example.org/fred"},
		{"/foo/bar/", "/fred", "/fred"},
		{"foo/bar/", "/fred", "/fred"},
		{"https://example.com/foo/bar#fragment", "fred", "https://example.com/foo/fred"},
		{"https://example.com/foo/bar#fragment", "fred#f2", "https://example.com/foo/fred#f2"},
		{"https://example.com/foo/bar?query", "fred", "https://example.com/foo/fred"},
	}
	for _, tc := range resolveTestCases {
		u, _ := url.Parse(tc.base)
		res := MustResolve(u, tc.path)
		res2, err := Resolve(u, tc.path)

		assert.Equal(t, tc.must, res.String())
		if assert.Nil(t, err) {
			assert.Equal(t, tc.must, res2.String())
		}
	}
}

func TestResolveTemplate(t *testing.T) {
	vars := map[string]string{
		"dub":   "me/too",
		"hello": "Hello World!",
		"half":  "50%",
		"var":   "value",
		"who":   "fred",
		"base":  "http://example.com/home/",
		"path":  "/foo/bar",
		"v":     "6",
		"x":     "1024",
		"y":     "768",
		"empty": "",
		// undef not set
	}
	simpleStringExpansionTestCases := []struct{ tpl, expected string }{
		{"{var}", `value`},
		{"{hello}", `Hello%20World%21`},
		{"{half}", `50%25`},
		{"O{empty}X", `OX`},
		{"O{undef}X", `OX`},
		{"{x,y}", `1024,768`},
		{"{x,hello,y}", `1024,Hello%20World%21,768`},
		{"?{x,empty}", `?1024,`},
		{"?{x,undef}", `?1024`},
		{"?{undef,y}", `?768`},
		{"{var:3}", `val`},
		{"{var:30}", `value`},
	}
	for _, tc := range simpleStringExpansionTestCases {
		res := MustResolveTemplate(nil, tc.tpl, vars)
		res2, err := ResolveTemplate(nil, tc.tpl, vars)
		assert.Equal(t, tc.expected, res.String())
		if assert.Nil(t, err) {
			assert.Equal(t, tc.expected, res2.String())
		}
	}
}
