// SPDX-License-Identifier: MIT

package errwrap

import (
	"errors"
	"testing"

	"github.com/issue9/assert/v3"
)

func TestBuffer(t *testing.T) {
	a := assert.New(t, false)

	buf := &Buffer{}
	buf.WString("123").WRunes([]rune("456")).WByte('7').
		Printf("%s", "89").Print("10").Println()

	a.Equal(buf.String(), "12345678910\n").
		NotError(buf.Err)

	buf.Err = errors.New("error")
	buf.Reset()
	a.Empty(buf.String()).Nil(buf.Err)
}
