package transform

import (
	"bytes"
	"strings"
	"testing"
)

const H5_JS_CONTENT_DOUBLE_QUOTE = "<!DOCTYPE html><html><head><script src=\"foobar.js\"></script></head><body><nav><h1>title</h1></nav><article>content <a href='/foobar'>foobar</a>. Follow up</article></body></html>"
const H5_JS_CONTENT_SINGLE_QUOTE = "<!DOCTYPE html><html><head><script src='foobar.js'></script></head><body><nav><h1>title</h1></nav><article>content <a href='/foobar'>foobar</a>. Follow up</article></body></html>"
const H5_JS_CONTENT_ABS_URL = "<!DOCTYPE html><html><head><script src=\"http://user@host:10234/foobar.js\"></script></head><body><nav><h1>title</h1></nav><article>content <a href=\"https://host/foobar\">foobar</a>. Follow up</article></body></html>"

// URL doesn't recognize authorities.  BUG?
//const H5_JS_CONTENT_ABS_URL = "<!DOCTYPE html><html><head><script src=\"//host/foobar.js\"></script></head><body><nav><h1>title</h1></nav><article>content <a href=\"https://host/foobar\">foobar</a>. Follow up</article></body></html>"

const CORRECT_OUTPUT_SRC_HREF = "<!DOCTYPE html><html><head><script src=\"http://base/foobar.js\"></script></head><body><nav><h1>title</h1></nav><article>content <a href=\"http://base/foobar\">foobar</a>. Follow up</article></body></html>"

func TestAbsUrlify(t *testing.T) {

	tr := &AbsURL{
		BaseURL: "http://base",
	}

	apply(t, tr, abs_url_tests)
}

type test struct {
	content  string
	expected string
}

var abs_url_tests = []test{
	{H5_JS_CONTENT_DOUBLE_QUOTE, CORRECT_OUTPUT_SRC_HREF},
	{H5_JS_CONTENT_SINGLE_QUOTE, CORRECT_OUTPUT_SRC_HREF},
	{H5_JS_CONTENT_ABS_URL, H5_JS_CONTENT_ABS_URL},
}

func apply(t *testing.T, tr Transformer, tests []test) {
	for _, test := range tests {
		out := new(bytes.Buffer)
		err := tr.Apply(out, strings.NewReader(test.content))
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if test.expected != string(out.Bytes()) {
			t.Errorf("Expected:\n%s\nGot:\n%s", test.expected, string(out.Bytes()))
		}
	}
}
