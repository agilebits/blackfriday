//
// Blackfriday Markdown Processor
// Available at http://github.com/russross/blackfriday
//
// Copyright © 2011 Russ Ross <russ@russross.com>.
// Distributed under the Simplified BSD License.
// See README.md for details.
//

//
// Unit tests for full document parsing and rendering
//

package blackfriday

import "testing"

func TestDocument(t *testing.T) {
	var tests = []string{
		// Empty document.
		"",
		"",

		" ",
		"",

		// This shouldn't panic.
		// https://github.com/russross/blackfriday/issues/172
		"[]:<",
		"<p>[]:&lt;</p>\n",

		// This shouldn't panic.
		// https://github.com/russross/blackfriday/issues/173
		"   [",
		"<p>[</p>\n",
	}
	doTests(t, tests)
}

func TestAGExtensions(t *testing.T) {
	var tests = []string{
		"Ignore [this](https://this.com)",
		"<p>Ignore [this](https://this.com)</p>\n",

		"Ignore ![this image](https://this.com)",
		"<p>Ignore ![this image](https://this.com)</p>\n",

		"Divide\n---\nDon't head",
		"<p>Divide</p>\n\n<hr>\n\n<p>Don't head</p>\n",

		"Hello *world*",
		"<p>Hello <strong>world</strong></p>\n",

		"Hello _world_",
		"<p>Hello <em>world</em></p>\n",

		"Hello *_world_*",
		"<p>Hello <strong><em>world</em></strong></p>\n",

		"Hello ~world~",
		"<p>Hello <del>world</del></p>\n",

		"Hello **world**",
		"<p>Hello **world**</p>\n",

		"This *is_p_unctuation*",
		"<p>This <strong>is_p_unctuation</strong></p>\n",

		"This should be `normal <strong>code</strong>`",
		"<p>This should be <code>normal &lt;strong&gt;code&lt;/strong&gt;</code></p>\n",

		"This should be ```multi-line\n pre-formatted *text*```",
		"<p>This should be \n<pre>multi-line\n pre-formatted *text*</pre>\n</p>\n",
		"```pre```",
		"<p>\n<pre>pre</pre>\n</p>\n",

		"1. Hello\n2. World\n* Bye\n* World",
		"<ol>\n<li>Hello</li>\n<li>World</li>\n</ol>\n\n<ul>\n<li>Bye</li>\n<li>World</li>\n</ul>\n",
	}
	doAGTests(t, tests)
}
