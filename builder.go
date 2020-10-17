// SPDX-License-Identifier: MIT

package errwrap

import (
	"fmt"
	"strings"
)

// StringBuilder 提供缓存错误的 strings.Buffer
type StringBuilder struct {
	strings.Builder
	Err error
}

// WString 写入字符串
func (builder *StringBuilder) WString(str string) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = builder.WriteString(str)
	}
	return builder
}

// WByte 写入单个字节内容
func (builder *StringBuilder) WByte(b byte) *StringBuilder {
	return builder.WBytes([]byte{b})
}

// WBytes 写入字节内容
func (builder *StringBuilder) WBytes(data []byte) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = builder.Write(data)
	}
	return builder
}

// WRune 写入单个字节内容
func (builder *StringBuilder) WRune(r rune) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = builder.WriteRune(r)
	}
	return builder
}

// WRunes 写入字节内容
func (builder *StringBuilder) WRunes(rs []rune) *StringBuilder {
	for _, r := range rs {
		builder.WRune(r)
	}
	return builder
}

// Print 相当于 fmt.Fprint(builder, v...)
func (builder *StringBuilder) Print(v ...interface{}) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = fmt.Fprint(builder, v...)
	}
	return builder
}

// Println 相当于 fmt.Fprintln(builder, v...)
func (builder *StringBuilder) Println(v ...interface{}) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = fmt.Fprintln(builder, v...)
	}
	return builder
}

// Printf 相当于 fmt.Fprintf(builder, format, v...)
func (builder *StringBuilder) Printf(format string, v ...interface{}) *StringBuilder {
	if builder.Err == nil {
		_, builder.Err = fmt.Fprintf(builder, format, v...)
	}
	return builder
}
