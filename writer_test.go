// SPDX-License-Identifier: MIT

package errwrap

import (
	"bytes"
	"testing"

	"github.com/issue9/assert"
)

func TestWriter(t *testing.T) {
	a := assert.New(t)

	buf := new(bytes.Buffer)
	w := &Writer{
		Writer: buf,
	}
	w.WString("123").WRunes([]rune("456")).WByte('7').
		Printf("%s", "89").Print("10").Println()

	a.Equal(buf.String(), "12345678910\n").
		NotError(w.Err)
}
