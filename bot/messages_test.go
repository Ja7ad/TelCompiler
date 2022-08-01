package bot

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	t.Logf(fmt.Sprintf(botResponseMessage(), "python", "test", "print('hello')", "hello world", "example stats", "@test"))
}

func TestEscapeChars(t *testing.T) {
	test := "` _ + = [ ] ) ( { } | - >"
	want := "\\` \\_ \\+ \\= \\[ \\] \\) \\( \\{ \\} \\| \\- \\>"
	escaped := escapeSpecialChar(test)
	if escaped != want {
		t.Errorf("want %s but got %s", want, escaped)
	}
	t.Log(escaped)
}
