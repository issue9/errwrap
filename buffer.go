// SPDX-FileCopyrightText: 2020-2024 caixw
//
// SPDX-License-Identifier: MIT

package errwrap

import (
	"bytes"
	"fmt"
)

// Buffer 提供缓存错误的 bytes.Buffer
type Buffer struct {
	bytes.Buffer
	Err error
}

// WString 写入字符串
func (w *Buffer) WString(str string) *Buffer {
	if w.Err == nil {
		_, w.Err = w.WriteString(str)
	}
	return w
}

// WByte 写入单个字节内容
func (w *Buffer) WByte(b byte) *Buffer {
	return w.WBytes([]byte{b})
}

// WBytes 写入字节内容
func (w *Buffer) WBytes(data []byte) *Buffer {
	if w.Err == nil {
		_, w.Err = w.Write(data)
	}
	return w
}

// WRune 写入单个字节内容
func (w *Buffer) WRune(r rune) *Buffer {
	if w.Err == nil {
		_, w.Err = w.WriteRune(r)
	}
	return w
}

// WRunes 写入字节内容
func (w *Buffer) WRunes(rs []rune) *Buffer {
	for _, r := range rs {
		w.WRune(r)
	}
	return w
}

// Print 相当于 fmt.Fprint(builder, v...)
func (w *Buffer) Print(v ...interface{}) *Buffer {
	if w.Err == nil {
		_, w.Err = fmt.Fprint(w, v...)
	}
	return w
}

// Println 相当于 fmt.Fprintln(builder, v...)
func (w *Buffer) Println(v ...interface{}) *Buffer {
	if w.Err == nil {
		_, w.Err = fmt.Fprintln(w, v...)
	}
	return w
}

// Printf 相当于 fmt.Fprintf(builder, format, v...)
func (w *Buffer) Printf(format string, v ...interface{}) *Buffer {
	if w.Err == nil {
		_, w.Err = fmt.Fprintf(w, format, v...)
	}
	return w
}

// Reset 重置内容
//
// 同时也会清除错误标记。
func (w *Buffer) Reset() *Buffer {
	w.Buffer.Reset()
	w.Err = nil
	return w
}
