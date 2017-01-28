package tags_test

import (
	"testing"

	"github.com/markbates/tags"
	"github.com/stretchr/testify/require"
)

func Test_Options_String(t *testing.T) {
	r := require.New(t)
	o := tags.Options{
		"value": "Mark",
		"id":    "foo-bar",
		"class": "foo bar baz",
	}
	s := o.String()
	r.Contains(s, `value="Mark"`)
	r.Contains(s, `id="foo-bar"`)
	r.Contains(s, `class="foo bar baz"`)
}

func Test_Options_String_Escaped(t *testing.T) {
	r := require.New(t)
	o := tags.Options{
		"<b>": "<p>",
	}
	s := o.String()
	r.Equal(`&lt;b&gt;="&lt;p&gt;"`, s)
}
