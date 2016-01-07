// Package frontman is a collection of useful utilities for static web
package frontman

import (
	"io"

	"github.com/gernest/frontman/down"
)

// Markdown renders a markdown document to html.The rendered output is writen to
// out.
//
// It accept markdown input as an io.reader in src. To enable syntax highlighting
// set highlighting to true. This will add css classess as supported by
// https://github.com/google/code-prettify.
//
// If inline is set to true, styles will be inlined using the style html
// attribute.
//
//  You can optionally pass a string as a theme name. The following themes are
//  shipped with this package
// 		* prettify
//		* desert
//		* doxy
//		* sunburst
//		* obsidian
func Markdown(out io.Writer, src io.Reader, highlight, inline bool, theme ...string) error {
	return down.Mark(out, src, highlight, inline, theme...)
}
