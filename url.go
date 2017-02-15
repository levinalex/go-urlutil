// Package urlutil resolves relative uritemplate references.
package urlutil

import (
	"net/url"

	"github.com/levinalex/go-urlutil/internal/uritemplates"
)

// Resolve resolves a URI reference to a absolute URL according to the
// algorithm of net/url.ResolveReference.
// The method always returns a new URL.
//
func Resolve(u *url.URL, path string) (*url.URL, error) {
	urlpath, err := url.Parse(path)
	if u != nil {
		urlpath = u.ResolveReference(urlpath)
	}
	return urlpath, err
}

// MustResolve runs Resolve and panics on error.
func MustResolve(u *url.URL, path string) *url.URL {
	if res, err := Resolve(u, path); err != nil {
		panic(err)
	} else {
		return res
	}
}

// ResolveTemplate resolves a URI template, expands it and resolves it relative to
// the base URL.
//
func ResolveTemplate(u *url.URL, tpl string, vars map[string]string) (*url.URL, error) {
	path, _, _ := uritemplates.Expand(tpl, vars)
	return Resolve(u, path)
}

// MustResolveTemplate runs ResolveTemplate and panics on error.
func MustResolveTemplate(u *url.URL, tpl string, vars map[string]string) *url.URL {
	u, err := ResolveTemplate(u, tpl, vars)
	if err != nil {
		panic(err)
	}
	return u
}
