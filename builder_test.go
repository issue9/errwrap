// SPDX-License-Identifier: MIT

package errwrap

import (
	"testing"

	"github.com/issue9/assert"
)

func TestStringBuilder(t *testing.T) {
	a := assert.New(t)

	builder := &StringBuilder{}
	builder.WString("123").WRunes([]rune("456")).WByte('7').
		Printf("%s", "89").Print("10").Println()

	a.Equal(builder.String(), "12345678910\n").
		NotError(builder.Err)
}
