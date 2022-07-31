package bot

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	fmt.Println(fmt.Sprintf(MSG_CODE, "golang", "hello world", "a", "b", "c", "d"))
}
